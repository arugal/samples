package hello

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SkyAPM/go2sky"
	httpp "github.com/SkyAPM/go2sky/plugins/http"
)

func whenErr(w http.ResponseWriter, err error) {
	fmt.Fprint(w, err.Error())
	w.WriteHeader(http.StatusInternalServerError)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	request, err := http.NewRequest("POST", fmt.Sprintf("%s/end", "http://192.168.2.101:8081"), nil)
	if err != nil {
		whenErr(w, err)
		return
	}

	client, err := httpp.NewClient(go2sky.GetGlobalTracer())
	if err != nil {
		whenErr(w, err)
		return
	}

	resp, err := client.Do(request)
	if err != nil {
		whenErr(w, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		whenErr(w, err)
		return
	}

	fmt.Fprint(w, body)
}
