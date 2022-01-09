package module2

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func Server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str, err := json.Marshal(r.Header)
		if err == nil {
			panic(errors.New("header jsonize eror"))
		}
		w.Header().Set("Location", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(str))
	})
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}