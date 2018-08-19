package router

import "regexp"

const (
	// all http verb methods
	GET     = "GET"
	PUT     = "PUT"
	HEAD    = "HEAD"
	POST    = "POST"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"

	// some help constants
	FavIcon = "/favicon.ico"
)

type patternType int8

const (
	PatternStatic   patternType = iota // /home
	PatternRegexp                      // /:id([0-9]+)
	PatternPathExt                     // /*.*
	PatternHolder                      // /:user
	PatternMatchAll                    // /*
)

// Options for the router
type Options struct {
	// Intercept all request. eg. "/site/error"
	Intercept string
	// Ignore last slash char('/'). If is True, will clear last '/'.
	IgnoreLastSlash bool
	// If True, the router checks if another method is allowed for the current route,
	HandleMethodNotAllowed bool
}

// Router definition
type Router struct {
	name string

	initialized bool

	// static routes
	staticRoutes map[string]interface{}

	// cached dynamic routes
	cachedRoutes map[string]interface{}

	// Regular dynamic routing 规律的动态路由
	regularRoutes map[string]interface{}

	// Irregular dynamic routing 无规律的动态路由
	irregularRoutes map[string]interface{}

	currentGroupPrefix string
	currentGroupOption string
}

// "/users/:id" "/users/:id(\d+)"
var varPattern = regexp.MustCompile(`:[a-zA-Z0-9]+`)
var anyPattern = regexp.MustCompile(`[^/]+`)

var globalPatterns = map[string]string{
	"all": `.*`,
	"any": `[^/]+`,
	"num": `[1-9][0-9]*`,
}

// New
func New() *Router {
	return &Router{name: "default"}
}

// Add a route to router
func (r *Router) Add(method, path string, handler interface{}) (route *Route) {

	return
}

// AddRoute a route to router by Route
func (r *Router) AddRoute(route *Route) *Route {
	return route
}

type RouteGroup struct {
	prefix string
}

func (r *Router) Group(path string, mds ...interface{}) *RouteGroup {
	return &RouteGroup{
		prefix: path,
	}
}
