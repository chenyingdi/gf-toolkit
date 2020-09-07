package wx

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/chenyingdi/gf-toolkit/utils"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"net/http"
)

type Client interface {
	Code2Session(code string) (g.Map, utils.Error)                                // 微信小程序获取openid和session
	Oauth2(code string) (g.Map, utils.Error)                                      // web端oauth2登录
	UnifiedOrder(p *utils.RequestParams) (g.Map, utils.Error)                     // 微信支付统一下单
	Refund(p *utils.RequestParams, keyPath, certPath string) (g.Map, utils.Error) // 微信支付退款
	CloseOrder(p *utils.RequestParams) (g.Map, utils.Error)                       // 微信支付关闭订单
	GetAccessToken() (g.Map, utils.Error)                                         // 获取access_token
	GetDailyRetain(accessToken, date string) (g.Map, utils.Error)                 // 微信小程序获取日访问留存
	GetWeeklyRetain(accessToken, begin, end string) (g.Map, utils.Error)          // 微信小程序获取周访问留存
	GetMonthlyRetain(accessToken string, ym string) (g.Map, utils.Error)          // 微信小程序获取月访问留存
	GetDailySummary(accessToken, date string) (g.Map, utils.Error)                // 微信小程序获取日统计
	GetDailyVisitTrend(accessToken, date string) (g.Map, utils.Error)             // 微信小程序获取日趋势
	GetWeeklyVisitTrend(accessToken, begin, end string) (g.Map, utils.Error)      // 微信小程序获取周趋势
	GetMonthlyVisitTrend(accessToken string, ym string) (g.Map, utils.Error)      // 微信小程序获取月趋势
	GetDailyUserPortrait(accessToken, date string) (g.Map, utils.Error)           // 微信小程序日用户画像
	GetDailyVisitDistribution(accessToken, date string) (g.Map, utils.Error)      // 微信小程序日访问分布
	GetDailyVisitPage(accessToken, date string) (g.Map, utils.Error)              // 微信小程序日页面数据
}

type client struct {
	AppID     string // app id
	AppSecret string // app密钥
	MchID     string // 商户号
	ApiKey    string // api密钥
}

// 简单模式的client，用于获取openid，以及access_token
func NewSimpleClient(appID, appSecret string) Client {
	return &client{
		AppID:     appID,
		AppSecret: appSecret,
	}
}

// 支付模式的client，用于支付，退款
func NewPaymentClient(appID, mchID, apiKey string) Client {
	return &client{
		AppID:  appID,
		MchID:  mchID,
		ApiKey: apiKey,
	}
}

// 微信小程序获取openid和session
func (c *client) Code2Session(code string) (g.Map, utils.Error) {
	e := utils.NewErr()
	res, err := http.Get(fmt.Sprintf(Code2SessionUrl, c.AppID, c.AppSecret, code))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// web端微信oauth2登录
func (c *client) Oauth2(code string) (g.Map, utils.Error) {
	e := utils.NewErr()

	res, err := http.Get(fmt.Sprintf(Oauth2Url, c.AppID, c.AppSecret, code))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 微信支付统一下单api
func (c *client) UnifiedOrder(p *utils.RequestParams) (g.Map, utils.Error) {
	e := utils.NewErr()

	// 签名
	c.signParamMD5(p, c.ApiKey)

	buf, err := xml.Marshal(utils.Xml(p.Values()))
	e.Append(err)

	res, err := http.Post(UnifiedOrderUrl, "application/xml", bytes.NewReader(buf))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 微信支付退款
func (c *client) Refund(p *utils.RequestParams, keyPath, certPath string) (g.Map, utils.Error) {
	e := utils.NewErr()

	// 签名
	c.signParamMD5(p, c.ApiKey)

	buf, err := xml.Marshal(utils.Xml(p.Values()))
	e.Append(err)

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	e.Append(err)

	cl := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	res, err := cl.Post(RefundUrl, "application/xml", bytes.NewReader(buf))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 微信支付关闭订单
func (c *client) CloseOrder(p *utils.RequestParams) (g.Map, utils.Error) {
	e := utils.NewErr()

	// 签名
	c.signParamMD5(p, c.ApiKey)

	buf, err := xml.Marshal(utils.Xml(p.Values()))
	e.Append(err)

	res, err := http.Post(CloseOrderUrl, "application/xml", bytes.NewReader(buf))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取接口调用凭证
func (c *client) GetAccessToken() (g.Map, utils.Error) {
	e := utils.NewErr()

	res, err := http.Post(
		fmt.Sprintf(GetAccessTokenUrl, c.AppID, c.AppSecret),
		"application/x-www-form-urlencoded", nil)
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取日访问留存
func (c *client) GetDailyRetain(accessToken, date string) (g.Map, utils.Error) {
	e := utils.NewErr()

	req, err := gjson.Encode(g.Map{
		"begin_date": date,
		"end_date":   date,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetDailyRetainUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取周访问留存
func (c *client) GetWeeklyRetain(accessToken, begin, end string) (g.Map, utils.Error) {
	e := utils.NewErr()

	req, err := gjson.Encode(g.Map{
		"begin_date": begin,
		"end_date":   end,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetWeeklyRetainUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取月访问留存
func (c *client) GetMonthlyRetain(accessToken string, ym string) (g.Map, utils.Error) {
	e := utils.NewErr()

	begin := ym + "01"
	end := ym + gconv.String(gtime.Now().DaysInMonth())

	req, err := gjson.Encode(g.Map{
		"begin_date": begin,
		"end_date":   end,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetMonthlyRetainUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取日统计
func (c *client) GetDailySummary(accessToken, date string) (g.Map, utils.Error) {
	e := utils.NewErr()

	req, err := gjson.Encode(g.Map{
		"begin_date": date,
		"end_date":   date,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetDailySummaryUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取日趋势
func (c *client) GetDailyVisitTrend(accessToken, date string) (g.Map, utils.Error) {
	e := utils.NewErr()

	req, err := gjson.Encode(g.Map{
		"begin_date": date,
		"end_date":   date,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetDailyVisitTrendUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取周趋势
func (c *client) GetWeeklyVisitTrend(accessToken, begin, end string) (g.Map, utils.Error) {
	e := utils.NewErr()

	req, err := gjson.Encode(g.Map{
		"begin_date": begin,
		"end_date":   end,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetWeeklyVisitTrendUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取月趋势
func (c *client) GetMonthlyVisitTrend(accessToken string, ym string) (g.Map, utils.Error) {
	e := utils.NewErr()

	begin := ym + "01"
	end := ym + gconv.String(gtime.Now().DaysInMonth())

	req, err := gjson.Encode(g.Map{
		"begin_date": begin,
		"end_date":   end,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetMonthlyVisitTrendUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取日用户画像
func (c *client) GetDailyUserPortrait(accessToken, date string) (g.Map, utils.Error) {
	e := utils.NewErr()

	req, err := gjson.Encode(g.Map{
		"begin_date": date,
		"end_date":   date,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetDailyUserPortraitUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取日用户分布
func (c *client) GetDailyVisitDistribution(accessToken, date string) (g.Map, utils.Error) {
	e := utils.NewErr()

	req, err := gjson.Encode(g.Map{
		"begin_date": date,
		"end_date":   date,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetDailyVisitDistributionUrl, accessToken),
		"application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 获取日页面数据
func (c *client) GetDailyVisitPage(accessToken, date string) (g.Map, utils.Error) {
	e := utils.NewErr()

	req, err := gjson.Encode(g.Map{
		"begin_date": date,
		"end_date":   date,
	})
	e.Append(err)

	res, err := http.Post(
		fmt.Sprintf(GetDailyVisitPageUrl, accessToken), "application/json", bytes.NewReader(req))
	e.Append(err)

	data, er := utils.GetMapFromJsonResp(res)
	e.Append(er.Errs()...)

	return data, e
}

// 签名
func (c *client) signParamMD5(p *utils.RequestParams, key string) {
	p.SetString("sign", GeneSign(p.Values(), key))
}
