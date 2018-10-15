package aws

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Jeffail/gabs"
)

type AWSHandler struct {
	InstanceId string
}

const (
	instanceIdMetaURL = "http://169.254.169.254/latest/meta-data/instance-id"
	ecsTasksURL       = "http://localhost:51678/v1/tasks"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewAWSHandler() *AWSHandler {
	return &AWSHandler{
		InstanceId: getInstanceId(),
	}
}

func getInstanceId() string {
	instanceId, _ := os.Hostname()

	var httpClient = &http.Client{
		Timeout: 5 * time.Second,
	}
	response, err := httpClient.Get(instanceIdMetaURL)

	if timeout, ok := err.(*url.Error); ok && timeout.Timeout() {
		log.Println("Most likely not running in EC2 instance, falling back to hostname!")
	} else {
		check(err)

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		check(err)

		instanceId = string(contents)
	}

	return instanceId
}

func (a *AWSHandler) GetEcsTasksOnInstance() [][]string {
	var tasksInfo [][]string
	var httpClient = &http.Client{
		Timeout: 1 * time.Second,
	}

	response, err := httpClient.Get(ecsTasksURL)

	if timeout, ok := err.(*url.Error); ok && timeout.Timeout() {
		log.Println("Task list fetch timed out, am I running in ECS Container Instance?")
	} else {
		check(err)
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		check(err)
		jsonParsed, err := gabs.ParseJSON(contents)
		tasks, _ := jsonParsed.S("Tasks").Children()
		for _, task := range tasks {
			taskArn := task.Path("Arn").Data().(string)
			serviceName := task.Path("Family").Data().(string)
			parts := strings.Split(taskArn, "/")
			taskID := parts[len(parts)-1]
			tasksInfo = append(tasksInfo, []string{taskID, serviceName})
		}
	}
	return tasksInfo
}
