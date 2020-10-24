module apap04.com/server

go 1.15

require (
	apap04.com/server/config v0.0.0-00010101000000-000000000000
	apap04.com/server/handlers v0.0.0-00010101000000-000000000000
	apap04.com/server/utils v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/rs/xid v1.2.1 // indirect
	github.com/sirupsen/logrus v1.7.0
)

replace apap04.com/server/handlers => ./handlers

replace apap04.com/server/utils => ./utils

replace apap04.com/server/config => ./config
