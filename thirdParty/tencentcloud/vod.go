package tencentcloud

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/chenyingdi/gf-toolkit/utils"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/google/uuid"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vodSdk "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod/v20180717"
	"github.com/tencentyun/vod-go-sdk"
	"net/url"
	"strings"
	"time"
)

// 上传视频
func VodUpload(secretID, secretKey, region, filePath, coverPath, procedure, mediaType string) (g.Map, utils.Error) {
	e := utils.NewErr()

	// 1. 初始化上传对象
	client := &vod.VodUploadClient{
		SecretId:  secretID,
		SecretKey: secretKey,
	}

	// 2. 构造上传请求对象
	req := vod.NewVodUploadRequest()
	req.MediaFilePath = common.StringPtr(filePath)
	req.MediaName = common.StringPtr(strings.ReplaceAll(uuid.New().String(), "-", ""))
	req.MediaType = common.StringPtr(mediaType)
	req.CoverFilePath = common.StringPtr(coverPath)
	req.Procedure = common.StringPtr(procedure)
	req.ConcurrentUploadNumber = common.Uint64Ptr(5)

	// 3. 调用上传
	rsp, err := client.Upload(region, req)
	e.Append(err)

	return gconv.Map(rsp.ToJsonString()), e
}

// 删除视频
func VodDelete(secretID, secretKey, fileID string) (g.Map, utils.Error) {
	e := utils.NewErr()

	credential := common.NewCredential(
		secretID,
		secretKey,
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vod.tencentcloudapi.com"
	client, _ := vodSdk.NewClient(credential, "", cpf)

	request := vodSdk.NewDeleteMediaRequest()

	params := fmt.Sprintf("{\"FileId\":\"%s\"}", fileID)

	err := request.FromJsonString(params)
	if err != nil {
		panic(err)
	}
	response, err := client.DeleteMedia(request)
	e.Append(err)

	return gconv.Map(response.ToJsonString()), e
}

// 获取防盗链
func GetRefererUrl(rawurl, key string, rLimit, exper int) (string, utils.Error) {
	e := utils.NewErr()

	u, err := url.Parse(rawurl)
	e.Append(err)

	t := fmt.Sprintf("%x", gtime.Now().Add(300*time.Second).Unix())

	l := strings.Split(u.Path, "/")

	// dir: 路径
	dir := strings.Join(l[0:len(l)-1], "/") + "/"

	// us: 随机字符串
	us := utils.GeneNonceStr(10)

	// 将 key dir t exper rLimit us拼接
	s := fmt.Sprintf("%s%s%s%d%d%s", key, dir, t, exper, rLimit, us)

	// md5签名
	h := md5.New()

	h.Write([]byte(s))

	sign := hex.EncodeToString(h.Sum(nil))

	referer := fmt.Sprintf(
		"%s?t=%s$exper=%d&rlimit=%d&us=%s&sign=%s",
		rawurl, t, exper, rLimit, us, sign,
	)

	return referer, e
}
