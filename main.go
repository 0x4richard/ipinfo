package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const IPINFO_URL = "https://ipinfo.io"

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func main() {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, IPINFO_URL, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	var ipInfo IPInfo
	if err := json.Unmarshal(b, &ipInfo); err != nil {
		log.Fatalf("Failed to parse response body: %v", err)
	}

	fmt.Println(ipInfo)
}
