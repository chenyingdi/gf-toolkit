package utils

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/iGoogle-ink/gotil"
)

type PaymentParams interface {
}

type WxPaymentParams struct {
	PaymentParams
	AppID     string `json:"app_id"`
	MchID     string `json:"mch_id"`
	ApiKey    string `json:"api_key"`
	MchName   string `json:"mch_name"`
	Detail    string `json:"detail"`
	OrderSn   string `json:"order_sn"`
	NotifyUrl string `json:"notify_url"`
	TradeType string `json:"trade_type"`
	SignType  string `json:"sign_type"`
	Openid    string `json:"openid"`
	TotalFee  int    `json:"total_fee"`
}

type AlipayPaymentParams struct {
	PaymentParams
	AppID          string
	PrivateKey     string
	PrivateKeyType alipay.PKCSType
	TotalAmount    string
	Subject        string
	OrderSn        string
	Channel        string
}

func Payment(way string, p PaymentParams, args g.Map) (g.Map, Error) {
	e := NewErr()

	switch way {
	case "wx":
		param, ok := p.(*WxPaymentParams)
		if !ok {
			e.Append(errors.New("参数类型错误！"))
		}

		prepayID, er := wxPayment(param, args)
		e.Append(er.Errs()...)

		return g.Map{"prepay_id": prepayID, "ts": gtime.Now().Unix()}, e

	case "alipay":
		param, ok := p.(*AlipayPaymentParams)
		if !ok {
			e.Append(errors.New("参数类型错误！"))
		}

		payParam, er := alipayPayment(param)
		e.Append(er.Errs()...)

		return g.Map{"pay_param": payParam, "ts": gtime.Now().Unix()}, e

	default:
		e.Append(errors.New("暂未支持的模式！"))
		return nil, e
	}
}

// 微信支付
func wxPayment(p *WxPaymentParams, args g.Map) (string, Error) {
	e := NewErr()

	// 初始化客户端
	client := wechat.NewClient(p.AppID, p.MchID, p.ApiKey, true)

	// 获取id
	ip, err := GetIp()
	e.Append(err)

	// 设置参数
	bm := make(gopay.BodyMap)

	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("body", fmt.Sprintf("%s-%s", p.MchName, p.Detail))
	bm.Set("out_trade_no", p.OrderSn)
	bm.Set("total_fee", p.TotalFee)
	bm.Set("spbill_create_ip", ip)
	bm.Set("notify_url", p.NotifyUrl)
	bm.Set("trade_type", p.TradeType)
	bm.Set("sign_type", p.SignType)
	bm.Set("openid", p.Openid)

	for i, v := range args {
		bm.Set(i, v)
	}

	// 签名
	sign := wechat.GetParamSign(p.AppID, p.MchID, p.ApiKey, bm)
	bm.Set("sign", sign)

	rsp, err := client.UnifiedOrder(bm)
	e.Append(err)

	return rsp.PrepayId, e
}

// 支付宝支付
func alipayPayment(p *AlipayPaymentParams) (string, Error) {
	e := NewErr()

	client := alipay.NewClient(p.AppID, p.PrivateKey, true)

	client.SetPrivateKeyType(p.PrivateKeyType)

	bm := make(gopay.BodyMap)
	bm.Set("total_amount", p.TotalAmount)
	bm.Set("subject", p.Subject)
	bm.Set("out_trade_no", p.OrderSn)


	access := make(g.MapStrStr)

	access["channel"] = p.Channel

	bm.Set("access_params", access)

	payParam, err := client.TradeAppPay(bm)
	e.Append(err)

	return payParam, e
}
