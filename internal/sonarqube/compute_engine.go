package sonarqube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gungorugur/ciqube/pkg/httpclient"
)

type computeEngineComponent struct {
	Queue []queue `json:"queue"`
}

type queue struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	SubmittedAt string `json:"submittedAt"`
}

func (c computeEngineComponent) printQueue() {

	for _, element := range c.Queue {
		fmt.Printf("Background task in queue [Id:%v, Status:%v, SubmittedAt:%v]\n", element.ID, element.Status, element.SubmittedAt)
	}
}

//CheckBackgroundTask checks if given project key has a background task in compute engine's queue
func CheckBackgroundTask(host string, token string, projectKey string, timeout int, client httpclient.Client) error {

	path := host + "/api/ce/component"

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("component", projectKey)
	req.URL.RawQuery = q.Encode()
	req.SetBasicAuth(token, "")

	startTime := time.Now()

	for {
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var computeEngineComponent computeEngineComponent
		if err := json.Unmarshal(body, &computeEngineComponent); err != nil {
			return err
		}

		currentTime := time.Now()
		diffInSeconds := currentTime.Sub(startTime).Seconds()

		if len(computeEngineComponent.Queue) > 0 {
			computeEngineComponent.printQueue()
		} else {
			return nil
		}

		if int(diffInSeconds) > timeout {
			return errors.New("background task timeout expired")
		}

		time.Sleep(5 * time.Second)
	}
}
