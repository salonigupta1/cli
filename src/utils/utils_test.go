package utils


import (

"io/ioutil"
"testing"
)

func TestGetCall(t *testing.T) {
	res, err := GetCall("https://www.google.com/")
	if err != nil {
		t.Errorf("Fail! Not detecting correct URLs.")
	} else if res.Status != "200 OK" {
		t.Errorf("Fail! Wrong response coming")
	}
}

func TestIsReturnTypeJson(t *testing.T) {
	response, _ := GetCall("https://www.google.com/")
	if IsReturnTypeJson(response)  {
		t.Errorf("Fail! Not giving correct response type")
	}
}

func TestIsStatusOk(t *testing.T) {
	response, _ := GetCall("https://www.google.com/")
	if !IsStatusOk(response) {
		t.Errorf("Fail! Not giving correct status result")
	}
}

func TestParseXml(t *testing.T) {
	response, _ := GetCall("http://api.open-notify.org/astros.json")
	data, err := ioutil.ReadAll(response.Body)
	_, err = ParseXml(string(data))
	if err!=nil {
		t.Errorf("Fail! Not parsing XML from JSON")
	}

}
