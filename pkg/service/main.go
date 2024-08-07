package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
	middleware "wvm/overrides/middleware"
	config_service "wvm/pkg/config"
)

func Init() {

	config_service.LoadConfig()
	config := config_service.ConfigData

	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

		val := middleware.ProxyMiddleware(req, config)

		// define origin server URL
		originServerURL, err := url.Parse(val)
		if err != nil {
			log.Fatal("invalid origin server URL")
		}

		log.Printf("originServerURL: %s", originServerURL)

		// set req Host, URL and Request URI to forward a request to the origin server
		req.Host = originServerURL.Host
		req.URL.Host = originServerURL.Host
		req.URL.Scheme = originServerURL.Scheme
		req.RequestURI = ""

		// save the response from the origin server
		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			return
		}

		// return response to the client
		rw.WriteHeader(http.StatusOK)
		io.Copy(rw, originServerResponse.Body)
	})

	port := fmt.Sprintf(":%d", config.Port)

	fmt.Println("Server is listening on port ", port)
	log.Fatal(http.ListenAndServe(port, reverseProxy))
}
