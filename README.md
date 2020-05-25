# Draught Log

Simple log viewer for development environment.

## Installation

`$ go get -u github.com/pedrofaria/draught-log`

## Usage

```
$ draught-log -h                                                                                                                                                       ✔  6533  17:59:47
Usage of draught-log:
  -config string
    	Config file path
  -dev
    	Development mode
  -h	Show help
  -port int
    	Port to bind (default 5000)
```

## Configuration

YAML file:

````yaml
resources:
  # name to be displayed in provider column in UI.
  - name: example1
    # provider defines from where the server will fetch the log.
    # possible values: docker | file
    provider: docker 
    # docker: container name or container ID
    # file: file path
    providerId: my_container_name
    # Considering the log line:
    # [14:42:59][API] : {"elapsed":67.047686,"env":"development","level":"info","message":"Request handled","timestamp":"2020-05-25T14:42:59.722714315Z"}
    formatter:
      type: json
      messageField: message
      levelField:  level
      timestampField: timestamp
      # Use the date: "Mon Jan 2 15:04:05 -0700 MST 2006" to define your format.
      # https://golang.org/pkg/time/#Time.Format
      timestampFormat: "2006-01-02T15:04:05.999999999Z"
      # Golang regular expression
      # https://github.com/google/re2/wiki/Syntax
      preFilterRegex: "^.*? : "
      preFilterRegexReplace: ""

  - name: example2
    # ...
```` 

## Contribute

* Clone this repository.
* Run `$ cd client && npm install && cd ..` to get all svelte dependencies.
* Run `$ make dev` to make NPM watch for changes in the ./client directory and building it automatically.
* In a new terminal run draught-log in dev mode, so the server will serve the UI files from the ./client/public directory.

`$ go run main.go -dev -config=/path/to/config/file.yml [-port=5000]`


## Build

`$ make build`
