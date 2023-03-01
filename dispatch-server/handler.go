package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestPayload struct {
	Action   string      `json:"action" binding:"required"`
	TODOdata todoPayload `json:"action,omitempty" binding:"required"`
}

type todoPayload struct {
	Title       string `json:"title" binding:"required"`
	ID          string `json:"id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"`
	IsDone      bool   `json:"is_done"`
	IsDelete    bool   `json:"is_delete"`
	IsFavorite  bool   `json:"is_favorite"`
}

func (app *Config) submitHandler(ctx *gin.Context) {
	var req RequestPayload

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("caanot binding")
		return
	}

	switch req.Action {
	case "request-data":
		app.requestData(&req.TODOdata)
	}
}

func (app *Config) requestData(payload *todoPayload) {
	//轉發data到另一個server
	//發送請求到其他server 必須用net/http

	reqBody, err := json.Marshal(&payload)
	if err != nil {
		log.Println("cannot marshal request body")
		return
	}

	reqURL := "http://localhost:9090/pending/add-task"
	reqHeaders := make(http.Header)
	reqHeaders.Set("Content-Type", "application/json")

	// 发送 HTTP 请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", reqURL, bytes.NewReader(reqBody))
	if err != nil {
		log.Println("cannot create http request")
		return
	}

	req.Header = reqHeaders
	res, err := client.Do(req)
	if err != nil {
		log.Println("cannot send http request")
		return
	}
	defer res.Body.Close()

	// 处理 HTTP 响应
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("cannot read http response")
		return
	}
	log.Println(string(resBody))
}
