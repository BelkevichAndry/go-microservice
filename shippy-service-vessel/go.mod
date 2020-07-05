module microservices/shippy-service-vessel

go 1.12

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/BelkevichAndry/go-microservice/shippy-service-vessel v0.0.0-20200705123803-d40d707074cc
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v2 v2.8.0
	go.mongodb.org/mongo-driver v1.3.4
)