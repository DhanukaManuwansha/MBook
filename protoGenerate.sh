#!/bin/bash

# echo "Enter The Name of Attributes Proto File: "
# read at_proto_file

# echo "Enter The Name of Service Proto File: "
# read svc_proto_file

# Thaveesha - 12/07/2022

# user 
protoc --go_out=internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/user_attributes.proto
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/user_svc.proto

# nurse
protoc --go_out=internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/nurse_attributes.proto
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/nurse_svc.proto

# center admin
protoc --go_out=internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/center_admin_attributes.proto
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/center_admin_svc.proto

# super admin
protoc --go_out=internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/super_admin_attributes.proto
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/super_admin_svc.proto

# doctor
protoc --go_out=internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/doctor_attributes.proto
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/doctor_svc.proto

#center
protoc --go_out=internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/center_attributes.proto
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/drive/grpc --proto_path=internal/adapters/framework/drive/grpc/proto internal/adapters/framework/drive/grpc/proto/center_svc.proto
