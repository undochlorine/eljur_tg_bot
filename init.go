package main

import (
	"net/http"
	"encoding/json"
	"bytes"
	"strconv"
	"io/ioutil"
)

func SetCommands(botUrl string, commands BotCommands) (error) {
	buf, err := json.Marshal(commands)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl + "/setMyCommands", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func GetUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}