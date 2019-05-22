module main

require (
	gopkg.in/oauth2.v3 v3.10.0
	pkg/api v1.0.0
	pkg/config v1.0.0
	pkg/log v1.0.0
	pkg/router v1.0.0
	pkg/server v1.0.0
	pkg/static v1.0.0
)

replace (
	pkg/api => ./pkg/api
	pkg/config => ./pkg/config
	pkg/log => ./pkg/log
	pkg/router => ./pkg/router
	pkg/server => ./pkg/server
	pkg/static => ./pkg/static
)

go 1.12
