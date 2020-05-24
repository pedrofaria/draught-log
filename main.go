package main

import (
	"flag"
	"log"
	"net/http"

	config2 "github.com/pedrofaria/draught-log/internal/app/config"
	"github.com/pedrofaria/draught-log/internal/pkg/config"

	"github.com/pedrofaria/draught-log/internal/app/handler"
	_ "github.com/pedrofaria/draught-log/internal/pkg/statik"
	"github.com/pedrofaria/draught-log/internal/pkg/stream"
	statikFs "github.com/rakyll/statik/fs"
)

var attrDevMode bool
var attrConfigPath string

func main() {
	flag.BoolVar(&attrDevMode, "dev", false, "Development mode")
	flag.StringVar(&attrConfigPath, "config", "", "Config file path")
	flag.Parse()

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

	log.Println("Access http://localhost:5000 to see your logs...")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
