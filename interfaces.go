package api

type Object interface {
	IsAnAPIObject()
}

// RESTStorage is a generic interface for RESTful storage services.
// Resources which are exported to the RESTful API of apiserver need to implement this interface. It is expected
// that objects may implement any of the REST* interfaces.
// TODO: implement dynamic introspection (so GenericREST objects can indicate what they implement)
type Resource interface {
	// New returns an empty object that can be used with Create and Update after request data has been put into it.
	// This object must be a pointer type for use with Codec.DecodeInto([]byte, Object)
	New() Object
}

type ResourceLister interface {
	// NewList returns an empty object that can be used with the List call.
	// This object must be a pointer type for use with Codec.DecodeInto([]byte, Object)
	NewList() Object
	// List selects resources in the storage which match to the selector.
	List(ctx Context) (Object, error)
}

type ResourceGetter interface {
	// Get finds a resource in the storage by id and returns it.
	// Although it can return an arbitrary error value, IsNotFound(err) is true for the
	// returned error value err when the specified resource is not found.
	Get(ctx Context, id string) (Object, error)
}

type ResourceDeleter interface {
	// Delete finds a resource in the storage and deletes it.
	// Although it can return an arbitrary error value, IsNotFound(err) is true for the
	// returned error value err when the specified resource is not found.
	Delete(ctx Context, id string) (<-chan ResourceResult, error)
}

type ResourceCreater interface {
	// Create creates a new version of a resource.
	Create(ctx Context, obj Object) (<-chan ResourceResult, error)
}

type ResourceUpdater interface {
	// Update finds a resource in the storage and updates it. Some implementations
	// may allow updates creates the object - they should set the Created flag of
	// the returned RESTResultto true. In the event of an asynchronous error returned
	// via an Status object, the Created flag is ignored.
	Update(ctx Context, obj Object) (<-chan ResourceResult, error)
}

// ResourceResult indicates the result of a REST transformation.
type ResourceResult struct {
	// The result of this operation. May be nil if the operation has no meaningful
	// result (like Delete)
	Object

	// May be set true to indicate that the Update operation resulted in the object
	// being created.
	Created bool
}

// Redirector know how to return a remote resource's location.
type Redirector interface {
	// ResourceLocation should return the remote location of the given resource, or an error.
	ResourceLocation(ctx Context, id string) (remoteLocation string, err error)
}
