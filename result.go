package main

import (
	"fmt"
	"log"
	"time"
)

type Result struct {
	Code  string
	Count int
	Use   bool
}

func getResult(query string) {
	var Datas = map[string]*Result{}
	hasCounty := map[string]bool{}
	//fofas, err := getFoFaStat(`app="UNIFI-unifi-摄像头"`)
	fofas, err := getFoFaStat(query)
	if err != nil {
		log.Fatalln(err)
	}
	if fofas.Code != 0 {
		log.Printf("err")
	}
	countries := fofas.Data.Countries
	for _, v := range countries {
		if v.Count <= fofaSize {
			Datas[v.Name] = &Result{
				Code:  v.Code.String(),
				Count: v.Count,
				Use:   false,
			}
		} else {
			hasCounty[v.Name] = true
			for _, city := range v.Regions {
				code := fmt.Sprintf("country=\"%s\" && %s", v.NameCode, city.Code.String())
				if v.NameCode == "HK" || v.NameCode == "MO" || v.NameCode == "TW" {
					code = fmt.Sprintf("region=\"%s\" && %s", v.NameCode, city.Code.String())
				}
				Datas[fmt.Sprintf("%s-%s", v.Name, city.Name)] = &Result{
					Code:  code,
					Count: city.Count,
					Use:   false,
				}
			}
		}
	}

	for _, v := range fofas.Data.CountryChartData {
		if _, ok := hasCounty[v.Cname]; ok {
			continue
		}
		if v.Value <= fofaSize {
			Datas[v.Cname] = &Result{
				Code:  v.Code.String(),
				Count: v.Value,
				Use:   false,
			}
		} else {
			time.Sleep(time.Second * 1)
			country, err := getFoFaStat(v.Code.String())
			if err != nil {
				log.Fatalln(err)
				continue
			}
			cty := country.Data.Countries[0]
			for _, city := range cty.Regions {
				code := fmt.Sprintf("country=\"%s\" && %s", cty.NameCode, city.Code.String())
				if cty.NameCode == "HK" || cty.NameCode == "MO" || cty.NameCode == "TW" {
					code = fmt.Sprintf("region=\"%s\" && %s", cty.NameCode, city.Code.String())
				}
				Datas[fmt.Sprintf("%s-%s", v.Name, city.Name)] = &Result{
					Code:  code,
					Use:   false,
					Count: city.Count,
				}
			}

		}
	}

	var results []Result
	for _, v := range Datas {
		//log.Printf("%s,%d\n", v.Code, v.Count)
		if v.Use {
			//log.Printf("use : %s,%d\n", v.Code, v.Count)
			continue
		}
		if v.Count > fofaSize {
			results = append(results, Result{Code: v.Code, Count: fofaSize, Use: true})
			v.Use = true
			continue
		}
		var code = fmt.Sprintf("(%s)", v.Code)
		var count = v.Count
		v.Use = true
		for _, k := range Datas {
			if k.Use {
				continue
			}
			if k.Count+count <= fofaSize {
				code = fmt.Sprintf("%s||(%s)", code, k.Code)
				count = count + k.Count
				k.Use = true
			}
		}

		results = append(results, Result{Code: code, Count: count})
	}
	var total = 0
	for _, v := range results {
		if debug {
			log.Printf("The current number of fofa queries is %d\n", v.Count)
		}
		total += v.Count
		fmt.Printf("%s\n", v.Code)
	}
	log.Printf("\n\nThe number of queries in the %s is  %d\n", query, total)
}
