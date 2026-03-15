module messenger/sso

go 1.25.4

require (
	github.com/lib/pq v1.11.2
	messenger/router v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
)

replace messenger/router => ../common/router
