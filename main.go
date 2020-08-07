// time service
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		type T struct {
			Unixtime int64  `json:"unixtime"`
			RFC3339  string `json:"rfc3339"`
		}
		var t T
		now := time.Now().UTC()
		t.Unixtime = now.Unix()
		t.RFC3339 = now.Format(time.RFC3339)
		j, err := json.Marshal(t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(j)

	})
	http.ListenAndServe(":8123", nil)
}
