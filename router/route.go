package router

// Route in the router
type Route struct {
	Name string
	// Pattern definition for the route. eg "/users/{id}"
	Pattern string
	// handler for the route. eg. myFunc, &MyController.SomeAction
	Handler interface{}
	Options map[string]string

	// middleware list
	Middles []interface{}

	// metadata
	meta map[string]string

	// domains defaults
}
