package order

//go:generate protoc --experimental_allow_proto3_optional -I . -I .. --go_out=plugins=grpc:. --go_opt paths=source_relative pb/apisix.proto
//go:generate gomodifytags -skip-unexported -w -file pb/apisix.pb.go -all -remove-tags json
//go:generate gomodifytags -skip-unexported -w -file pb/apisix.pb.go -all -add-tags json
//go:generate gomodifytags -skip-unexported -w -file pb/apisix.pb.go -struct ListReq -add-tags form
