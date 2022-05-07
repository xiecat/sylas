package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func getFoFaStat(query string) (result FoFaStatusResp, err error) {
	qbase := base64.StdEncoding.EncodeToString([]byte(query))
	log.Printf("query:%s,base64:%s", query, qbase)
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if chromePath != "" {
		chromedp.ExecPath(chromePath)
	}
	done := make(chan bool)
	var requestID network.RequestID

	chromedp.ListenTarget(ctx, func(v interface{}) {
		switch ev := v.(type) {
		case *network.EventRequestWillBeSent:

			if strings.Contains(ev.Request.URL, fmt.Sprintf("https://api.%s/v1/search/stats?qbase64=", fofaDomain)) && ev.Type == "XHR" {
				if debug {
					log.Printf("EventRequestWillBeSent: %v: %v\n", ev.RequestID, ev.Request.URL)
				} else {
					log.Printf("request fofa query :%s stat data\n", query)
				}

				requestID = ev.RequestID

			}
		case *network.EventLoadingFinished:

			if ev.RequestID == requestID {
				if debug {
					log.Printf("EventLoadingFinished: %v\n", ev.RequestID)
				} else {
					log.Printf("get fofa query :%s stat data\n", query)
				}

				close(done)
			}
		}
	})
	// run task list
	fofaUrl := fmt.Sprintf(`https://%s/result?qbase64=%s`, fofaDomain, qbase)
	log.Printf("fofa Query URL:%s \n", fofaUrl)

	var res string
	err = chromedp.Run(ctx,
		network.Enable(),
		chromedp.Navigate(fofaUrl),
		chromedp.Evaluate(`window.__NUXT__.data[0].url_key`, &res),
	)
	if err != nil {
		log.Fatalf("err: %s", err)
		return result, err
	}
	//log.Printf("window object keys: %v", res)
	<-done

	// get the downloaded bytes for the request id
	var buf []byte
	if err = chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		buf, err = network.GetResponseBody(requestID).Do(ctx)
		return err
	})); err != nil {
		return result, err
	}

	//fmt.Println(string(buf))
	err = json.Unmarshal(buf, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
