
kratos proto add api/user/v1/user.proto

// -t destination dir
kratos proto server api/user/v1/user.proto -t internal/service

// errors generate
protoc --proto_path=. \
--proto_path=./third_party \
--go_out=paths=source_relative:. \
--go-errors_out=paths=source_relative:. \
$(API_PROTO_FILES)

or 

make errors
