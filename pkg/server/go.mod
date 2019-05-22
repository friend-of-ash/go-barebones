module server

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	pkg/config v1.0.0
)

replace pkg/config => ../config
