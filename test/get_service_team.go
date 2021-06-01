package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AlertService struct {
	endpoint string
}

type ResourceService struct {
	endpoint string
}

type ResourceResp struct {
	Code int `json:"code"`
	Data struct {
		Items []map[string]interface{} `json:"items"`
	} `json:"data"`
	Message string `json:"message"`
}

func (res ResourceService) GetServiceList() []map[string]interface{} {
	path := "/api/resource/list"
	req, err := http.NewRequest("GET", res.endpoint+path, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("per_page", "3000")
	q.Add("model", "project_info")
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	respText, _ := ioutil.ReadAll(resp.Body)
	respData := ResourceResp{}
	json.Unmarshal(respText, &respData)
	return respData.Data.Items
}

func (s *AlertService) GetServiceTeam(name string) string {
	serviceInfo := s.GetServiceInfo(name)
	if team, ok := serviceInfo["group_name"]; ok {
		return team.(string)
	}
	return ""
}

func (s *AlertService) GetServiceInfo(name string) (respData map[string]interface{}) {
	path := "/api/services/name"
	req, err := http.NewRequest("GET", s.endpoint+path, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("name", name)
	req.URL.RawQuery = q.Encode()
	var resp *http.Response
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(2)
	}
	respText, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respText, &respData)
	defer resp.Body.Close()
	return
}

func main() {
	resourceService := &ResourceService{endpoint: "http://resource-tree.luojilab.com"}
	serviceList := resourceService.GetServiceList()
	for _, service := range serviceList {
		team := service["team"].(string)
		serviceName := service["name"].(string)
		owner := service["owner_name"].(string)
		fmt.Printf("%s\t%s\t%s\t\n", serviceName, team, owner)
	}
}
