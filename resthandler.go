package api

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

func NewAPI() *API {
	return &API{
		resources: make(map[string]Resource),
	}
}

// API implements HTTP verbs on a set of RESTful resources identified by name.
type API struct {
	resources map[string]Resource
	// codec            runtime.Codec
	// canonicalPrefix string
	// selfLinker       runtime.SelfLinker
	// ops              *Operations
	// asyncOpWait      time.Duration
	// admissionControl admission.Interface
}

func (h *API) AddResource(obj Resource) {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		panic("All types must be pointers to structs.")
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		panic("All types must be pointers to structs.")
	}

	name := strings.ToLower(t.Name())

	h.resources[name] = obj

	log.Printf("Added Resource: %s, %s", name, t)
}

// ServeHTTP handles requests to all Resource objects.
func (h *API) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	kind, parts, err := Kind(req)
	if err != nil {
		notFound(w, req)
		return
	}
	resource := h.resources[kind]
	if resource == nil {
		notFound(w, req)
		return
	}

	h.handleAPIResource(parts, req, w, resource)
}

// handleAPIResource is the main dispatcher for a storage object.  It switches on the HTTP method, and then
// on path length, according to the following table:
//   Method     Path          Action
//   GET        /foo          list
//   GET        /foo/bar      get 'bar'
//   POST       /foo          create
//   PUT        /foo/bar      update 'bar'
//   DELETE     /foo/bar      delete 'bar'
// Returns 404 if the method/pattern doesn't match one of these entries
// The s accepts several query parameters:
//    sync=[false|true] Synchronous request (only applies to create, update, delete operations)
//    timeout=<duration> Timeout for synchronous requests, only applies if sync=true
//    labels=<label-selector> Used for filtering list operations
func (h *API) handleAPIResource(parts []string, req *http.Request, w http.ResponseWriter, storage Resource) {
	// ctx := api.WithNamespace(api.NewContext(), namespace)

	switch req.Method {
	case "GET":
		switch len(parts) {
		case 1:
			lister, ok := storage.(ResourceLister)
			if !ok {
				http.Error(w, "Not Supported", http.StatusMethodNotAllowed)
				return
			}
			list, err := lister.List(Context{})
			if err != nil {

				http.Error(w, "Idk!", http.StatusInternalServerError)
				return
			}
			log.Println(list)
		default:
			notFound(w, req)
		}

	case "POST":
	case "DELETE":

	case "PUT":

	default:
		log.Println("Not found")
		notFound(w, req)
	}
}
