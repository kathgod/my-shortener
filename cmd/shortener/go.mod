module shortener

require internal/app v1.0.0

require (
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/go-chi/chi/v5 v5.0.8 // indirect
)

replace internal/app => ../../internal/app

go 1.19
