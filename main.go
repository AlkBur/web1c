package main

import (
	"github.com/AlkBur/web1c/server"
	"log"
)

func main() {
	conf, err := parseConfig(fileCFG)
	checkError(err)
	debug = conf.Debug

	//Create server
	s := server.New()
	s.Debug = conf.Debug

	//Add bases
	for _, val := range conf.Bases {
		checkError(s.AddDB(val.Name, val.Path, val.Url))
	}

	//Run server
	checkError(s.Run(conf.Addr))
}

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
