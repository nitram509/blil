package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v1"
	"log"
	"net/http"
	"strconv"
)

const (
	VERSION           = "0.0.1"
	DEFAULT_PORT      = 8080
	DEFAULT_INTERFACE = "127.0.0.1"
)

var (
	flagPort      = kingpin.Flag("port", "listen on port").Short('p').Default(strconv.Itoa(DEFAULT_PORT)).Int()
	flagInterface = kingpin.Flag("interface", "listen on interface").Short('i').Default(DEFAULT_INTERFACE).String()
)

func main() {

	kingpin.Version(VERSION)
	kingpin.Parse()

	router := NewRouter()

	inetAdress := fmt.Sprintf("%s:%d", *flagInterface, *flagPort)
	log.Fatal(http.ListenAndServe(inetAdress, router))
}
