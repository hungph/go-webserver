package main

import (
	"./apps"
	"./models"
	"./utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var curApp apps.App

func TestMain(m *testing.M) {
	curApp = apps.App{}

	curApp.Initialize("test")

	port := os.Getenv("port")

	if port == "" {
		port = "8080"
	}

	fmt.Print(port)

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	curApp.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestSignUpSuccessfully(t *testing.T) {
	payload := []byte(`{"user_nm": "hung", "passwd": "hung12222"}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-up", bytes.NewBuffer(payload))
	response := executeRequest(req)

	models.DeleteUser("hung")

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestSignUpFailDuplicate(t *testing.T) {
	payload := []byte(`{"user_nm": "hung", "passwd": "hung12222"}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-up", bytes.NewBuffer(payload))
	response := executeRequest(req)

	req, _ = http.NewRequest("POST", "/v1/user/sign-up", bytes.NewBuffer(payload))
	response = executeRequest(req)

	models.DeleteUser("hung")

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if m["code"] != utils.ErrorConstants.UsernameDuplicate.Code() {
		t.Errorf("Expected response code to be '%v'. Got '%v'", utils.ErrorConstants.UsernameDuplicate.Code(), m["code"])
	}
}

func TestSignUpFailUsernameEmpty(t *testing.T) {
	payload := []byte(`{"user_nm": "", "passwd": "hung12222"}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-up", bytes.NewBuffer(payload))
	response := executeRequest(req)

	models.DeleteUser("hung")

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if m["code"] != utils.ErrorConstants.UsernameEmpty.Code() {
		t.Errorf("Expected response code to be '%v'. Got '%v'", utils.ErrorConstants.UsernameEmpty.Code(), m["code"])
	}
}

func TestSignUpFailPasswordEmpty(t *testing.T) {
	payload := []byte(`{"user_nm": "hung", "passwd": ""}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-up", bytes.NewBuffer(payload))
	response := executeRequest(req)

	models.DeleteUser("hung")

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if m["code"] != utils.ErrorConstants.PasswordEmpty.Code() {
		t.Errorf("Expected response code to be '%v'. Got '%v'", utils.ErrorConstants.PasswordEmpty.Code(), m["code"])
	}
}

func TestSignInSuccessfully(t *testing.T) {
	payload := []byte(`{"user_nm": "hung", "passwd": "hung12222"}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-up", bytes.NewBuffer(payload))
	response := executeRequest(req)

	req, _ = http.NewRequest("POST", "/v1/user/sign-in", bytes.NewBuffer(payload))
	response = executeRequest(req)

	models.DeleteUser("hung")

	checkResponseCode(t, http.StatusOK, response.Code)

}

func TestSignInFailUsernameEmpty(t *testing.T) {
	payload := []byte(`{"user_nm": "", "passwd": "hung12222"}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-in", bytes.NewBuffer(payload))
	response := executeRequest(req)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if m["code"] != utils.ErrorConstants.UsernameEmpty.Code() {
		t.Errorf("Expected response code to be '%v'. Got '%v'", utils.ErrorConstants.UsernameEmpty.Code(), m["code"])
	}

}

func TestSignInFailPasswordEmpty(t *testing.T) {
	payload := []byte(`{"user_nm": "hung", "passwd": ""}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-in", bytes.NewBuffer(payload))
	response := executeRequest(req)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if m["code"] != utils.ErrorConstants.PasswordEmpty.Code() {
		t.Errorf("Expected response code to be '%v'. Got '%v'", utils.ErrorConstants.PasswordEmpty.Code(), m["code"])
	}

}

func TestSignInFailUserNotExisted(t *testing.T) {
	payload := []byte(`{"user_nm": "hung", "passwd": "hung11111"}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-in", bytes.NewBuffer(payload))
	response := executeRequest(req)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if m["code"] != utils.ErrorConstants.UsernameNotExisted.Code() {
		t.Errorf("Expected response code to be '%v'. Got '%v'", utils.ErrorConstants.UsernameNotExisted.Code(), m["code"])
	}

}

func TestSignInFailPasswordNotMatched(t *testing.T) {
	payload := []byte(`{"user_nm": "hung", "passwd": "hung12222"}`)

	req, _ := http.NewRequest("POST", "/v1/user/sign-up", bytes.NewBuffer(payload))
	response := executeRequest(req)

	payload = []byte(`{"user_nm": "hung", "passwd": "hung122"}`)

	req, _ = http.NewRequest("POST", "/v1/user/sign-in", bytes.NewBuffer(payload))
	response = executeRequest(req)

	models.DeleteUser("hung")

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	checkResponseCode(t, http.StatusBadRequest, response.Code)

	if m["code"] != utils.ErrorConstants.PasswordNotMatched.Code() {
		t.Errorf("Expected response code to be '%v'. Got '%v'", utils.ErrorConstants.PasswordNotMatched.Code(), m["code"])
	}
}
