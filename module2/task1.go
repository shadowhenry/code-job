package module2

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

func Server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str, err := json.Marshal(r.Header)
		if err != nil {
			panic(errors.New("header jsonize eror"))
		}
		w.Header().Set("Location", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(str))
	})


	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host)
		w.Write([]byte("Os: " + runtime.GOOS + runtime.GOARCH + "\n"))
		w.Write([]byte("Go Version: " + runtime.Version() + "\n"))
		w.Write([]byte("Requet Path: " + r.URL.Path + "\n"))
		w.Write([]byte("Http Code: " + strconv.Itoa(http.StatusOK)))
	})

	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}