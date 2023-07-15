package main

import (
	"context"
	"database/sql"
	"final_project_backend/config"
	"final_project_backend/internal/pkg/database"
	pb "final_project_backend/pbGenerated"
	"final_project_backend/pkg/slice"
	"flag"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type server struct {
	pb.UnimplementedUsersServiceServer
	db database.Querier
}

func (s *server) SignUpUser(ctx context.Context, req *pb.SignUpUserRequest) (*pb.SignUpUserResponse, error) {
	username := req.Username
	password := req.Password
	_, err := s.db.CreatePerson(ctx, database.CreatePersonParams{Username: username, Password: password})
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok &&
			pqerr.Code == "23505" {
			return &pb.SignUpUserResponse{Success: false, Message: "this username is already taken"}, nil
		} else {
			return nil, fmt.Errorf("create person %w", err)
		}
	}

	return &pb.SignUpUserResponse{Success: true, Message: "person created successfully"}, nil
}

func (s *server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	username := req.Username
	password := req.Password

	person, err := s.db.GetPersonByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get person %w", err)
	}

	if person.Password != password {
		return &pb.LoginUserResponse{Success: false, Message: "username or password is invalid"}, nil
	}

	return &pb.LoginUserResponse{Success: true, Message: "logged in successfully"}, nil
}

func (s *server) AddCredit(ctx context.Context, req *pb.AddCreditRequest) (*pb.AddCreditResponse, error) {
	username := req.Username

	person, err := s.db.GetPersonByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get person %w", err)
	}

	totalCredit := req.Credit + person.Credit.Int32

	err = s.db.UpdateCredit(ctx, database.UpdateCreditParams{Credit: sql.NullInt32{Valid: true, Int32: totalCredit}})
	if err != nil {
		return nil, fmt.Errorf("add credit %w", err)
	}

	return &pb.AddCreditResponse{Credit: totalCredit, Message: "add credit successfully"}, nil
}

func (s *server) UnavailableDates(ctx context.Context, req *pb.UnavailableDatesRequest) (*pb.UnavailableDatesResponse, error) {
	roomId := req.RoomId

	reservations, err := s.db.GetRoomsReservedDates(ctx, roomId)
	if err != nil {
		return nil, fmt.Errorf("get person %w", err)
	}

	var reservedDates []int32
	for _, reservation := range reservations {
		reservedDates = append(reservedDates, reservation...)
	}

	reservedDates = slice.Unique(reservedDates)

	return &pb.UnavailableDatesResponse{Dates: reservedDates}, nil
}

func main() {
	flag.Parse()
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("load config %v", err)
	}
	conn, err := database.NewDBConnection(&conf.DB)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db := database.New(conn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUsersServiceServer(s, &server{db: db})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	wg.Wait()
}
