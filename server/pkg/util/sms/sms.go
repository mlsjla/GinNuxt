package sms

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/mlsjla/gin-nuxt/server/internal/app/setting"
)

// OSS 对象存储接口

type SMS interface {
	Send(params SendParams) error
}

type SendParams struct {
	Mobile string
	Value  string
	Params string
}

type SmsParams struct {
	Code int `json:"code"`
}

type SmsTemplate struct {
	Label        string `json:"label"`
	Value        string `json:"value"`
	TemplateCode string `json:"templateCode"`
	Template     string `json:"template"`
}

type SmsTemplates []SmsTemplate

// NewOss OSS的实例化方法

func NewSMS() SMS {
	res := setting.Setting("sms")
	smsType := jsoniter.Get([]byte(res), "smsType", 0).ToInt()
	switch smsType {
	case 0:
		return &Aliyun{}
	default:
		return &Aliyun{}
	}
}
