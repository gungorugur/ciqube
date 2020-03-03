package sonarqube

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gungorugur/ciqube/pkg/httpclient"
)

type projectStatusResponse struct {
	ProjectStatus ProjectStatus `json:"projectStatus"`
}

//ProjectStatus model
type ProjectStatus struct {
	Conditions []Condition `json:"conditions"`
	Status     string      `json:"status"`
}

//IsPassed returns if quality gate passed or not
func (ps ProjectStatus) IsPassed() bool {
	return ps.Status == "OK"
}

//Condition model
type Condition struct {
	Status         string `json:"status,omitempty"`
	MetricKey      string `json:"metricKey,omitempty"`
	Comparator     string `json:"comparator,omitempty"`
	ErrorThreshold string `json:"errorThreshold,omitempty"`
	ActualValue    string `json:"actualValue,omitempty"`
}

//GetQualityGateProjectStatus  return project status
func GetQualityGateProjectStatus(host string, token string, projectKey string, client httpclient.Client) (*ProjectStatus, error) {

	path := host + "/api/qualitygates/project_status"

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("projectKey", projectKey)
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(token, "")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var projectStatusResponse projectStatusResponse
	if err := json.Unmarshal(body, &projectStatusResponse); err != nil {
		return nil, err
	}

	return &projectStatusResponse.ProjectStatus, nil
}
