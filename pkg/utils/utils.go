// we need unmarshalled data
// the recieved request will be in json
// so we need to unmarshall the data in order access it

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {

	//read the body of the request
	if body, err := ioutil.ReadAll(r.Body); err == nil {

		//unmarshal the data into byte format in the body var
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
