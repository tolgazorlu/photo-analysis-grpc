package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	pb "github.com/tolgazorlu/photo-analysis/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPhotoManagementClient(conn)

	fmt.Println("Welcome to Photo Management System")
	fmt.Println("----------------------------------")
	fmt.Println("Choose what you want to do:")
	fmt.Println("A. Upload new image")
	fmt.Println("B. Update exist image")

	reader := bufio.NewReader(os.Stdin)
	choose, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(choose)

	switch choose {
	case 'A':
		uploadImage(c)
	case 'B':
		updateImage(c)
	}

}
