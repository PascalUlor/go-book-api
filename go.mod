module github.com/PascalUlor/go-book-api

go 1.13

require (
	example.com/me/bookdbconfig v0.0.0 // indirect
	example.com/me/controllers v0.0.0
	example.com/me/driver v0.0.0
	example.com/me/models v0.0.0
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/vcs v1.13.1 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/golang/dep v0.5.4 // indirect
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/gorilla/mux v1.7.4
	github.com/jmank88/nuts v0.4.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/nightlyone/lockfile v1.0.0 // indirect
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sdboyer/constext v0.0.0-20170321163424-836a14457353 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/subosito/gotenv v1.2.0
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
	golang.org/x/sys v0.0.0-20200321134203-328b4cd54aae // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace example.com/me/bookdbconfig => ./config/book

replace example.com/me/models => ./models

replace example.com/me/controllers => ./controllers

replace example.com/me/driver => ./driver
