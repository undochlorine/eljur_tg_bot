package send

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

func Document(botUrl string, chatId int, fileData []byte, filename string) error {
	// Create a new buffer to write the multipart/form-data request body
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Create a new form field with the chat ID and write it to the request body
	err := writer.WriteField("chat_id", strconv.Itoa(chatId))
	if err != nil {
		return err
	}

	// Create a new form field with the file contents and write it to the request body
	fileWriter, err := writer.CreateFormFile("document", filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(fileWriter, bytes.NewReader(fileData))
	if err != nil {
		return err
	}

	// Close the multipart writer to finalize the request body
	writer.Close()

	req, err := http.NewRequest("POST", botUrl+"/sendDocument", &buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Content-Length", strconv.Itoa(buf.Len()))

	sent := make(chan struct{})
	go WaitMock(botUrl, chatId, sent)

	// send the HTTP request to the telegram api
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		sent <- struct{}{}
		close(sent)
		resp.Body.Close()
	}()

	return nil
}
