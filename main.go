package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vitormoschetta/go-gpt/models"
)

func main() {
	appSettings := loadAppSettings()

	requestBody, err := json.Marshal(appSettings.CompletionRequest)
	request, err := http.NewRequest("POST", appSettings.Url, bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+appSettings.Token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var completionResponse models.CompletionResponse
	err = json.Unmarshal(responseBody, &completionResponse)
	if err != nil {
		panic(err)
	}

	println(completionResponse.Choices[0].Text)
}

func loadAppSettings() models.AppSettings {
	fileData, err := ioutil.ReadFile("appsettings.json")
	if err != nil {
		panic(err)
	}

	var appSettings models.AppSettings
	err = json.Unmarshal(fileData, &appSettings)
	if err != nil {
		panic(err)
	}

	return appSettings
}
