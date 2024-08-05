package apiserver

import (
	"net/http"
)

// Route ...
type Route struct {
	Path       string
	Methods    []string
	HandleFunc func(http.ResponseWriter, http.Request)
}
