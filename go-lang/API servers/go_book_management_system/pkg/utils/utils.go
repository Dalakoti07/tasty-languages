package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(request *http.Request, X interface{}) {
	if body, err := ioutil.ReadAll(request.Body); err == nil {
		if err := json.Unmarshal([]byte(body), X); err != nil {
			return
		}
	}
}
