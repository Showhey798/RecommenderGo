package request

import (
	"encoding/json"
	"net/http"
)

func Bind(r *http.Request, out any) error {
	defer func() {
		_ = r.Body.Close()
	}()
	return json.NewDecoder(r.Body).Decode(out)

}
