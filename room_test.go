package main

import (
	"fmt"
	"testing"
)

func Test_TestDescription(t *testing.T) {
	r := NewRoom()
	offerKey, _ := r.NewDescription("pickles")
	offerResponse, _ := r.GetDescription(offerKey)

	fmt.Println(r.ListDescriptions())
	if offerResponse != "pickles" {
		t.Error("Description isn't pickles")
	}
	r.AnswerDescription(offerKey, "bananas")
	answer, err := r.GetDescription(offerKey)
	if answer != "bananas" {
		t.Error("Description not updated with answer", answer)
	}
	if err != nil {
		t.Error(err.Error())
	}
	if len(r.ListDescriptions()) != 0 {
		t.Error("Description still stands!", r.ListDescriptions())
	}
}

func Test_TestDescriptionNotFound(t *testing.T) {
	r := NewRoom()
	_, err := r.GetDescription("asdf")
	if err == nil {
		t.Error("error should be set")
	}
}
