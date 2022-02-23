package grafana

import (
	//"encoding/json"
	"bytes"
	"fmt"
	//"io/ioutil"
	"strings"
	"net/http"
	"os"
	"time"
)

func postToApi(data string) int {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", "http://localhost:3000/api/dashboards/db", bytes.NewBuffer([]byte(data)))
	if err != nil {
		fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", `Bearer eyJrIjoiM01TUzVJdHAxRHMzTzkzdkVVNVc1WWRnaFhSV1NjUGwiLCJuIjoiZ3JhZmFuYV90b2tsZW4iLCJpZCI6MX0=`)
	resp, err := client.Do(req)
	// TODO: Return POST status to Front-end feedback
	return resp.StatusCode
}

func CreateDashboard(build string, model string, serial string) int {
	fmt.Printf("### createDashboard: Device: %s // %s // %s \n", build, model, serial)
	dat, _ := os.ReadFile(fmt.Sprintf("/home/adrianozdp/workspace/pi_das_2021_2/protocol/grafana/templates/%s.json", model))
	content := string(dat)
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