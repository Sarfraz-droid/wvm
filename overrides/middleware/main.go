package middleware

import (
	"log"
	"net/http"
	"wvm/pkg/config"
)

func ProxyMiddleware(req *http.Request, config *config.Config) string {
	// Modify middleware here - this is where you can add your own middleware
	version, err := req.Cookie("version")

	if version == nil || err != nil {
		version = &http.Cookie{Name: "version", Value: "v1"}
	}

	val, ok := config.Version[version.Value]

	if !ok {
		val = config.Version["default"]
	}

	log.Printf("version: %s", version.Value)

	return val
}
