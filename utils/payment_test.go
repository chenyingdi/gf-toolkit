package utils

import (
	"github.com/iGoogle-ink/gopay/alipay"
	"testing"
)

const (
	AppPublicCert = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAirsq3/Gp2rDR3dNqlXf5ogSVG48+Rw9tuSRK0+UPYL5kfPSni2QL6BPWfgXuDQGmPm8t9oDjmjw20Y+LCx8fGuE8wK8TN/pMc8GI6iVLcWwcOOIio1Zz0aT7PILmMcoLgH6DTeBAb9/hJsUdbS6bVaFlOdPPuyHBL2qfJp063bf6QoMRv5r1sjtRHp6im/PR3lBEij+4LlgAcM7RPP77wBpQ+aPPkRD1GTdgFG7NhcRjyK77P4L17U+3cU7wOHZJkQhSZoNhIb2xpj4mXq+tW+eKl77JOJNg/kG7UfCsdHF8fjcR7NqtaRFmQP8qWZ6oSW5nCyWC9z5b/sJ3lZCeRwIDAQAB`
)

func TestPayment(t *testing.T) {
	res, err := Payment("alipay", &AlipayPaymentParams{
		AppID: "2016101600703352",
		PrivateKey: `MIIEowIBAAKCAQEAiCzta8cSdfBGur8iwioS21YPg4jJ8mO9tRqPPMN1npIty0A+9aE2U+9sF/3Bbp27o4zYczAY6u6YwMzNt/nqEKiOn0RnIodc6rXgIJAi/kcASYtZjx3AZ7BYI85z0h76ED/WO0Sij5Mj4Py5/g2pbwurafBNvwsdW0c2yAz3Kt7LCjB+Xs0BC6L0Pg/1OYxW7UioCXcu1sONtrtniPHeE1YIotZc0GCz+7uPDblH8u6u7VwEjLTLAdgvAxn/RMVUqpOhj6LZA4eQ+D6KYpTw2+N/rNedxWULF7DU0jH+IvJeX8MEOwSxwl6hgfIXZoMWfnJ6AcXK+G5It1etBNU/OwIDAQABAoIBAEkKUA6PKIe56TYhbzNV/edHseqLZEZu2UJajJm6/UugfI+YVizJ26oJsaDsquP6FrsSwMaH9dWNRMGlGHKlybZFsroapncOw/fgtebBaQOacb0A0XjCLIFxRNVv1w/NLbOpie8gUVFRSt1SsTBjg43cZITeL9VY0NY0zF1hFvIW5arZOaKsAQXXUC6ZvbyRjk/PoN0eXg2+HKQjR5owKPkfV69uHOAGU3vNrgI73WM3ZTeaqB5w/tbbXtd+fFGLTJ93lBjLH8qfPhl3NQsVYl3zxHz6ffkyMy0ACzcsReJmcDYsAnhBagFQ+qh7iZFFVspVYx9CtCfIP33x/uOchgECgYEAw9x3wyFRgAUfFpOEhBiki6qolT8ksyUZtVdJP+0jKSWUeUSWz5XZw8MmCJxwG/mtbHcSbGuT8PDCFkZo7Vg0+RXqM4hfE/A46NXxPngM0vwaYjmxzDSWBpOjLVf+wsaOGBvJe2Wy6my1jM2svS2ClNlvP41+7nlMpQBhSU+AXqsCgYEAsfzoH5iP3li4JACy+D+sK/bIvB8KKGihusJ9mt4rM97lrRcglWYVTagWcbO/j7k/zMgWU4xDPusGjsDn7KL7kmJo94WyhS3hpQp+IV5lS4bVkK9ltbCeOJtzb2CgjStczEqHC+PbQj75oJQunZiO9GvxxLuL0VnA/UaB49efYbECgYBVUgQlz6zWLOT3C7oNZULAyM35ffE5zO6fDXAOVfocIY/FJ/jeYvPjEG7QD33S4fgHKPOwoUhoDCkwVOm+gs9ItqA4ZK6uW9Yer5wQz3Ees822flWSlFHKeaP6y7tiE+awX+JsS8gd0M9hj/Mw0dNxjiqlL7lcnyChPEIlEmnkUwKBgBQPJrH45a3vXcFg+sqTknnZ9EGPVfu73w5HQRSlGUVdR5E8XAW6XYhE+1KRKXOvMwuHOUztL971aeXIw9qde7DBuoa64KW1yAprpk9obg4XAhauTc4uO4axrk7NGwsN6gV0GMg0Q8+xfTyltqM4QFQ3niXH5TgQ33kr6xOch2/hAoGBAI3XxmO9GetEXPLGjKML06TtDDh3NeW6GOjo+f2RyxpTLXWrkXPjMF6UeL6VlSXZznMkjIqRvIz/ztty1PndKz4HMhBnoC0nhl+cp6ciKjr9zxNCflEPL0VTp1FiVj0oMGkxEDn08mOM6/XjkazpyahHw/IyHAEKegv+9pNvJWXb`,
		PrivateKeyType: alipay.PKCS1,
		TotalAmount: "0.06",
		Subject: "九域优谷-商品",
		OrderSn: "test1",
		Channel: "ALIPAYAPP",
	}, nil)
	t.Log(err)

	t.Log(res)

	//res, err := Payment("wx", &WxPaymentParams{
	//	AppID: "wx6eed9af8e39fc073",
	//	ApiKey: "8c42c5af8a59e751a625f8e5c5687119",
	//	MchID: "1597844801",
	//	MchName: "九域优谷",
	//	Detail: "商品购买",
	//	OrderSn: "test",
	//	NotifyUrl: "https://47.108.169.164",
	//	TradeType: "JSAPI",
	//	SignType: "MD5",
	//	Openid: "oMx7V5PFR9rwzJXwTuwkMcLk7bxk",
	//	TotalFee: 10,
	//}, nil)
	//t.Log(err)
	//t.Log(res)
}
