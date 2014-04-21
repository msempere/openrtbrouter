package main

import (
    "flag"
    "fmt"
    "github.com/bsm/openrtb"
    "gopkg.in/yaml.v1"
    "io/ioutil"
    "log"
    "net/http"
)

type Tyaml struct {
    ROUTER struct {
        PATH string
        PORT int
        LOG  struct {
            INFO string
        }
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
}

func main() {

    config_file := flag.String("conf", "", "Router yaml configuration file")
    flag.Parse()

    router_conf := Tyaml{}
    content, err := ioutil.ReadFile(*config_file)

    if err != nil {
        log.Fatalf("error: %v", err)
    }

    err = yaml.Unmarshal([]byte(content), &router_conf)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    http.HandleFunc(router_conf.ROUTER.PATH, handler)
    http.ListenAndServe(fmt.Sprintf(":%d", router_conf.ROUTER.PORT), nil)
}
