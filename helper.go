package ghasedak

import (
	"net/url"
	"strings"
)

func BuildQueryString(baseUrl string, params map[string]string, arrayParamName string, arrayParams []string) string {
	query := url.Values{}
	for key, value := range params {
		query.Add(key, value)
	}

	for _, value := range arrayParams {
		query.Add(arrayParamName, value)
	}

	queryString := query.Encode()
	return baseUrl + "?" + strings.TrimSuffix(queryString, "&")
}
