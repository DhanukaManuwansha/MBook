network:
	docker network create medbook-network

postgress: 
	docker run --name medbookpsqlsvr --network medbook-network -p 5432\:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Medbook@1234NTEC -d postgres\:14.2-alpine

createdb: 
	docker exec -it medbookpsqlsvr createdb --username=root --owner=root medbookDB

createpatientdb: 
	docker exec -it medbookpsqlsvr createdb --username=root --owner=root medbookPatientDB

dropdb: 
	docker exec -it medbookpsqlsvr dropdb medbookDB

droppatientdb: 
	docker exec -it medbookpsqlsvr dropdb medbookPatientDB

migrateup: 
	migrate -path db/migration -database postgres\://root\:Medbook@1234NTEC@localhost\:5432/medbookDB?sslmode=disable -verbose up

migratepatientup: 
	migrate -path db/migration/patient -database postgres\://root\:Medbook@1234NTEC@localhost\:5432/medbookPatientDB?sslmode=disable -verbose up

migratedown: 
	migrate -path db/migration -database postgresql\://root\:Medbook@1234NTEC@localhost\:5432/medbookDB?sslmode=disable -verbose down

sqlc:
	sqlc generate

svrimgbuild:
	docker build -t medbook:latest .

svrcreate:
	docker run --name medbooksvr --network medbook-network -p 9000:9000 -e DB_SOURCE_NAME="postgres://root:Medbook@1234NTEC@medbookpsqlsvr:5432/medbookDB?sslmode=disable"  medbook:latest

run:
	cd cmd && go run main.go

proto:
	bash protoGenerate.sh
	
.PHONY: network postgres createdb dropdb migrateup migratedown sqlc svrimgbuild svrcreate createpatientdb droppatientdb migratepatientup run proto