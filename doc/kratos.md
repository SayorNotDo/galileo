
kratos proto add api/user/v1/user.proto

// -t destination dir
kratos proto server api/user/v1/user.proto -t internal/service
