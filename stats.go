package main

import (
	"crypto/tls"
	"encoding/json"
	"github.com/murlokswarm/app"
	"net/http"
	"net/url"
	"time"
)

type Stats struct {
	Data []Pool
}

type Pool struct {
	Name        string  `json:name`
	Url         string  `json:url`
	StatsUrl    string  `json:statsUrl`
	Miners      int64   `json:miners`
	Hashrate    float64 `json:hashrate`
	Uptime      float64 `json:uptime`
	Software    string  `json:software`
	Count       int64   `json:count`
	OnlineCount int64   `json:onlineCount`
}

func (s *Stats) OnNavigate(u *url.URL) {
	app.Debug(u.String())

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	var timeout = time.Duration(30 * time.Second)
	var client = http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	var pools []Pool

	res, err := client.Get("https://www.ubiq.cc/api/v1/stats/pools")
	if err != nil {
		app.Debug("Error getting json: %#v", err)
	}

	err = json.NewDecoder(res.Body).Decode(&pools)
	if err != nil {
		app.Debug("Error reading json: %v", err)
	}

	app.Debug("%v", pools)

	s.Data = pools
	app.Render(s)

}

// func (s *Stats) OnMount() {
//
// 	transport := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
//
// 	var timeout = time.Duration(30 * time.Second)
// 	var client = http.Client{
// 		Timeout:   timeout,
// 		Transport: transport,
// 	}
//
// 	var pools []Pool
//
// 	res, err := client.Get("https://www.ubiq.cc/api/v1/stats/pools")
// 	if err != nil {
// 		app.Debug("Error getting json: %#v", err)
// 	}
//
// 	err = json.NewDecoder(res.Body).Decode(&pools)
// 	if err != nil {
// 		app.Debug("Error reading json: %v", err)
// 	}
//
// 	app.Debug("%v", pools)
//
// 	s.Data = pools
// 	app.Render(s)
//
// }

func (s *Stats) Render() string {

	return `{{.Data}}`
}
