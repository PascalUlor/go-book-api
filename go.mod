module github.com/PascalUlor/go-book-api

go 1.13

require (
	example.com/me/bookdbconfig v0.0.0
	example.com/me/controllers v0.0.0
	example.com/me/driver v0.0.0
	example.com/me/models v0.0.0
	github.com/gorilla/mux v1.7.4
	github.com/lib/pq v1.3.0 // indirect
	github.com/subosito/gotenv v1.2.0
)

replace example.com/me/bookdbconfig => ./config/book

replace example.com/me/models => ./models

replace example.com/me/controllers => ./controllers

replace example.com/me/driver => ./driver
