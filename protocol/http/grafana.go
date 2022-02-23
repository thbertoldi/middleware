package grafana

import (
	"bytes"
	"fmt"
	"strings"
	"net/http"
	"time"
	"middleware/config"
)

func postToApi(data string) int {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", "http://localhost:3000/api/dashboards/db", bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Printf("Got error %s", err.Error())
		return 400
	}
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", `Bearer eyJrIjoiM01TUzVJdHAxRHMzTzkzdkVVNVc1WWRnaFhSV1NjUGwiLCJuIjoiZ3JhZmFuYV90b2tsZW4iLCJpZCI6MX0=`)
	resp, err := client.Do(req)
	// TODO: Return POST status to Front-end feedback
	return resp.StatusCode
}

func CreateDashboard(build string, model string, serial string, conf config.Internal) int {
	fmt.Printf("### createDashboard: Device: %s // %s // %s \n", build, model, serial)
	var content string
	if model == "tf" {
		content = conf.Template.Tf
	} else if model == "mf" {
		content = conf.Template.Mf
	} else if model == "cm" {
		content = conf.Template.Cm
	}
	content = strings.Replace(content, "#DEVMODEL#", model, -1)
	content = strings.Replace(content, "#DEVMODEL_UPPER#", strings.ToUpper(model), -1)
	content = strings.Replace(content, "#SERIAL#", serial, -1)
	content = strings.Replace(content, "#BUILD#", build, -1)
	ret := postToApi(content)
	return ret
}

func main() {
	fmt.Printf("### Grafana main")
}
