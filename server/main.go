package main

import (
	"context"
	"flag"
	"fmt"

	// "fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/Shubham-Rasal/goser"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	pb.UnimplementedTodoServer
}

var db = make(map[int]*pb.TodoResponse)

func (s *server) CreateTodo(ctx context.Context, req *pb.TodoRequest) (*pb.TodoResponse, error) {

	todo := &pb.TodoResponse{
		Id:                   string(rune(len(db) + 1)),
		Text:                 req.Text,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     []byte{},
		XXX_sizecache:        0,
	}

	db[len(db)+1] = todo

	return todo, nil

}

func (s *server) ReadAllTodo(ctx context.Context, req *pb.EmptyRequest) (*pb.AllTodoResponse, error) {

	var todos []*pb.TodoResponse

	for _, v := range db {
		todos = append(todos, v)
	}

	return &pb.AllTodoResponse{Todos: todos}, nil

}

func main() {

	//flag package is used to parse command line arguments
	//if not used then default value of port will be 50051
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create a new grpc server
	s := grpc.NewServer()

	//register the server with the grpc server
	pb.RegisterTodoServer(s, &server{})

	log.Printf("server listening at %v", listener.Addr())

	//start the grpc server
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	

}

// func main() {
// 	flag.Parse()
// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGreeterServer(s, &server{})
// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
