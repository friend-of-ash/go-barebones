module router

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/render v1.0.1
	pkg/api v1.0.0
	pkg/static v1.0.0
)

replace (
	pkg/api => ../api
	pkg/static => ../static
)

go 1.12
