package controllers

import (
	"../models"
	res "../utils"
	"encoding/json"
	"net/http"
	"time"
)

var SignUp = func(resWriter http.ResponseWriter, req *http.Request) {
	curTime := time.Now().UnixNano() / 1000000

	resUser := &models.RequestUser{}

	err := json.NewDecoder(req.Body).Decode(resUser)

	if err != nil {
		resEntity := res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.SystemError.Code(), res.TimeDiff(curTime),
			err.Error())
		res.Respond(resWriter, resEntity)

		return
	}

	resp := models.CreateUser(resUser.GetUserName(), resUser.GetPasswd())
	res.Respond(resWriter, resp)
}

var SignIn = func(resWriter http.ResponseWriter, req *http.Request) {
	curTime := time.Now().UnixNano() / 1000000

	resUser := &models.RequestUser{}

	err := json.NewDecoder(req.Body).Decode(resUser)

	if err != nil {
		resEntity := res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.SystemError.Code(), res.TimeDiff(curTime),
			err.Error())
		res.Respond(resWriter, resEntity)

		return
	}

	resp := models.Signin(resUser.GetUserName(), resUser.GetPasswd())
	res.Respond(resWriter, resp)
}
