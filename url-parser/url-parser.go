package main

import (
	//	"fmt"
	"log"
	"net"
	"net/url"
)

type protocol struct {
	port int
}

func main() {
	uri := "https://identity.zalando.com/oauth2/authorize?realm=emp#foo"

	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(u.Scheme)
	host, port, _ := net.SplitHostPort(u.Host)
	if u.Scheme == "https" {
		port := "443"
		log.Println(port)
		s := protocol{port: 1443}
		log.Println("Struct Port:", s.port)
	} else {

		log.Println("Boo")
	}
	log.Println(u.Host)
	log.Println("host:", host)
	log.Println("port", port)

	log.Println("Path", u.Path)
	log.Println("Fragment", u.Fragment)
	log.Println(u.RawQuery)
}
