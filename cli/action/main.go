package action

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type responseJson struct {
	Value string
}

func storeURL(store, key string) string {
	var u url.URL
	u.Scheme = "https"
	u.Host = store
	u.Path = "/v1/kv/" + key
	return u.String()
}

func request(method, url string, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return &http.Response{}, err
	}

	c := &http.Client{}
	return c.Do(req)
}

func getValueFromConsulBody(body []byte) (string, error) {
	var consulResponse []responseJson
	err := json.Unmarshal(body, &consulResponse)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("json unmarshal error %s", err))
	}
	value, err := base64.StdEncoding.DecodeString(consulResponse[0].Value)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("json unmarshal error %s", err))
	}
	return string(value), nil
}

func keyValueRequest(method, url, data string) ([]byte, error) {
	resp, err := request(method, url, data)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode >= 300 {
		return []byte{}, fmt.Errorf(fmt.Sprintf("response status '%s'", resp.Status))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf(fmt.Sprintf("response body %s", err))
	}
	return body, nil
}

func GetKey(store string, key string) (string, error) {
	url := storeURL(store, key)
	body, err := keyValueRequest(http.MethodGet, url, key)
	if err != nil {
		return "", err
	}
	return getValueFromConsulBody(body)
}

func PutKey(store, key, value string) error {
	url := storeURL(store, key)
	_, err := keyValueRequest(http.MethodPut, url, value)
	if err != nil {
		return err
	}
	return nil
}

func DeleteKey(store, key string) error {
	fmt.Println("deleting")
	url := storeURL(store, key)

	body, err := keyValueRequest(http.MethodDelete, url, key)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
