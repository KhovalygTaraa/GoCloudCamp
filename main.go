package main

import (
	"fmt"
	"music_service/service"
	"database/sql"
	"net"
	"github.com/KhovalygTaraa/music_service/api"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	connStr := "user=gocloud password=gocloud dbname=playlist sslmode=disable host=192.168.1.3 port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("DB connection error")
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", "192.168.1.2:9000")
	if err != nil {
		fmt.Printf("listen error")
        panic("listen error")
    }
	api.RegisterMusicServiceServer(grpcServer, service.NewService(db))
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}