package main

import (
	"flag"
	"log"
	"server/common/config"
)

var file = flag.String("f", "", "config file path")

func main() {
	flag.Parse()
	config.Init(*file)
	log.Printf("【 Api 】addr on %s", config.RunData.Addr.ApiAddr)

}
