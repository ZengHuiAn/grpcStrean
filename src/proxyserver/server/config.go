package server

import (
	"fmt"
	"gprcStream/src/config"
)

var g_listen_addr = ":10000"

func ParseConfig() error {
	_, g_listen_addr = config.GetServiceAddr("proxy", 1)
	fmt.Println("ParseConfig", g_listen_addr)
	return nil
}

func GetServeAddr() string {
	return g_listen_addr
}
