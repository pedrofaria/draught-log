package main

import (
	"log"
	"net/http"

	"github.com/pedrofaria/draught-log/internal/pkg/stream/formatter"

	"github.com/pedrofaria/draught-log/internal/app/handler"
	_ "github.com/pedrofaria/draught-log/internal/pkg/statik"
	"github.com/pedrofaria/draught-log/internal/pkg/stream"
	"github.com/pedrofaria/draught-log/internal/pkg/stream/provider"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client", 301)
	})
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(statikFS)))

	manager := &stream.Manager{
		IsDockerEnabled: true,
		FileDirectory:   "",
	}

	providerConfig := formatter.Config{
		MessageField:    "message",
		LevelField:      "level",
		TimestampField:  "timestamp",
		TimestampFormat: "2006-01-02T15:04:05Z",
	}

	dockerP := formatter.NewJSON(provider.NewDockerProvider("c003c3399e21"), providerConfig)

	manager.RegisterProvider(dockerP) // TODO move to a handler

	http.Handle("/api/resources", handler.NewResourcesHandler(manager))
	http.Handle("/api/stream", handler.NewStreamHandler(manager))

	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
