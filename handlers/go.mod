module apap04.com/server/handlers

go 1.15

replace apap04.com/server/models => ../models

replace apap04.com/server/utils => ../utils

require (
	apap04.com/server/utils v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)
