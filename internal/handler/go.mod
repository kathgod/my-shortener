module handler

require (
	github.com/go-chi/chi/v5 v5.0.8
	internal/handler v1.0.0
)

replace internal/handler => ./

go 1.19
