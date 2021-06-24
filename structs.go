package main

import "time"

type tomlConfig struct {
	Title      string
	HttpServer httpserver
	Modbus    mbus
	Controller controller
}

type httpserver struct {
	Listen string
	Port string
}

type mbus struct {
	Timeout duration
}

type duration struct {
	time.Duration
}

type controller struct {
	IP   string
	Port string
}