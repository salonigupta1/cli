package commands

import (
	"cli/src/utils"
	"io/ioutil"
	"testing"
)

func TestGetCall(t *testing.T) {
	res, err := utils.GetCall("https://www.google.com/")
	if err != nil {
		t.Errorf("Fail! Not detecting correct URLs.")
	} else if res.Status != "200 OK" {
		t.Errorf("Fail! Wrong response coming")
	}
}

func TestIsReturnTypeJson(t *testing.T) {
	response, _ := utils.GetCall("https://www.google.com/")
	if utils.IsReturnTypeJson(response)  {
		t.Errorf("Fail! Not giving correct response type")
	}
}

func TestIsStatusOk(t *testing.T) {
	response, _ := utils.GetCall("https://www.google.com/")
	if !utils.IsStatusOk(response) {
		t.Errorf("Fail! Not giving correct status result")
	}
}

func TestParseXml(t *testing.T) {
	response, _ := utils.GetCall("http://api.open-notify.org/astros.json")
	data, err := ioutil.ReadAll(response.Body)
	_, err = utils.ParseXml(string(data))
	if err!=nil {
		t.Errorf("Fail! Not parsing XML from JSON")
	}

}

func TestValidateURLResponse(t *testing.T) {
	response, _ := utils.GetCall("https://www.google.com")
	if ValidateURLResponse(response)==nil {
		t.Errorf("Fail! Wrong Validation of URL")
	}

	response, _ = utils.GetCall("http://api.open-notify.org/astros.json")
	if ValidateURLResponse(response)!=nil {
		t.Errorf("Fail! Wrong Validation of URL")
	}
}
