package e2e

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func httpRequest(method, body, path string) (string, error) {
	url := fmt.Sprintf("http://localhost:8080/%s", path)

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}

	client := http.Client{}
	request, err := client.Do(req)
	if err != nil {
		return "", err
	}
	all, err := io.ReadAll(request.Body)
	if err != nil {
		return "", err
	}
	defer request.Body.Close()
	return string(all), nil
}
