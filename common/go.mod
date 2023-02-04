module common

go 1.18

require (
	github.com/mitchellh/mapstructure v1.3.3
	go-micro.dev/v4 v4.9.0
)

require (
	github.com/google/uuid v1.2.0 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.3.0 // indirect

)

replace (
	common => ../common
	rpc => ../rpc
)
