package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	config2 "github.com/pedrofaria/draught-log/internal/app/config"
	"github.com/pedrofaria/draught-log/internal/app/handler"
	"github.com/pedrofaria/draught-log/internal/pkg/config"
	_ "github.com/pedrofaria/draught-log/internal/pkg/statik"
	"github.com/pedrofaria/draught-log/internal/pkg/stream"
	statikFs "github.com/rakyll/statik/fs"
)

var attrPort int
var attrDevMode bool
var attrConfigPath string
var attrHelp bool

func main() {
	flag.StringVar(&attrConfigPath, "config", "", "Config file path")
	flag.BoolVar(&attrDevMode, "dev", false, "Development mode")
	flag.IntVar(&attrPort, "port", 5000, "Port to bind")
	flag.BoolVar(&attrHelp, "h", false, "Show help")
	flag.Parse()

	if attrHelp {
		flag.Usage()
		return
	}

	var fs http.FileSystem

	if attrDevMode {
		log.Println("Dev mode enabled")
		fs = http.Dir("./client/public")
	} else {
		statikFS, err := statikFs.New()
		if err != nil {
			log.Fatal(err)
		}
		fs = statikFS
	}

	manager := &stream.Manager{}

	if attrConfigPath != "" {
		cfg, err := config.LoadFromFile(attrConfigPath)
		if err != nil {
			panic(err)
		}

		if err := config2.RegisterProvidersFromConfig(manager, cfg); err != nil {
			panic(err)
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client", 301)
	})
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(fs)))
	http.Handle("/api/stream", handler.NewStreamHandler(manager))

	hostAddr := fmt.Sprintf("%d", attrPort)
	log.Println("Access http://localhost:" + hostAddr + " to see your logs...")

	if err := http.ListenAndServe(":"+hostAddr, nil); err != nil {
		panic(err)
	}
}
