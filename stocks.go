package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const PolygonPath = "https://api.polygon.io"
const ApiKey = "4stgmx8gbZTZAVS8Dim9edSd6w2XnkJS"

type Stock struct {
    Ticker string `json:"ticker"`
    Name   string `json:"name"`
    Price  float64
}

type Values struct {
    Open float64 `json:"open"`
}

func SearchTicker(ticker string) []Stock {
    resp, err := http.Get(PolygonPath + "/v3/reference/tickers/" + 
                strings.ToUpper(ticker) + "?apiKey=" + ApiKey)

    if err != nil {
        log.Fatal(err)
    }

    body, err := io.ReadAll(resp.Body)

    data := struct {
        Results []Stock `json:"results"`
    }{}

    json.Unmarshal(body, &data)
    return data.Results
}

func GetDailyValues(ticker string) Values {
    resp, err := http.Get(PolygonPath + "/v1/open-close/" + strings.ToUpper(ticker) + "/2024-04-09/?" + ApiKey)
    if err != nil {
        log.Fatal(err)
    }
    body, err := io.ReadAll(resp.Body)

    data := Values{}
    json.Unmarshal(body, &data)
    return data
}
