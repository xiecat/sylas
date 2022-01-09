package main

import (
	"encoding/base64"
	"log"
)

type b64str string

func (b b64str) String() string {
	bckd, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		log.Print(err)
	}
	return string(bckd)
}

type common struct {
	Code  b64str `json:"code"`
	Count int    `json:"count"`
	Name  string `json:"name"`
}

type Country struct {
	Code     b64str   `json:"code"`
	Count    int      `json:"count"`
	Ename    string   `json:"ename"`
	Name     string   `json:"name"`
	NameCode string   `json:"name_code"`
	Regions  []common `json:"regions"`
}
type CountryChartData struct {
	Cname string `json:"cname"`
	Code  b64str `json:"code"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type List struct {
	hash     []common `json:"hash"`
	port     []common `json:"port"`
	protocol []common `json:"protocol"`
	server   []common `json:"server"`
	title    []common `json:"title"`
}
type Ranks struct {
	List List `json:"List"`
}

type Data struct {
	Countries        []Country          `json:"countries"`
	CountryChartData []CountryChartData `json:"country_chart_data"`
	DistinctIps      int                `json:"distinct_ips"`
	Query            string             `json:"q"`
	Ranks            Ranks              `json:"ranks"`
}

type FoFaStatusResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    Data   `json:"data"`
}
