package wx

const (
	Code2SessionUrl              = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	Oauth2Url                    = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	UnifiedOrderUrl              = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	PushMessageUrl               = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s"
	RefundUrl                    = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	CloseOrderUrl                = "https://api.mch.weixin.qq.com/pay/closeorder"
	GetAccessTokenUrl            = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	GetDailyRetainUrl            = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=%s"
	GetMonthlyRetainUrl          = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=%s"
	GetWeeklyRetainUrl           = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=%s"
	GetDailySummaryUrl           = "https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=%s"
	GetDailyVisitTrendUrl        = "https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=%s"
	GetWeeklyVisitTrendUrl       = "https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=%s"
	GetMonthlyVisitTrendUrl      = "https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=%s"
	GetDailyUserPortraitUrl      = "https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=%s"
	GetDailyVisitDistributionUrl = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=%s"
	GetDailyVisitPageUrl         = "https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=%s"
)
