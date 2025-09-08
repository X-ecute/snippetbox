install and init golang-migrate:
    docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate:v4.17.0 create -ext sql -dir /db/migrations init 

run mysql:
    docker run --name go-ecomm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password  -d mysql:8.4 

mysql docker shell:
    docker exec -i go-ecomm mysql -uroot -ppassword <<< "CREATE DATABASE go_ecomm;"#does not work in powershell, use gitbash or bash

docker migrate:
    docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate:v4.17.0 -path=/db/migrations -database "mysql://root:password@tcp(localhost:3306)/go_ecomm" up

compile protobuffers:
    protoc \
  --proto_path=ecomm-grpc/pb \
  --go_out=ecomm-grpc/pb --go_opt=paths=source_relative \
  --go-grpc_out=ecomm-grpc/pb --go-grpc_opt=paths=source_relative \
  ecomm-grpc/pb/api.proto