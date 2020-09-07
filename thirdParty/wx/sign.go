package wx

import (
	"github.com/chenyingdi/gf-toolkit/utils"
	"github.com/gogf/gf/frame/g"
)

func GetWxAppOpenid(appID, appSecret, code string) (g.Map, utils.Error) {
	c := NewSimpleClient(appID, appSecret)
	return c.Code2Session(code)
}
