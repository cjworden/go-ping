package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	URLS []string
)

func main() {
	log.Debug(URLS)
	for _, url := range URLS {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		log.Debug(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Errorf("Failed pining %s.\n\t%s", url, err)
		} else {
			log.Infof("Ping of %s resulted in status code \"%s\".", url, resp.Status)
			if log.GetLevel() == log.DebugLevel {
				body, err := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					log.Errorf("Error reading response body: %s", err)
				} else {
					log.Debugf("%s", body)
				}
			}
		}

	}
}

func init() {
	var urls string
	flag.StringVar(&urls, "urls", "", "Comma serparated list of urls to ping.")
	logLevel := flag.String("log", "Info", "Log level. Defaults to Info.")
	flag.Parse()
	InitializeLogging(*logLevel)
	URLS = strings.Split(urls, ",")
}

func InitializeLogging(logLevel string) {
	switch logLevel {
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		fallthrough
	default:
		log.SetLevel(log.InfoLevel)
	}
}
