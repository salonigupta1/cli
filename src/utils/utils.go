package utils

import (
	"cli/src/constants"
	_ "cli/src/constants"
	"encoding/xml"
	"github.com/golang/gddo/httputil/header"
	"net/http"
)

func GetCall(url string) (*http.Response, error) {
	response, err := http.Get(url)
	return response, err
}

func IsReturnTypeJson(response *http.Response) bool {
	if response.Header.Get(constants.ContentType) != "" {
		value, _ := header.ParseValueAndParams(response.Header, constants.ContentType)
		return value == constants.JsonKey
	}
	return false
}

func IsStatusOk(response *http.Response) bool {
	return response.StatusCode == http.StatusOK
}

func ParseXml(data string) (string, error) {
	x, err := xml.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(x), nil
}
