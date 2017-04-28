package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

//MatrixConfig stores the data from json matrixConfig.json file
type MatrixConfig struct {
	RoomId   string `json:"room_id"`
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
}

//MatrixToken stores the token data from matrix
type MatrixToken struct {
	AccessToken string `json:"access_token"`
	Server      string `json:"server"`
	UserId      string `json:"user_id"`
	DeviceId    string `json:"device_id"`
}

var matrixConfig MatrixConfig
var matrixToken MatrixToken

func readMatrixConfig() {
	file, e := ioutil.ReadFile("matrixConfig.json")
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &matrixConfig)
}

func loginMatrix() {
	url := matrixConfig.Server + "/_matrix/client/r0/login"
	jsonStr := `{
		"type":"m.login.password",
		"user":"` + matrixConfig.User + `",
		"password":"` + matrixConfig.Password + `"
	}`
	b := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("POST", url, b)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal([]byte(body), &matrixToken)

}
func matrixSendMsg(senderScreenName string, message string, createdAt string) {
	txnId := strconv.Itoa(rand.Int())
	c.Green(txnId)
	url := matrixConfig.Server + "/_matrix/client/r0/rooms/" + matrixConfig.RoomId + "/send/m.room.message/" + txnId + "?access_token=" + matrixToken.AccessToken
	jsonStr := `{
		"body":"[NEW DM] - at ` + createdAt + `\n@` + senderScreenName + `: ` + message + `",
		"msgtype":"m.text"
	}`
	b := strings.NewReader(jsonStr)
	req, _ := http.NewRequest("PUT", url, b)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	fmt.Println(createdAt)
	fmt.Print("received DM sent to Matrix: ")
	c.Green(message)

}
