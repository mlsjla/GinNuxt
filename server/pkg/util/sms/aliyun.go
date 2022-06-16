package sms

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	jsoniter "github.com/json-iterator/go"
	"github.com/mlsjla/gin-nuxt/server/internal/app/setting"
)

type Aliyun struct {
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func _main(params SendParams) (_err error) {
	res := setting.Setting("sms")
	accessKeyId := jsoniter.Get([]byte(res), "accessKeyId").ToString()
	accessKeySecret := jsoniter.Get([]byte(res), "accessKeySecret").ToString()
	signName := jsoniter.Get([]byte(res), "signName").ToString()
	s := jsoniter.Get([]byte(res), "smsTemplate").ToString()

	var smsTemplates SmsTemplates
	jsoniter.UnmarshalFromString(s, &smsTemplates)

	fmt.Println("smsTemplates", smsTemplates)
	var template SmsTemplate
	for i := 0; i < len(smsTemplates); i++ {
		if smsTemplates[i].Value == params.Value {
			template = smsTemplates[i]
		}
	}

	client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(params.Mobile),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(template.TemplateCode),
		TemplateParam: tea.String(params.Params),
	}
	resp, _err := client.SendSms(sendSmsRequest)
	fmt.Println("resp", resp, _err)
	if _err != nil {
		return _err
	}

	return _err
}

//@object: *Aliyun
//@function: Send
//@description: 发送短信
//@param: key string
//@return: error

func (*Aliyun) Send(params SendParams) error {
	return _main(params)
}
