module github.com/aglide100/chicken_review_webserver

go 1.14

require (
	cloud.google.com/go v0.70.0 // indirect
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/lib/pq v1.8.0
	github.com/markbates/goth v1.64.2
	google.golang.org/appengine v1.6.6 // indirect
)

replace github.com/aglide100/chicken_review_webserver => ../chicken_review_webserver
