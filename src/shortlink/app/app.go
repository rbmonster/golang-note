package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortLinkController struct {
}

type shortLinkResq struct {
	ShortLink string `json:"shortlink"`
}

func main() {
	mux := http.NewServeMux()
	controller := ShortLinkController{}
	// http.HandlerFunc 为一个方法类型， HandlerFunc定义了类型的ServeHTTP方法，因此实现该类型即可实现handler方法
	mux.Handle("/api/shorten", http.HandlerFunc(controller.createShortLink))
	mux.Handle("/api/info", http.HandlerFunc(controller.getShortLinkInfo))
	mux.Handle("/{shortlink:[a-zA-Z0-9]{1,11}}", http.HandlerFunc(controller.redirect))
	http.ListenAndServe("localhost:8000", mux)
}

func (controller ShortLinkController) createShortLink(w http.ResponseWriter, req *http.Request) {
	var shortLinkResq shortLinkResq
	if err := json.NewDecoder(req.Body).Decode(&shortLinkResq); err != nil {
		return
	}
	fmt.Println(w, req, shortLinkResq)
	defer req.Body.Close()
}

func (controller ShortLinkController) getShortLinkInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
	vals := req.URL.Query()
	s := vals.Get("shortlink")
	fmt.Println(s)
}

func (controller ShortLinkController) redirect(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}
