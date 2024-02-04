PROJECT_DIR = github.com/tolgazorlu/photo-analysis
BIN_DIR = bin
SERVER_DIR = server
CLIENT_DIR = client

# make protoc
protoc:
	protoc -Iproto --go_out=. --go_opt=module=${PROJECT_DIR} --go-grpc_out=. --go-grpc_opt=module=${PROJECT_DIR} proto/*.proto

client:
	go build -o ${BIN_DIR}/${CLIENT_BIN} ./${CLIENT_DIR}

server: 
	go build -o ${BIN_DIR}/${SERVER_BIN} ./${SERVER_DIR}

# make all build
build: 
	protoc -Iproto --go_out=. --go_opt=module=${PROJECT_DIR} --go-grpc_out=. --go-grpc_opt=module=${PROJECT_DIR} proto/*.proto
	go build -o ${BIN_DIR}/${SERVER_BIN} ./${SERVER_DIR}
	go build -o ${BIN_DIR}/${CLIENT_BIN} ./${CLIENT_DIR}

# make clean
clean: ## Clean generated files for blog
	rm -f proto/*.pb.go
