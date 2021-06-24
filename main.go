package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pterm/pterm"
	"net/http"
	"os"
)

func init() {
	//prometheus.MustRegister(newSolarCollector())
}

var (
	conf tomlConfig
)

func main() {
	pterm.EnableDebugMessages()
	//if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
	//	pterm.Error.Println(fmt.Sprintf("Config error: %s", err))
	//	os.Exit(1)
	//}
	confTree, err := toml.LoadFile("config.toml")
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Config Error: %s", err))
		os.Exit(1)
	}
	err = confTree.Unmarshal(&conf)
	if err != nil {
		pterm.Error.Println(fmt.Sprintf("Config Error: %s", err))
		os.Exit(1)
	}

	pterm.Info.Println("Starting epsolar_exporter")
	pterm.Info.Println(conf)
	pterm.Info.Println(fmt.Sprintf("Listening on %s:%s", conf.HttpServer.Listen, conf.HttpServer.Port))
	r := prometheus.NewRegistry()
	r.MustRegister(newSolarCollector())
	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)

	pterm.Fatal.Println(http.ListenAndServe(fmt.Sprintf("%s:%s", conf.HttpServer.Listen, conf.HttpServer.Port), nil))
}
