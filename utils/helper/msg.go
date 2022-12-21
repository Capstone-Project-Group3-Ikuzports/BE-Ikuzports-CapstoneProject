package helper

import (
	"errors"
	"log"
	"strings"
)

func ServiceErrorMsg(errData error) error {
	if strings.Contains(errData.Error(), "table") {
		return errors.New("Failed. Error on process. Please contact your administrator.")
	} else if strings.Contains(errData.Error(), "found") {
		return errors.New("Failed. Data not found. Please check input again.")
	} else if strings.Contains(errData.Error(), "failed on the 'required' tag") {
		return errors.New("Failed. Required field is empty. Please check input again.")
	} else if strings.Contains(errData.Error(), "foreign key constraint fails ") {
		return errors.New("Failed. Reference ID not found. Please check input again.")
	} else {
		return errors.New("Failed. Other Error. Please contact your administrator.")
	}
}

func HandlerErrorMsg(errData error) error {
	if strings.Contains(errData.Error(), "table") {
		return errors.New("Failed. Error on process. Please contact your administrator.")
	} else if strings.Contains(errData.Error(), "found") {
		return errors.New("Failed. Data not found. Please check input again.")
	} else {
		return errors.New("Failed. Other Error. Please contact your administrator.")
	}
}

func LogDebug(msg ...interface{}) {
	log.Println(msg)
}
