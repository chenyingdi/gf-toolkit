package tencentcloud

import (
	"bytes"
	"context"
	"github.com/chenyingdi/gf-toolkit/utils"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 下载对象
func CosGet(rawurl, secretID, secretKey, name string) ([]byte, utils.Error) {
	e := utils.NewErr()
	u, _ := url.Parse(rawurl)

	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})

	resp, err := client.Object.Get(context.Background(), name, nil)
	e.Append(err)

	body, err := ioutil.ReadAll(resp.Body)
	e.Append(err)

	return body, e
}

// 上传对象
func CosUpload(rawurl, secretID, secretKey, filePath, fileName string, f []byte) utils.Error {
	e := utils.NewErr()
	u, _ := url.Parse(rawurl)

	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})

	name := filePath + "/" + fileName

	r := bytes.NewReader(f)

	_, err := client.Object.Put(context.Background(), name, r, nil)
	e.Append(err)

	return e
}

// 删除对象
func CosDelete(rawurl, secretID, secretKey, name string) utils.Error {
	e := utils.NewErr()
	u, _ := url.Parse(rawurl)

	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})

	_, err := client.Object.Delete(context.Background(), name)
	e.Append(err)

	return e
}
