package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joselitofilho/golang-fake-bigquery/internal/logging"
	"github.com/joselitofilho/golang-fake-bigquery/pkg/routes"
)

func main() {
	var discoveryJsonPath string
	var portNum int

	if err := buildConfig(&discoveryJsonPath, &portNum); err != nil {
		logging.Error.Println(err)
		os.Exit(1)
	}

	discoveryJson, err := buildDiscoveryJson(discoveryJsonPath, portNum)
	if err != nil {
		logging.Error.Println(err)
		os.Exit(1)
	}

	if err := listenAndServe(discoveryJson, portNum); err != nil {
		logging.Error.Println(err)
		os.Exit(1)
	}
}

func buildConfig(discoveryJsonPath *string, port *int) error {
	envPort, _ := strconv.Atoi(os.Getenv("FAKE_BQ_PORT"))

	discoveryJsonPathFlag := flag.String("discovery-json-path", "", "path to discovery.json")
	portFlag := flag.Int("port", envPort, "port number to listen at")
	flag.Parse()

	if *discoveryJsonPathFlag == "" {
		return fmt.Errorf("please specify -discovery-json-path")
	}

	if *portFlag == 0 {
		return fmt.Errorf("please specify -port")
	}

	*discoveryJsonPath = *discoveryJsonPathFlag
	*port = *portFlag

	return nil
}

func buildDiscoveryJson(path string, port int) ([]byte, error) {
	discoveryJson, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	myUrl := fmt.Sprintf("http://localhost:%d", port)
	discoveryJson = bytes.Replace(discoveryJson,
		[]byte("https://www.googleapis.com"), []byte(myUrl), -1)

	return discoveryJson, nil
}

func listenAndServe(discoveryJson []byte, portNum int) error {
	app := routes.NewApp(discoveryJson)
	http.HandleFunc("/", app.Route)

	log.Printf("Listening on :%d...", portNum)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", portNum), nil); err != nil {
		return fmt.Errorf("ListenAndServe: %v", err)
	}

	return nil
}
