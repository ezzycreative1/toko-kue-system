package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var client = &http.Client{}
var headers = map[string]string{
	"Content-Type": "application/json",
}

// GetClient to get client instance
func GetClient() *http.Client {
	return client
}

// Get do get http request
func Get(url string, body []byte) ([]byte, error) {
	return doRequest("GET", url, body)
}

// Post do post http request
func Post(url string, body []byte) ([]byte, error) {
	return doRequest("POST", url, body)
}

func doRequest(method string, url string, body []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Atur timeout sesuai kebutuhan
	defer cancel()

	payload := strings.NewReader(string(body))

	req, err := http.NewRequestWithContext(ctx, method, url, payload)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if isStatusError(resp.StatusCode) {
		return nil, fmt.Errorf("status error: %v %v", resp.StatusCode, string(respBody))
	}
	return respBody, err
}

func isStatusError(statusCode int) bool {
	return statusCode > http.StatusBadRequest
}

// PostRequest ...
func PostRequest(urls string, headers map[string]string, body []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Atur timeout sesuai kebutuhan
	defer cancel()

	client := &http.Client{}
	payload := bytes.NewReader(body)

	req, err := http.NewRequestWithContext(ctx, "POST", urls, payload)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if isStatusError(resp.StatusCode) {
		return nil, fmt.Errorf("status error: %v %v", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// DoAuthPostRequest ..
func DoAuthPostRequest(URL string, clientID string, clientSecret string) ([]byte, error) {

	client := &http.Client{}
	// body
	data := url.Values{}
	resource := "grant_type=" + "client_credentials"

	urlStr := URL + resource

	data.Add("client_id", clientID)
	data.Add("client_secret", clientSecret)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	// Set Headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	resBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	return resBody, err
}
