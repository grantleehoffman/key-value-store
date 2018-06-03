package action

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRequester interface {
	request() (*http.Response, error)
}
type mockRequestClient struct {
	body       io.ReadCloser
	statusCode int
}

func (r mockRequestClient) request() (*http.Response, error) {
	return &http.Response{Body: r.body, StatusCode: r.statusCode}, nil
}

func TestStoreUrl(t *testing.T) {
	receivedStoreURI := storeURL("kvstore.com", "my-test-key")
	assert.Equal(t, "https://kvstore.com/v1/kv/my-test-key", receivedStoreURI)
}

func TestKeyValueRequestsStatusCodes(t *testing.T) {
	var testCases = []struct {
		testCase      string
		statusCode    int
		body          []byte
		expectedError bool
	}{

		{
			testCase:      "200 status code and returned body test case",
			statusCode:    200,
			body:          []byte("testValue"),
			expectedError: false,
		},
		{
			testCase:      "400 status code test case",
			statusCode:    404,
			body:          []byte{},
			expectedError: true,
		},
		{
			testCase:      "503 status code test case",
			statusCode:    503,
			body:          []byte{},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		m := mockRequestClient{ioutil.NopCloser(bytes.NewBuffer(tc.body)), tc.statusCode}
		body, err := keyValueRequest(m)
		retErr := err != nil

		assert.Equal(t, tc.expectedError, retErr, tc.testCase)
		assert.Equal(t, tc.body, body, tc.testCase)
	}
}

func TestGetValueFromConsulBody(t *testing.T) {
	var testCases = []struct {
		testCase      string
		body          string
		expectedValue string
		expectedError bool
	}{

		{
			testCase: "proper values test case",
			body: `[{"LockIndex":0,
			                "Key":"test",
					"Flags":0,
					"Value":"dGVzdHZhbHVl",
					"CreateIndex":254,
					"ModifyIndex":254}]`,
			expectedValue: "testvalue",
			expectedError: false,
		},
		{
			testCase: "bad base64 test case",
			body: `[{"LockIndex":0,
			                "Key":"test",
					"Flags":0,
					"Value":"/$@",
					"CreateIndex":254,
					"ModifyIndex":254}]`,
			expectedValue: "",
			expectedError: true,
		},
		{
			testCase: "bad json test case",
			body: `[{{{{{"LockIndex":0,
			                "Key":"test",
					"Flags":0,
					"Value":"dGVzdHZhbHVl",
					"CreateIndex":254,
					"ModifyIndex":254i]`,
			expectedValue: "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		value, err := getValueFromConsulBody([]byte(tc.body))
		retErr := err != nil

		assert.Equal(t, tc.expectedError, retErr, tc.testCase)
		assert.Equal(t, tc.expectedValue, value, tc.testCase)
	}

}
