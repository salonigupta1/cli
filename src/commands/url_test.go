package commands

import (
	"cli/src/utils"
	"testing"
)



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

func TestCallFunction(t *testing.T) {
	response := CallFunction("https://www.google.com/")
	if response == nil {
		t.Errorf("Fail, Wrong Results")
	}

	response = CallFunction("http://api.open-notify.org/astros.json")
	if response!=nil {
		t.Errorf("Fail, Worong Results")
	}
}
