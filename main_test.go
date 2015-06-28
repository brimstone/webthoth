package main

import (
	"net/http/httptest"
	"testing"

	"github.com/brimstone/go-saverequest"
)

func Test_Hello(t *testing.T) {
	// need to reset root
	t.Log("Testing Hello")
	req, _ := saverequest.FakeRequest("GET", "/", map[string]string{}, "")
	w := httptest.NewRecorder()
	helloFunc(w, req)
	if w.Body.String() != "hello!\n" {
		t.Errorf("Got unexpected hello")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
	t.Log("Got Proper hello")
}

func Test_NullRoom(t *testing.T) {
	t.Log("Testing Null Room")
	req, _ := saverequest.FakeRequest("GET", "/rooms/", map[string]string{}, "")
	w := httptest.NewRecorder()
	roomFunc(w, req)
	if w.Body.String() != "" {
		t.Errorf("Unable to get null room")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
	t.Log("Got null room")
}

func Test_SingleOffer(t *testing.T) {
	t.Log("Looking for empty room")
	req, _ := saverequest.FakeRequest("GET", "/rooms/asdf", map[string]string{}, "")
	w := httptest.NewRecorder()
	roomFunc(w, req)
	if w.Body.String() != "[]" {
		t.Errorf("Unable to get empty room")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	} else {
		t.Log("Got empty room")
	}

	offerDescription := "This is my offer"
	t.Log("Trying to save a description")
	req, _ = saverequest.FakeRequest("POST", "/rooms/asdf", map[string]string{}, offerDescription)
	w = httptest.NewRecorder()
	roomFunc(w, req)
	descKey := w.Body.String()
	if descKey == "" {
		t.Errorf("Unable to save description")
		t.Errorf("%d: %s", w.Code, descKey)
		return
	}

	t.Log("Trying to retrieve description")
	req, _ = saverequest.FakeRequest("GET", "/rooms/asdf/"+descKey, map[string]string{}, "")
	w = httptest.NewRecorder()
	roomFunc(w, req)
	if w.Body.String() != offerDescription {
		t.Errorf("Unable to retrieve description")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}

	answerDescription := "This is my answer"
	t.Log("Trying to answer description")
	req, _ = saverequest.FakeRequest("POST", "/rooms/asdf/"+descKey, map[string]string{}, answerDescription)
	w = httptest.NewRecorder()
	roomFunc(w, req)
	if w.Body.String() != "" {
		t.Errorf("Unable to answer description")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}

	t.Log("Trying to get answer description")
	req, _ = saverequest.FakeRequest("GET", "/rooms/asdf/"+descKey, map[string]string{}, "")
	w = httptest.NewRecorder()
	roomFunc(w, req)
	if w.Body.String() != answerDescription {
		t.Errorf("Unable to answer description")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}

	t.Log("Trying to get answer description, again")
	req, _ = saverequest.FakeRequest("GET", "/rooms/asdf/"+descKey, map[string]string{}, "")
	w = httptest.NewRecorder()
	roomFunc(w, req)
	if w.Code != 404 {
		t.Errorf("Able to answer description")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
}

func Test_GetNonOffer(t *testing.T) {
	t.Log("Trying to retrieve nonexistant offer description")
	req, _ := saverequest.FakeRequest("GET", "/rooms/asdf/asdf", map[string]string{}, "")
	w := httptest.NewRecorder()
	roomFunc(w, req)
	if w.Code != 404 {
		t.Errorf("Able to retrieve nonexistant offer description")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
}
func Test_AnswerNonexistantOffer(t *testing.T) {
	t.Log("Trying to answer description")
	req, _ := saverequest.FakeRequest("POST", "/rooms/nope/asdf", map[string]string{}, "asdf")
	w := httptest.NewRecorder()
	roomFunc(w, req)
	if w.Code == 200 {
		t.Errorf("Able to answer fake description")
		t.Errorf("%d: %s", w.Code, w.Body.String())
		return
	}
}
