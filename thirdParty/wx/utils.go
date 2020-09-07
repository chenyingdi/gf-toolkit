package wx

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/chenyingdi/gf-toolkit/utils"
	"strings"
)

// 签名
func GeneSign(args map[string]interface{}, key string) string {
	// 1. 字典序排序
	stringA := utils.ParseMap(args)
	stringSignTemp := ""

	// 2. 与key拼接得到stringSignTemp
	if key != "" {
		stringSignTemp = stringA + "&key=" + key
	}

	// 3. 加密
	m := md5.New()
	m.Write([]byte(stringSignTemp))

	return strings.ToUpper(hex.EncodeToString(m.Sum(nil)))
}

// 校验签名
func CheckSign(key string, params map[string]interface{}) (bool, error) {
	var (
		ok      bool
		sign    string
		newSign string
	)

	sign, ok = params["sign"].(string)
	if !ok {
		return false, errors.New("sign类型错误！")
	}

	delete(params, "sign")

	newSign = GeneSign(params, key)

	return newSign == sign, nil
}
