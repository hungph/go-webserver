package models

import (
	res "../utils"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"strings"
	"time"
)

type RequestUser struct {
	UserName string `json:"user_nm"`
	Password string `json:"passwd"`
}

func (reqUser RequestUser) GetUserName() string {
	return reqUser.UserName
}

func (reqUser RequestUser) GetPasswd() string {
	return reqUser.Password
}

type User struct {
	UserNm    string
	EncPasswd string
	CreDt     int64
}

func ValidateUsername(userName string, passWd string, checkExisted bool) (map[string]interface{}, bool) {
	//Username must not be empty
	if strings.TrimSpace(userName) == "" {
		return res.ResponseEntity(
				res.ErrorConstants.Failed.Code(), res.ErrorConstants.UsernameEmpty.Code(),
				0, res.ErrorConstants.UsernameEmpty.Message()),
			false
	}

	//Password must not be empty
	if passWd == "" {
		return res.ResponseEntity(
				res.ErrorConstants.Failed.Code(), res.ErrorConstants.PasswordEmpty.Code(),
				0, res.ErrorConstants.PasswordEmpty.Message()),
			false
	}

	if checkExisted {
		//Check if username is existed
		dynamoRes, err := GetDynamoDBClient().GetItem(&dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"UserNm": {
					S: aws.String(userName),
				},
			},
			TableName: aws.String(GetTableName()),
		})

		if err != nil {
			return res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.SystemError.Code(), 0,
				err.Error()), false
		}

		dbUser := User{}
		err = dynamodbattribute.UnmarshalMap(dynamoRes.Item, &dbUser)

		if err != nil {
			return res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.SystemError.Code(), 0,
				err.Error()), false
		}

		if dbUser.UserNm != "" {
			return res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.UsernameDuplicate.Code(), 0,
				res.ErrorConstants.UsernameDuplicate.Message()), false
		}
	}

	return res.ResponseEntity(res.ErrorConstants.Success.Code(), "", 0, ""), true
}

func CreateUser(userName string, passWd string) (map[string]interface{}, int) {
	curTime := time.Now().UnixNano() / 1000000

	if resp, ok := ValidateUsername(userName, passWd, true); !ok {
		return resp, 400
	}

	encryptedPass := res.EncryptString(passWd)
	user := User{}
	user.UserNm = userName
	user.EncPasswd = encryptedPass
	user.CreDt = curTime

	marshalMap, err := dynamodbattribute.MarshalMap(user)

	if err != nil {
		return res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.SystemError.Code(), res.TimeDiff(curTime),
			err.Error()), 500
	}

	_, err = GetDynamoDBClient().PutItem(&dynamodb.PutItemInput{
		Item:      marshalMap,
		TableName: aws.String(GetTableName()),
	})

	if err != nil {
		return res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.SystemError.Code(), res.TimeDiff(curTime),
			err.Error()), 500
	}

	return res.ResponseEntity(res.ErrorConstants.Success.Code(), res.ErrorConstants.SignupSuccessfully.Code(), res.TimeDiff(curTime),
		res.ErrorConstants.SignupSuccessfully.Message()), 200
}

func Signin(userName string, passWd string) (map[string]interface{}, int) {
	curTime := time.Now().UnixNano() / 1000000

	if resp, ok := ValidateUsername(userName, passWd, false); !ok {
		return resp, 400
	}

	//Check username
	dynamoRes, err := GetDynamoDBClient().GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserNm": {
				S: aws.String(userName),
			},
		},
		TableName: aws.String(GetTableName()),
	})

	if err != nil {
		return res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.SystemError.Code(), res.TimeDiff(curTime),
			err.Error()), 500
	}

	dbUser := User{}
	err = dynamodbattribute.UnmarshalMap(dynamoRes.Item, &dbUser)

	if dbUser.UserNm == "" {
		return res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.UsernameNotExisted.Code(), res.TimeDiff(curTime),
			res.ErrorConstants.UsernameNotExisted.Message()), 400
	}

	//Compare passwd
	encryptedPass := res.EncryptString(passWd)

	if dbUser.EncPasswd != encryptedPass {
		return res.ResponseEntity(res.ErrorConstants.Failed.Code(), res.ErrorConstants.PasswordNotMatched.Code(), res.TimeDiff(curTime),
			res.ErrorConstants.PasswordNotMatched.Message()), 400
	}

	//Generate Random token
	randToken := res.RandomString()

	return res.ResponseEntity(res.ErrorConstants.Success.Code(), res.ErrorConstants.SignupSuccessfully.Code(), res.TimeDiff(curTime),
		randToken), 200
}

func DeleteUser(userName string) {
	_, err := GetDynamoDBClient().DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"UserNm": {
				S: aws.String(userName),
			},
		},
		TableName: aws.String(GetTableName()),
	})

	if err != nil {
		fmt.Println(err)
	}
}
