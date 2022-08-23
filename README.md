# Introduction 
TODO: Give a short introduction of your project. Let this section explain the objectives or the motivation behind this project. 

# Getting Started
TODO: Guide users through getting your code up and running on their own system. In this section you can talk about:
1.	Installation process
    ###required packages to be installed
    sqlc -
2.	Software dependencies
3.	Latest releases
4.	API references






# Database Migration
:- to do the migration you should install migrate libary in go lang. visit https://pkg.go.dev/github.com/golang-migrate/migrate/v4 for more :- references
:- installation instructions: https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md
:- if yo have new migrations just follow below 
    1. go to the project foldern
    2. opne terminal there and execute this command: migrate create -ext sql -dir db/migration -seq init_schema
    3. abouve command make two files up and down. Up file push the changes to the working db and down push the changes to the code base from the working db. you can get the idea from this tutorial : https://www.youtube.com/watch?v=0CYkrGIJkpw&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=3
    4. code the sql code you need to make in the migration files respectively on up and down migrations.
    5. the commands for migration up and down are written in Makefile. You can use that commands in folder bash for running the migrations


# proto setup

if you have import issues is date stamps follow this command
PROTOC_ZIP=protoc-3.24.0-osx-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.24.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP

# docker setup
all steps added in make file also
1. create custom network :  docker network create medbook-network
2. create postgresql svr in the custom network: docker run --name medbookpsqlsvr --network medbook-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Medbook@1234NTEC -d postgres:14.2-alpine

if you already created psql svr container you can change the network by using this command: docker network connect medbook-network medbookpsqlsv

3. make server image : docker build -t medbook:latest .
4. run built image on network medbook-network with the setting up the environment variable to mention sql server address: docker run --name medbooksvr --network medbook-network -p 9000:9000 -e DB_SOURCE_NAME="postgres://root:Medbook@1234NTEC@medbookpsqlsvr:5432/medbookDB?sslmode=disable"  medbook:latest

# Build and Test
TODO: Describe and show how to build your code and run the tests. 

# Contribute
TODO: Explain how other users and developers can contribute to make your code better. 

