package main

import (
	//"fmt"
	//"os"
	. "github.com/dtchanpura/cd-go/manage"
	"flag"
	"fmt"
	"log"
)

func main() {
	addPtr := flag.Bool("add", false, "Use this to add a configuration")
	servePtr := flag.Bool("serve", false, "Use this to serve and start listener")
	repoPtr := flag.String("repopath", "", "Path to Repository")
	namePtr := flag.String("name", "", "Name of configuration")

	flag.Parse()
	if *addPtr {
		if *namePtr == "" || *repoPtr == "" {
			log.Fatalf("Repository Name and Path required.")
		}
		AddConfiguration(*namePtr, *repoPtr, "")
	} else if *servePtr {
		fmt.Println("Serving...")

	} else {
		fmt.Println("Atleast add or serve flag is required.")
	}

}
