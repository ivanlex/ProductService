module ProductService

go 1.17

require (
	github.com/go-playground/validator/v10 v10.10.1
	github.com/gorilla/mux v1.8.0
)

require golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect

require (
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-openapi/errors v0.20.2
	github.com/go-openapi/runtime v0.24.0
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-openapi/swag v0.21.1
	github.com/go-openapi/validate v0.21.0
	github.com/gorilla/handlers v1.5.1
	github.com/kevin/currency v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	google.golang.org/grpc v1.46.0
)

replace github.com/kevin/currency => ./../Currency
