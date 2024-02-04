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

	for {
		fmt.Println("Welcome to Photo Management System")
		fmt.Println("----------------------------------")
		fmt.Println("Choose what you want to do:")
		fmt.Println("A. Upload A New Image")
		fmt.Println("B. Update Exist Image")
		fmt.Println("C. Get Image Detail")
		fmt.Println("D. Get Image Feed")
		fmt.Println("E. Exit Program")

		fmt.Print("Select: ")
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
		case 'C':
			getImageDetail(c)
		case 'D':
			getImageFeed(c)
		case 'E':
			os.Exit(0)
		}
	}

}
