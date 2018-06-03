package action

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type requester interface {
	request() (*http.Response, error)
}

type requestClient struct {
	req *http.Request
}

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

func newRequest(r *http.Request) requester {
	return requestClient{r}
}

func (r requestClient) request() (*http.Response, error) {
	c := &http.Client{}
	return c.Do(r.req)
}

func getValueFromConsulBody(body []byte) (string, error) {
	var consulResponse []responseJson
	err := json.Unmarshal(body, &consulResponse)
	if err != nil {
		return "", errors.New(fmt.Sprintf("json unmarshal error %s", err))
	}
	value, err := base64.StdEncoding.DecodeString(consulResponse[0].Value)
	if err != nil {
		return "", errors.New(fmt.Sprintf("base64 decode error %s", err))
	}
	return string(value), nil
}

func keyValueRequest(r requester) ([]byte, error) {
	resp, err := r.request()
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode >= 300 {
		return []byte{}, errors.New(fmt.Sprintf("response status '%s'", resp.Status))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("response body %s", err))
	}
	return body, nil
}

func GetKey(store string, key string) (string, error) {
	url := storeURL(store, key)
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewReader([]byte(key)))
	if err != nil {
		return "", err
	}
	body, err := keyValueRequest(newRequest(req))
	if err != nil {
		return "", err
	}
	return getValueFromConsulBody(body)
}

func PutKey(store, key, value string) error {
	url := storeURL(store, key)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader([]byte(value)))
	if err != nil {
		return err
	}
	_, err = keyValueRequest(newRequest(req))
	return err
}

func DeleteKey(store, key string) error {
	url := storeURL(store, key)

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewReader([]byte(key)))
	if err != nil {
		return err
	}
	_, err = keyValueRequest(newRequest(req))

	return err
}
