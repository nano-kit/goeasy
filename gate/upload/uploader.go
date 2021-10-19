package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type uploader struct {
	// pathPrefix is the prefix from the request URL's Path. It should end with '/'.
	pathPrefix string
	// documentRoot is the place where uploaded files are stored.
	documentRoot string
	// maxUploadSize limits the size of the uploaded content, specified with "byte".
	maxUploadSize int64
	// uploadEntry is the URL location to upload file with HTTP POST.
	uploadEntry string
}

func Uploader(pathPrefix, documentRoot string) http.HandlerFunc {
	s := uploader{
		pathPrefix:    pathPrefix,
		documentRoot:  documentRoot,
		maxUploadSize: 64 * 1024 * 1024,
		uploadEntry:   pathPrefix + "upload",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet, http.MethodHead:
			s.handleGet(w, r)
		case http.MethodPost:
			s.handlePost(w, r)
		case http.MethodPut:
			s.handlePut(w, r)
		default:
			w.Header().Set("Allow", "GET, HEAD, POST, PUT")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (s uploader) handleGet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == s.uploadEntry {
		http.Error(w, "invalid path", http.StatusNotFound)
		return
	}

	http.StripPrefix(s.pathPrefix, http.FileServer(http.Dir(s.documentRoot))).ServeHTTP(w, r)
}

func (s uploader) handlePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != s.uploadEntry {
		http.Error(w, "invalid path", http.StatusNotFound)
		return
	}

	uploadedURL, code, err := s.handleUpload(r, true)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	fmt.Fprint(w, uploadedURL)
}

func (s uploader) handlePut(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == s.uploadEntry {
		http.Error(w, "invalid path", http.StatusNotFound)
		return
	}

	uploadedURL, code, err := s.handleUpload(r, false)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	fmt.Fprint(w, uploadedURL)
}

func (s uploader) handleUpload(r *http.Request, usePartFilename bool) (uploadedURL string, code int, err error) {
	// reserve 10 MB for file parts stored in the memory.
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return "", http.StatusInternalServerError, fmt.Errorf("parse multipart form: %v", err)
	}

	// read and open the source file.
	srcFile, srcInfo, err := r.FormFile("file")
	if err != nil {
		return "", http.StatusBadRequest, fmt.Errorf("parse form: %v", err)
	}
	defer srcFile.Close()

	// check parameters
	if srcInfo.Size > s.maxUploadSize {
		return "", http.StatusRequestEntityTooLarge, fmt.Errorf("uploaded file size exceeds the limit")
	}
	if usePartFilename && srcInfo.Filename == "" {
		return "", http.StatusBadRequest, fmt.Errorf("missing the filename parameter")
	}

	// open the destination file.
	dstPath := path.Join(s.documentRoot, strings.TrimPrefix(r.URL.Path, s.pathPrefix))
	if usePartFilename {
		dstPath = path.Join(s.documentRoot, srcInfo.Filename)
	}
	if dstPath == path.Clean(s.documentRoot) {
		return "", http.StatusBadRequest, fmt.Errorf("destination path is effectively document root")
	}
	if err := os.MkdirAll(path.Dir(dstPath), 0777); err != nil {
		return "", http.StatusInternalServerError, err
	}
	dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	defer dstFile.Close()

	// save source file to destination.
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return "", http.StatusInternalServerError, err
	}

	// returns uploaded URL.
	uploadedURL = path.Join(s.pathPrefix, strings.TrimPrefix(dstPath, s.documentRoot))
	return uploadedURL, http.StatusOK, nil
}
