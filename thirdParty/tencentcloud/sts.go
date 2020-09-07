package tencentcloud

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"time"
)

// 获取sts临时密钥
func GetStsKey(appID, secretID, secretKey, bucket, path, region string) (g.Map, error) {
	c := sts.NewClient(
		secretID,
		secretKey,
		nil,
	)

	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + region + ":uid/" + appID + ":" + bucket + "/" + path,
					},
				},
			},
		},
	}

	res, err := c.GetCredential(opt)
	if err != nil {
		return nil, err
	}
	return g.Map{
		"credentials": res.Credentials,
		"ts":          gtime.Now().Unix(),
	}, nil
}
