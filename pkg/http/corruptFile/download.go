package corruptFile

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func GetFile(botToken, fileId string) ([]byte, error) {
	tgFilePath, err := getFilePath(botToken, fileId)
	if err != nil {
		return []byte{}, err
	}
	fileRes, err := downloadFile(botToken, tgFilePath)
	return fileRes, nil
}

func getFilePath(botToken, fileId string) (string, error) {
	// Example: https://api.telegram.org/bot+<botToken>/getFile?file_id=<fileId>
	baseURL, _ := url.Parse("https://api.telegram.org/bot" + botToken + "/getFile")
	endpoint := "getFile"
	params := url.Values{}
	params.Add("file_id", fileId)
	resolvedURL := baseURL.ResolveReference(&url.URL{Path: endpoint}).String() + "?" + params.Encode()

	resp, err := http.Get(resolvedURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type getFileType struct {
		Ok     bool `json:"ok"`
		Result struct {
			FilePath string `json:"file_path"`
		} `json:"result"`
	}
	var restResponse getFileType
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return "", err
	}
	if !restResponse.Ok {
		return "", err
	}

	return restResponse.Result.FilePath, nil
}

func downloadFile(botToken, filePath string) ([]byte, error) {
	// Example: https://api.telegram.org/file/bot+<botToken>/<filePath>
	resolvedURL := "https://api.telegram.org/file/bot" + botToken + "/" + filePath
	resp, err := http.Get(resolvedURL)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	fileData, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return fileData, nil
}
