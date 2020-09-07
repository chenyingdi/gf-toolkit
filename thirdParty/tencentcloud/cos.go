package tencentcloud

import (
	"bytes"
	"context"
	"github.com/chenyingdi/gf-toolkit/utils"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"strings"
)

// 上传对象
func CosUpload(rawurl, secretID, secretKey, filePath string, f []byte) utils.Error {
	e := utils.NewErr()
	u, _ := url.Parse(rawurl)

	b := &cos.BaseURL{BucketURL: u}

	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})

	name := filePath + "/" + strings.ReplaceAll(uuid.New().String(), "-", "")

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
