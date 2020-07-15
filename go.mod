module go-microservice

go 1.13

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
require (
	google.golang.org/grpc v1.26.0
	github.com/BelkevichAndry/go-microservice v0.0.0-20200714140122-06cdc53ed91a // indirect
	github.com/BelkevichAndry/go-microservice/shippy-service-consignment v0.0.0-20200714140122-06cdc53ed91a // indirect
	github.com/micro/go-micro v1.18.0 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
)
