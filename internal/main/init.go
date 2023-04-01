package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"tgbot/internal/types"
)

func SetCommands(botUrl string, commands types.BotCommands) error {
	buf, err := json.Marshal(commands)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/setMyCommands", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func GetUpdates(botUrl string, offset int) ([]types.Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse types.RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}
