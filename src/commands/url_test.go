package commands

import (
	"cli/src/utils"
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
