package stock

import (
	"io/ioutil"
	"net/http"
)

func httpget(url string) (bts []byte, err error) {
	rsp, err := http.Get(url)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	bts, err = ioutil.ReadAll(rsp.Body)
	return
}
