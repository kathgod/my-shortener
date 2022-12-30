module main

require internal/handler v1.0.0

require (
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/go-chi/chi/v5 v5.0.8 // indirect
)

replace internal/handler => ../../internal/handler

go 1.19
