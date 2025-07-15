package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"log"
)

func getMigrationPercent() int {
	val := os.Getenv("MOVIES_MIGRATION_PERCENT")
	percent, err := strconv.Atoi(val)
	if err != nil || percent < 0 || percent > 100 {
		return 0
	}
	return percent
}

func proxyRequest(w http.ResponseWriter, r *http.Request, url string) {
	targetURL := url + r.URL.Path
	if r.URL.RawQuery != "" {
		targetURL += "?" + r.URL.RawQuery
	}

	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	req.Header = r.Header.Clone()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("request failed:", err)
		log.Printf("Request: %s %s\n", req.Method, req.URL)
		//log.Printf("Response status: %s\n", resp.Status)

		http.Error(w, "bad gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, vv := range resp.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

