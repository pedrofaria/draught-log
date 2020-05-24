package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pedrofaria/draught-log/internal/app/handler"
	_ "github.com/pedrofaria/draught-log/internal/pkg/statik"
	"github.com/pedrofaria/draught-log/internal/pkg/stream"
	"github.com/pedrofaria/draught-log/internal/pkg/stream/formatter"
	"github.com/pedrofaria/draught-log/internal/pkg/stream/provider"
	statikFs "github.com/rakyll/statik/fs"
)

var envDev bool
var enableDocker bool
var logsDir string

func main() {
	flag.BoolVar(&envDev, "dev", false, "Development mode")
	flag.BoolVar(&enableDocker, "enable-docker", false, "Enable docker provider")
	flag.StringVar(&logsDir, "logs-dir", "", "Path to logs directory")
	flag.Parse()

	if enableDocker {
		log.Println("Docker provider enabled.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client", 301)
	})

	var fs http.FileSystem

	if envDev {
		log.Println("Dev mode enabled")
		fs = http.Dir("./client/public")
	} else {
		statikFS, err := statikFs.New()
		if err != nil {
			log.Fatal(err)
		}
		fs = statikFS
	}

	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(fs)))

	manager := &stream.Manager{
		IsDockerEnabled: enableDocker,
		FileDirectory:   logsDir,
	}

	providerConfig := formatter.Config{
		MessageField:    "message",
		LevelField:      "level",
		TimestampField:  "timestamp",
		TimestampFormat: "2006-01-02T15:04:05Z",
	}

	dockerP := formatter.NewJSON(provider.NewDockerProvider("/development_pablo_api_1"), providerConfig)
	manager.RegisterProvider(dockerP) // TODO move to a handler

	fileP := formatter.NewJSON(provider.NewFileProvider(logsDir, "/pablo.log"), providerConfig)
	manager.RegisterProvider(fileP) // TODO move to a handler

	http.Handle("/api/resources", handler.NewResourcesHandler(manager))
	http.Handle("/api/stream", handler.NewStreamHandler(manager))

	log.Println("Access http://localhost:5000 to see your logs...")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
