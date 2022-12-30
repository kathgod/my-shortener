module app

require (
	github.com/go-chi/chi/v5 v5.0.8
	internal/app v1.0.0
)

replace internal/app => ./

go 1.19
