package config

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Addr struct {
	Host string `xml:"host"`
	Port uint   `xml:"port"`
}

type Service struct {
	Addr

	Name string `xml:"name,attr"`

	Min uint `xml:"min,attr"`
	Max uint `xml:"max,attr"`
	ID  uint `xml:"id,attr"`
}

type XMLFile struct {
	PortBase uint
	GateWay  struct {
		Addr
		Max  uint `xml:"max"`
		Auth uint `xml:"auth"`
		Key  string
	}

	Proxy struct {
		Addr
		Max  uint `xml:"max"`
		Auth uint `xml:"auth"`
		Key  string
	}
	Social []Service `xml:"Social>Service"`

	SocialName map[string]*Service
}

var cfg = flag.String("c", "./services.xml", "config file")
var configFile XMLFile

func init() {
	flag.Parse()

	file, err := os.Open(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = xml.Unmarshal(bs, &configFile)
	if err != nil {
		log.Fatal(err)
	}

	configFile.SocialName = make(map[string]*Service)

	for i := 0; i < len(configFile.Social); i++ {
		service := &configFile.Social[i]
		configFile.SocialName[service.Name] = service
	}
}

func GetServiceAddr(name string, hint uint) (pro, addr string) {
	if strings.ToLower(name) == "gateway" {
		return GetGatewayAddr(hint)
	} else if strings.ToLower(name) == "proxy" {
		return GetProxyAddr(hint)
	} else {
		service := configFile.SocialName[name]
		if service != nil {
			return buildAddr(&service.Addr, service.ID)
		}
	}
	return "", ""
}

func GetGatewayAddr(hint uint) (protol string, addr string) {
	return buildAddr(&configFile.GateWay.Addr, hint)
}

func GetProxyAddr(hint uint) (protol string, addr string) {
	return buildAddr(&configFile.Proxy.Addr, hint)
}

func buildAddr(cfg *Addr, id uint) (protol, addr string) {
	host := cfg.Host
	port := cfg.Port

	if host == "" {
		host = "localhost"
	}

	if port == 0 {
		port = configFile.PortBase + id
	}
	protol = "tcp"
	addr = fmt.Sprintf("%s:%v", host, port)
	return

}
