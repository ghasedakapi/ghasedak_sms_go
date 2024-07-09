// package ghasedak

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/go-resty/resty/v2"
// )

// type Ghasedak struct {
// 	url     string
// 	headers map[string]string
// 	client  *resty.Client
// }

// func NewGhasedak(apiKey string) *Ghasedak {
// 	headers := map[string]string{
// 		"Content-Type":  "application/json",
// 		"Accept":        "application/json",
// 		"cache-control": "no-cache",
// 		"ApiKey":        apiKey,
// 	}
// 	client := resty.New()

// 	return &Ghasedak{
// 		url:     "https://gw.ghasedak.me/Rest/api/v1/WebService/",
// 		headers: headers,
// 		client:  client,
// 	}
// }

// func (g *Ghasedak) CheckSmsStatus(query map[string]interface{}) (*ResponseDto, error) {
// 	queryString := BuildQueryString(g.url+"CheckSmsStatus", map[string]string{"Type": fmt.Sprintf("%v", query["type"])}, "Ids", query["ids"].([]string))
// 	resp, err := g.client.R().SetHeaders(g.headers).Get(queryString)
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) GetAccountInformation() (*ResponseDto, error) {
// 	resp, err := g.client.R().SetHeaders(g.headers).Get(g.url + "GetAccountInformation")
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) GetReceivedSmses(query map[string]string) (*ResponseDto, error) {
// 	queryString := BuildQueryString(g.url+"GetReceivedSmses", query, "", nil)
// 	resp, err := g.client.R().SetHeaders(g.headers).Get(queryString)
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) GetReceivedSmsesPaging(query map[string]string) (*ResponseDto, error) {
// 	queryString := BuildQueryString(g.url+"GetReceivedSmsesPaging", query, "", nil)
// 	resp, err := g.client.R().SetHeaders(g.headers).Get(queryString)
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) GetOtpTemplateParameters(templateName string) (*ResponseDto, error) {
// 	queryString := fmt.Sprintf("%sGetOtpTemplateParameters?TemplateName=%s", g.url, templateName)
// 	resp, err := g.client.R().SetHeaders(g.headers).Get(queryString)
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) SendSimpleSms(command map[string]interface{}) (*ResponseDto, error) {
// 	resp, err := g.client.R().SetHeaders(g.headers).SetBody(command).Post(g.url + "SendSingleSMS")
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) SendBulkSms(command map[string]interface{}) (*ResponseDto, error) {
// 	resp, err := g.client.R().SetHeaders(g.headers).SetBody(command).Post(g.url + "SendBulkSMS")
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) SendPairToPairSms(command map[string]interface{}) (*ResponseDto, error) {
// 	resp, err := g.client.R().SetHeaders(g.headers).SetBody(command).Post(g.url + "SendPairToPairSMS")
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) SendOtpSmsOld(command map[string]interface{}) (*ResponseDto, error) {
// 	resp, err := g.client.R().SetHeaders(g.headers).SetBody(command).Post(g.url + "SendOtpSMSOld")
// 	return parseResponse(resp, err)
// }

// func (g *Ghasedak) SendOtpSms(command map[string]interface{}) (*ResponseDto, error) {
// 	resp, err := g.client.R().SetHeaders(g.headers).SetBody(command).Post(g.url + "SendOtpSMS")
// 	return parseResponse(resp, err)
// }

// func parseResponse(resp *resty.Response, err error) (*ResponseDto, error) {
// 	if err != nil {
// 		return &ResponseDto{IsSuccess: false, Message: err.Error(), StatusCode: resp.StatusCode()}, err
// 	}
// 	var result map[string]interface{}
// 	if err := json.Unmarshal(resp.Body(), &result); err != nil {
// 		return &ResponseDto{IsSuccess: false, Message: err.Error(), StatusCode: resp.StatusCode()}, err
// 	}
// 	return &ResponseDto{IsSuccess: true, Message: string(resp.Body()), StatusCode: resp.StatusCode()}, nil
// }
