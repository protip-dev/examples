# Protip Examples

This repository contains example server / client applications in a variety of languages.

Submit a Pull Request or file an Issue to see more!

# Go
```sh
$ cd golang

# Launch HelloWorld gRPC server running at localhost:8888 by default.
# Use -port to control the port.
#
# Connect to it using sample client below or a tool like grpcurl:
# https://github.com/fullstorydev/grpcurl
$ go run github.com/protip-dev/examples/cmd/grpcserver

# Run HelloWorld grpc client targeting localhost:8888 by default.
# Setting -language is required.
# Use -target to control the target.
$ go run github.com/protip-dev/examples/cmd/grpcclient

# Launch HelloWorld gRPC-Web server running at localhost:8888 by default.
# Use -port to control the port.
#
# Connect to it using the React Typescript demo web app below.
$ go run github.com/protip-dev/examples/cmd/grpcwebserver
```

# Typescript
```sh
$ cd typescript/client/
$ yarn

# Launch HelloWorld gRPC-Web client targeting http://localhost:8888 by default.
# Set REACT_APP_HELLOWORLD_API the environment variable to control the target,
# including the protocol ("http" or "https").
$ yarn start
```
