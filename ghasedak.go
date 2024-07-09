package ghasedak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	// "github.com/ghasedakapi/ghasedak_sms_go/helper" // update with the actual import path
)

type Ghasedak struct {
	url     string
	headers map[string]string
}

func NewGhasedak(apiKey string) *Ghasedak {
	return &Ghasedak{
		url: "https://gw.ghasedak.me/Rest/api/v1/WebService/",
		headers: map[string]string{
			"Content-Type":  "application/json",
			"Accept":        "application/json",
			"cache-control": "no-cache",
			"ApiKey":        apiKey,
		},
	}
}

func (g *Ghasedak) request(method, endpoint string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, g.url+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range g.headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (g *Ghasedak) CheckSmsStatus(query map[string]interface{}) (string, error) {
	params := map[string]string{
		"Type": fmt.Sprint(query["type"]),
	}
	ids := query["ids"].([]string)
	queryString := BuildQueryString(g.url+"CheckSmsStatus", params, "Ids", ids)
	response, err := g.request("GET", queryString, nil)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) GetAccountInformation() (string, error) {
	response, err := g.request("GET", "GetAccountInformation", nil)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) GetReceivedSmses(query map[string]interface{}) (string, error) {
	params := map[string]string{
		"LineNumber": query["line_number"].(string),
		"IsRead":     fmt.Sprint(query["is_read"]),
	}
	queryString := BuildQueryString(g.url+"GetReceivedSmses", params, "", nil)
	response, err := g.request("GET", queryString, nil)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) GetReceivedSmsesPaging(query map[string]interface{}) (string, error) {
	params := map[string]string{
		"LineNumber": query["line_number"].(string),
		"IsRead":     fmt.Sprint(query["is_read"]),
		"PageSize":   fmt.Sprint(query["page_size"]),
		"PageNumber": fmt.Sprint(query["page_number"]),
	}
	queryString := BuildQueryString(g.url+"GetReceivedSmsesPaging", params, "", nil)
	response, err := g.request("GET", queryString, nil)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) GetOtpTemplateParameters(templateName string) (string, error) {
	queryString := fmt.Sprintf("%sGetOtpTemplateParameters?TemplateName=%s", g.url, templateName)
	response, err := g.request("GET", queryString, nil)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) SendSimpleSms(command interface{}) (string, error) {
	body, err := json.Marshal(command)
	if err != nil {
		return "", err
	}
	response, err := g.request("POST", "SendSingleSMS", body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) SendBulkSms(command interface{}) (string, error) {
	body, err := json.Marshal(command)
	if err != nil {
		return "", err
	}
	response, err := g.request("POST", "SendBulkSMS", body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) SendPairToPairSms(command interface{}) (string, error) {
	body, err := json.Marshal(command)
	if err != nil {
		return "", err
	}
	response, err := g.request("POST", "SendPairToPairSMS", body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) SendOtpSmsOld(command interface{}) (string, error) {
	body, err := json.Marshal(command)
	if err != nil {
		return "", err
	}
	response, err := g.request("POST", "SendOtpSMSOld", body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

func (g *Ghasedak) SendOtpSms(command interface{}) (string, error) {
	body, err := json.Marshal(command)
	if err != nil {
		return "", err
	}
	response, err := g.request("POST", "SendOtpSMS", body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
