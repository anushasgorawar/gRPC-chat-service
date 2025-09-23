### Basic Server
https://www.youtube.com/watch?v=BdzYdN_Zd9Q

1. https://grpc.io/docs/languages/go/quickstart/
```
go mod init github.com/anushasgorawar/gRPC-chat-service
go get google.golang.org/grpc
go get github.com/golang/protobuf
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. gRPC is a open source RPC framework that can run anywhere.
Used in distributed systems that allows us to send messages, expose methods to other applications to invoke.
Expose functionality within your app using HTTP/2 as a means of communication

3. grpc vs rest
grpc uses http/2 and rest uses http1.1
grpc uses proto buffer data format and rest uses json data format
amount of data in grpc is lower.
grpc helps in bidirectional streamining.
postman can't help to test grpc. No native support.

4. net.Listen : generic low-level network listener (you manage everything).

http.ListenAndServe : high-level HTTP server that internally uses net.Listen("tcp", addr) but also handles the full HTTP protocol for you.

When using net.listen:
Your Go program is a raw TCP server, not an HTTP server.

Option 1: Keep it a raw TCP server

Don’t use curl — just test with nc (netcat) or telnet:
`nc localhost 8085`

Option 2: Make it an HTTP server

If you want curl to work, you need to send proper HTTP response headers.
```
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 11\r\n" +
		"\r\n" +
		"Hello There"

	_, err := conn.Write([]byte(response))
```

## gPRC

1. import grpc package
'go get google.golang.org/grpc'

2. recieve a message and reply to client

3. A message is a structured data type defined using Protocol Buffers.  

A service is a definition in protobuf that specifies a collection of remote methods (RPCs) the server provides, including their request and response message types.

ChatService has function names that can be called by any language of client.

4. create a package ie mkdir 
`protoc --go_out=plugins=grpc:chat chat.proto`

5. error:
(base) Anushas-MacBook-Air:gRPC anushasg$ protoc --go_out=plugins=grpc:chat chat.proto
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1.


protoc: protocol buffer's complier
tool that reads our .proto files and generates code.
--go_out = where to put the generated code of message
--go-grpc_out = where to put the service code of message

need to write `option go_package = "option go_package = "/chat";`

protoc --go_out=. --go-grpc_out=. chat.proto

This generated structs and interfaces.
*.pb.go files contains all the generated gRPC boilerplate: interfaces, structs, registration functions.

6. context = context of the request
`go get golang.org/x/net/context`

7. RegisterChatServiceServer: This is an auto-generated function created by --go-grpc_out.

Its job is to tell the gRPC server which service implementation [chat service] to use.

It takes two arguments:

> The gRPC server (*grpc.Server)

> Your service implementation (struct that has methods matching the service interface)

8. grpc service is now able to expose SayHello function of the chat service.

SayHello method is registered with GRPC