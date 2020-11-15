module github.com/aglide100/chicken_review_webserver

go 1.14

require (
	cloud.google.com/go v0.72.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/lib/pq v1.8.0
	github.com/markbates/goth v1.66.0
)

replace github.com/aglide100/chicken_review_webserver => ../chicken_review_webserver
