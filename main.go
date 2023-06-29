package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

const IPINFO_URL = "https://ipinfo.io"

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
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

	renderResult(ipInfo)
}

func renderResult(ipInfo IPInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"IP", "City", "Region", "Country", "Loc", "Org", "Timezone"})
	t.AppendRows([]table.Row{
		{
			ipInfo.IP,
			ipInfo.City,
			ipInfo.Region,
			ipInfo.Country,
			ipInfo.Loc,
			ipInfo.Org,
			ipInfo.Timezone,
		},
	})
	t.Render()
}
