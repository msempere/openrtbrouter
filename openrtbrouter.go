package main

import (
    //"encoding/json"
    "flag"
    "fmt"
    "github.com/bsm/openrtb"
    "github.com/msempere/httpresponse"
    "github.com/op/go-logging"
    "gopkg.in/yaml.v1"
    "io/ioutil"
    stdlog "log"
    "net/http"
    "os"
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

const PACKAGE = "msempere.openrtbrouter"

var log = logging.MustGetLogger(PACKAGE)

func sendErrorResponse(w *http.ResponseWriter, err, details *string) {
}

func sendDroppedBidResponse(w http.ResponseWriter) {
    response := httpresponse.NewHttpResponse(204, "", "none")
    fmt.Fprintf(w, "%s", response.Get())
}

func handler(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    req, err := openrtb.ParseRequest(r.Body)
    if err != nil {
        log.Error(err.Error())
    } else {
        log.Info("Received bid request %s", *req.Id)
    }
    sendDroppedBidResponse(w)
}

func configure_logger(filename, format string) {
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    logging.SetFormatter(logging.MustStringFormatter(format))

    if err == nil {
        logBackend := logging.NewLogBackend(file, "", stdlog.LstdFlags|stdlog.Lshortfile)
        logging.SetBackend(logBackend)
    }

    logging.SetLevel(logging.INFO, PACKAGE)
}

func main() {
    config_file := flag.String("conf", "", "Router yaml configuration file")
    flag.Parse()

    router_conf := Tyaml{}
    content, err := ioutil.ReadFile(*config_file)

    if err != nil {
        log.Error("error: %v", err)
    }

    err = yaml.Unmarshal([]byte(content), &router_conf)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    configure_logger(router_conf.ROUTER.LOG.INFO+"out.log", "%{level} %{message}")

    log.Info("Started Open RTB Router")

    http.HandleFunc(router_conf.ROUTER.PATH, handler)
    http.ListenAndServe(fmt.Sprintf(":%d", router_conf.ROUTER.PORT), nil)
}
