package img

import (
	"errors"
	"fmt"
	"hash/adler32"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

const (
	dpi         = 72.00
	maxFontSize = 512.00
	maxAge      = 7 * 24 * 3600
)

type placeholder struct {
	width           int
	height          int
	text            string
	foreground      color.RGBA
	background      color.RGBA
	ttfPath         string
	ttf             []byte
	backgroundImage image.Image
	marginRatio     float64

	font *truetype.Font
}

func hexToRGB(h string) (uint8, uint8, uint8, error) {
	rgb, err := strconv.ParseUint(h, 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	return uint8(rgb >> 16), uint8((rgb >> 8) & 0xFF), uint8(rgb & 0xFF), nil
}

func normalizeHex(h string) string {
	h = strings.TrimPrefix(h, "#")
	if len(h) != 3 && len(h) != 6 {
		return ""
	}
	if len(h) == 3 {
		h = h[:1] + h[:1] + h[1:2] + h[1:2] + h[2:] + h[2:]
	}
	return h
}

func paramToColor(param, defaultValue string) (color.RGBA, error) {
	if len(param) == 0 {
		param = defaultValue
	}

	hexColor := normalizeHex(param)
	if len(hexColor) == 0 {
		return color.RGBA{}, errors.New("bad hex color format")
	}

	R, G, B, err := hexToRGB(hexColor)
	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{R, G, B, 255}, nil
}

func parseQuery(q url.Values) (ph placeholder, err error) {
	width, err := strconv.ParseUint(q.Get("w"), 10, 32)
	if err != nil {
		return ph, fmt.Errorf("width is not an integer: %v", err)
	}
	if width < 1 || width > 5000 {
		return ph, fmt.Errorf("width should be [1, 5000]")
	}

	height, err := strconv.ParseUint(q.Get("h"), 10, 32)
	if err != nil {
		return ph, fmt.Errorf("height is not an integer: %v", err)
	}
	if height < 1 || height > 5000 {
		return ph, fmt.Errorf("height shoud be [1, 5000]")
	}

	text := q.Get("txt")
	if len(text) > 50 {
		return ph, fmt.Errorf("text should not be greater than 50 characters")
	}

	foreground, err := paramToColor(q.Get("fg"), "969696")
	if err != nil {
		return ph, fmt.Errorf("bad value for foreground color: %v", err)
	}

	background, err := paramToColor(q.Get("bg"), "CCCCCC")
	if err != nil {
		return ph, fmt.Errorf("bad value for background color: %v", err)
	}

	ph.width = int(width)
	ph.height = int(height)
	ph.text = text
	ph.foreground = foreground
	ph.background = background
	ph.ttf = goregular.TTF
	ph.backgroundImage = nil
	ph.marginRatio = 0.2

	if ph.ttfPath != "" {
		ph.ttf, err = ioutil.ReadFile(ph.ttfPath)
		if err != nil {
			return ph, fmt.Errorf("read font file: %v", err)
		}
	}

	ph.font, err = freetype.ParseFont(ph.ttf)
	if err != nil {
		return ph, fmt.Errorf("parse font: %v", err)
	}

	return
}

func Placeholder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		placeholder, err := parseQuery(r.URL.Query())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// set HTTP cache
		w.Header().Set("Expires", time.Now().Add(maxAge*time.Second).Format(http.TimeFormat))
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", maxAge))
		etag := placeholder.etag()
		w.Header().Set("Etag", etag)
		if match := r.Header.Get("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		// generate image
		img, err := placeholder.image()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// write as png
		w.Header().Set("Content-Type", "image/png")
		if err := png.Encode(w, img); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (ph placeholder) etag() string {
	spec := fmt.Sprintf("%x,%x,%s,%x,%x,%x,%x,%x,%x,%x,%x,%s,%f",
		ph.width, ph.height, ph.text,
		ph.foreground.R, ph.foreground.G, ph.foreground.B, ph.foreground.A,
		ph.background.R, ph.background.G, ph.background.B, ph.background.A,
		ph.ttfPath, ph.marginRatio)
	hash := adler32.Checksum([]byte(spec))
	return fmt.Sprintf(`"%x"`, hash)
}

// image returns a placeholder image with the given text, width & height
func (ph placeholder) image() (image.Image, error) {
	if ph.width < 0 || ph.height < 0 {
		return nil, errors.New("values for width or height must be positive")
	}
	if ph.width == 0 && ph.height == 0 {
		return nil, errors.New("values for width or height must be positive")
	}

	if ph.width == 0 {
		ph.width = ph.height
	} else if ph.height == 0 {
		ph.height = ph.width
	}

	if ph.text == "" {
		ph.text = strconv.Itoa(ph.width) + " x " + strconv.Itoa(ph.height)
	}

	img := image.NewRGBA(image.Rect(0, 0, ph.width, ph.height))

	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(ph.font)
	c.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 0}))
	c.SetDst(img)
	c.SetClip(img.Bounds())
	c.SetHinting(font.HintingNone)

	// draw the background
	draw.Draw(img, img.Bounds(), image.NewUniform(ph.background), image.ZP, draw.Src)

	// draw background image
	if ph.backgroundImage != nil {
		bgimg := imaging.Fill(ph.backgroundImage, ph.width, ph.height, imaging.Center, imaging.Lanczos)
		draw.Draw(img, img.Bounds(), bgimg, image.ZP, draw.Src)
	}

	if ph.text != "" {
		// draw with scaled fontsize to get the real text extent
		fontsize, actwidth := maxPointSize(ph.text, c,
			int(float64(ph.width)*(1.0-ph.marginRatio)),
			int(float64(ph.height)*(1.0-ph.marginRatio)))
		actheight := c.PointToFixed(fontsize/2.0) / 64
		xcenter := (float64(ph.width) / 2.0) - (float64(actwidth) / 2.0)
		ycenter := (float64(ph.height) / 2.0) + (float64(actheight) / 2.0)

		// draw the text
		c.SetFontSize(fontsize)
		c.SetSrc(image.NewUniform(ph.foreground))
		_, err := c.DrawString(ph.text, freetype.Pt(int(xcenter), int(ycenter)))
		if err != nil {
			return nil, fmt.Errorf("draw text: %v", err)
		}
	}

	return img, nil
}

// maxPointSize returns the maximum point size we can use to fit text inside width and height
// as well as the resulting text-width in pixels
func maxPointSize(text string, c *freetype.Context, width, height int) (float64, int) {
	// never let the font size exceed the requested height
	fontsize := maxFontSize
	for int(c.PointToFixed(fontsize)/64) > height {
		fontsize -= 2
	}

	// find the biggest matching font size for the requested width
	var actwidth int
	for actwidth = width + 1; actwidth > width; fontsize -= 2 {
		c.SetFontSize(fontsize)

		textExtent, err := c.DrawString(text, freetype.Pt(0, 0))
		if err != nil {
			return 0, 0
		}

		actwidth = int(float64(textExtent.X) / 64)
	}

	return fontsize, actwidth
}
