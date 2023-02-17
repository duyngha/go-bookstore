package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	err := json.NewDecoder(r.Body).Decode(&x)
	if err != nil {
		return
	}
	// if body, err := ioutil.ReadAll(r.Body); err != nil {
	// 	if err := json.Unmarshal([]byte(body), x); err != nil {
	// 		return
	// 	}
	// }
}
