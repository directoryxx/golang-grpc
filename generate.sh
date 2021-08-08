#/bin/bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc sum/protosum/sum.proto --go_out=plugins=grpc:.
protoc prime/protoprime/prime.proto --go_out=plugins=grpc:.
