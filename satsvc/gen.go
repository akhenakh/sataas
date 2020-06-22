//go:generate protoc -I . --go_out=plugins=grpc:. satsvc.proto
//go:generate protoc -I . satsvc.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcweb:.
package satsvc
