package wx

import (
	"github.com/chenyingdi/gf-toolkit/utils"
	"github.com/gogf/gf/frame/g"
)

// 微信支付
func Payment(fee int, orderSn, appID, mchID, apiKey, openid string) (g.Map, utils.Error) {
	e := utils.NewErr()

	c := NewPaymentClient(appID, mchID, apiKey)

	p := NewParams()

	ip, err := utils.GetIp()
	e.Append(err)

	p.SetString("appid", appID).
		SetString("mch_id", mchID).
		SetString("nonce_str", utils.GeneNonceStr(32)).
		SetString("body", g.Cfg().GetString("app.appName")+"-商品购买").
		SetString("out_trade_no", orderSn).
		SetInt("total_fee", fee).
		SetString("spbill_create_ip", ip).
		SetString("notify_url", g.Cfg().GetString("app.appNotifyUrl")).
		SetString("openid", openid).
		SetString("trade_type", "JSAPI")

	res, er := c.UnifiedOrder(p)
	e.Append(er.Errs()...)

	return res, e
}

