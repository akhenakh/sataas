//go:generate protoc -I . --go_out=plugins=grpc:. satsvc.proto
//go:generate protoc -I . satsvc.proto --js_out=import_style=commonjs:../js/src --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../js/src
package satsvc
