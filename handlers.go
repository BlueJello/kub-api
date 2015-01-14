package api

import (
	"fmt"
	"net/http"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	return
}

func Kind(req *http.Request) (kind string, parts []string, err error) {
	parts = splitPath(req.URL.Path)
	if len(parts) < 1 {
		err = fmt.Errorf("Unable to determine kind from an empty URL path")
		return
	}

	// handle input of form /api/{version}/* by adjusting special paths
	if parts[0] == "api" {
		if len(parts) > 2 {
			parts = parts[2:]
		} else {
			err = fmt.Errorf("Unable to determine kind from url, %v", req.URL)
			return
		}
	}

	// URL forms: /{kind}/*
	// URL forms: POST /{kind} is a legacy API convention to create in "default" namespace
	// URL forms: /{kind}/{resourceName} use the "default" namespace if omitted from query param
	// URL forms: /{kind} assume cross-namespace operation if omitted from query param
	kind = parts[0]

	return
}
