package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"github.com/nano-kit/goeasy/internal/ierr"
	ijson "github.com/nano-kit/goeasy/internal/json"
	"github.com/nano-kit/goeasy/servers/liveuser"
)

type Wx struct {
	httpClient *http.Client
	appID      string // 小程序 appId
	secret     string // 小程序 appSecret
}

func (w *Wx) Init() {
	w.httpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 2 * time.Second,
			}).DialContext,
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 2,
			IdleConnTimeout:     1 * time.Minute,
			MaxConnsPerHost:     10,
		},
		Timeout: 2 * time.Second,
	}
	w.appID = os.Getenv("WX_APP_ID")
	w.secret = os.Getenv("WX_APP_SECRET")
	if w.appID == "" || w.secret == "" {
		logger.Warn("set WX_APP_ID and WX_APP_SECRET to env")
	}
}

type SessionResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func (w *Wx) authCode2Session(code string) (ses SessionResponse, err error) {
	targetAddress := "https://api.weixin.qq.com/sns/jscode2session"
	req, err := http.NewRequest("GET", targetAddress, nil)
	if err != nil {
		logger.Warn(err)
		return
	}

	q := req.URL.Query()
	q.Add("appid", w.appID)
	q.Add("secret", w.secret)
	q.Add("js_code", code)
	q.Add("grant_type", "authorization_code")
	req.URL.RawQuery = q.Encode()

	resp, err := w.httpClient.Do(req)
	if err != nil {
		logger.Warn(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warn(err)
		return
	}
	if err = json.Unmarshal(body, &ses); err != nil {
		logger.Warn(err)
		return
	}
	if ses.ErrCode != 0 {
		err = fmt.Errorf("auth.code2Session: errcode(%v) %v", ses.ErrCode, ses.ErrMsg)
		logger.Warn(err)
		return
	}
	return
}

func (w *Wx) Login(ctx context.Context, req *liveuser.LoginReq, res *liveuser.LoginRes) error {
	logger.Infof("got wx.login request: %v", ijson.Stringify(req))
	if req.Code == "" {
		err := ierr.BadRequest("empty code")
		logger.Warn(err)
		return err
	}
	ses, err := w.authCode2Session(req.Code)
	if err != nil {
		return ierr.BadRequest("wx.login: %v", err)
	}
	logger.Infof("got session: %v", ijson.Stringify(ses))
	return nil
}
