package sonarqube

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

type mockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (c mockClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}

func TestGetQualityGateProjectStatus(t *testing.T) {

	host := "http://sonarqube.test.com"
	token := "token"
	projectKey := "test-api"

	mockDoFunc := func(req *http.Request) (*http.Response, error) {

		if req.Method != "GET" {
			t.Errorf("Method should be get")
		}

		if req.URL.Path != "/api/qualitygates/project_status" {
			t.Errorf("Path should be quality gate project status path")
		}

		expectedRawQuery := "projectKey=" + projectKey

		if req.URL.RawQuery != expectedRawQuery {
			t.Errorf("Current raw query: %s but want %s", req.URL.RawQuery, expectedRawQuery)
		}

		expectedAuthorizationHeader := "Basic dG9rZW46"

		if req.Header.Get("Authorization") != expectedAuthorizationHeader {
			t.Errorf("Curren authorization header: %s but want %s", req.Header.Get("Authorization"), expectedAuthorizationHeader)
		}

		json := `{   
			"projectStatus":{      
			  "status":"OK",
			  "conditions":[         
					{
					"status":"OK",
					"metricKey":"duplicated_lines_density",
					"comparator":"GT",
					"errorThreshold":"3",
					"actualValue":"1.7"
					}
				]
			}
		}`

		response := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString(json)),
		}

		return &response, nil
	}

	mockClient := mockClient{DoFunc: mockDoFunc}

	response, _ := GetQualityGateProjectStatus(host, token, projectKey, mockClient)

	if response.Status != "OK" {
		t.Errorf(response.Status)
	}

}
