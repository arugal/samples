package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/SkyAPM/go2sky"
	httpp "github.com/SkyAPM/go2sky/plugins/http"
	"github.com/SkyAPM/go2sky/reporter"
)

func main() {
	r, err := reporter.NewGRPCReporter("192.168.101.58:11800")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	tracer, err := go2sky.NewTracer("app3", go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}

	client, err := httpp.NewClient(tracer)

	for {
		// call end service
		request, err := http.NewRequest("GET", fmt.Sprintf("%s/end", "http://192.168.101.58:8081"), nil)
		if err != nil {
			log.Fatalf("unable to create http request: %+v\n", err)
		}
		res, err := client.Do(request)
		if err != nil {
			log.Fatalf("unable to do http request: %+v\n", err)
		}
		body, err1 := ioutil.ReadAll(res.Body)
		if err1 != nil {
			log.Println(err1)
		}
		_ = res.Body.Close()
		fmt.Println(string(body))
		time.Sleep(time.Second)
	}
}
