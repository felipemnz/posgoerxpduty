package httpclient

import (
	"net/http"
)

func MakeRequest(url string) (int, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
