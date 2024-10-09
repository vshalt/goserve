package main

import (
	"flag"
	"log"
	"os"

	"github.com/vshalt/goserve/internal/config"
	"github.com/vshalt/goserve/internal/server"
)

func main() {
	var opt config.Flags
	flag.BoolVar(&opt.Verbose, "verbose", false, "Show verbose output")
	flag.BoolVar(&opt.SSL, "ssl", false, "Enable HTTPS server")
	flag.StringVar(&opt.Directory, "dir", "", "Directory path to start the server from")
	flag.StringVar(&opt.Host, "host", "127.0.0.1", "Host address for the server")
	flag.StringVar(&opt.Port, "port", "3333", "Port to listen from")
	flag.StringVar(&opt.CertFile, "cert", "", "Path for the SSL cert file")
	flag.StringVar(&opt.KeyFile, "key", "", "Path for the SSL key file")

	flag.Parse()
	log := log.New(os.Stderr, "goserve: ", log.LstdFlags)

	if opt.SSL && (opt.CertFile == "" || opt.KeyFile == "") {
		log.Println("Cert file and key file are required if SSL flag is set")
		os.Exit(1)
	}

	if opt.Directory == "" {
		dir, err := os.Getwd()
		if err != nil {
			log.Printf("Cannot get current working directory: %v\n", err)
			os.Exit(1)
		}
		opt.Directory = dir
	}
	dir, err := config.Clean(opt.Directory)
	if err != nil {
		log.Printf("Directory error: %v\n", err)
		os.Exit(1)
	}
	opt.Directory = dir

	err = server.Server(log, opt)
	if err != nil {
		log.Printf("Server failed with error: %v\n", err)
		os.Exit(1)
	}
}
