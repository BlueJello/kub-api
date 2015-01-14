package main

import (
	"errors"
	"net/http"

	api "github.com/MattAitchison/kub-api"
)

type Post struct {
}

func (*Post) New() api.Object {
	return &Post{}
}

func (*Post) NewList() api.Object {

	return nil
}

// List selects resources in the storage which match to the selector.
func (*Post) List(ctx api.Context) (api.Object, error) {

	return nil, errors.New("Not working!")
}

func (*Post) IsAnAPIObject() {}

func main() {
	rest := api.NewAPI()

	rest.AddResource(&Post{})

	http.ListenAndServe(":8090", rest)

}
