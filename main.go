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
	connStr := "user=gocloud password=gocloud dbname=playlist sslmode=disable host=0.0.0.0 port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("DB connection error")
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
        panic(err)
    }
	api.RegisterMusicServiceServer(grpcServer, service.NewService(db))
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
	// playlist.PushBack(&core.Song{Name: "Blue bird", Duration: 40})
	// playlist.PushBack(&core.Song{Name: "All of you", Duration: 120})
	// p := core.CreateSimplePlaylist("My favorite playlist", playlist, context.Background())
	// var act string

	// for act != "stop" {
	// 	// fmt.Scanln(&act)
	// 	scanner  := bufio.NewScanner(os.Stdin)
	// 	scanner.Split(bufio.ScanLines)
	// 	scanner.Scan()
	// 	act = scanner.Text()
	// 	if act == "play" {
	// 		fmt.Println("Playing")
	// 		p.Play()
	// 	} else if act == "pause" {
	// 		fmt.Println("Paused")
	// 		p.Pause()
	// 	} else if act == "prev" {
	// 		fmt.Println("Prev song")
	// 		p.Prev()
	// 	} else if act == "next" {
	// 		fmt.Println("Next song")
	// 		p.Next()
	// 	} else if strings.Contains(act, "add") {
	// 		params := strings.Split(act, " ")
	// 		if (len(params) != 3 || params[0] != "add") {
	// 			fmt.Println("can't add new song, wrong format. Use: add <songName> <duration>")
	// 			continue
	// 		}
	// 		dur, err := strconv.Atoi(params[2])
	// 		if err != nil {
	// 			panic("Atoi error")
	// 		}
	// 		song := &core.Song{Name: params[1], Duration: dur}
	// 		fmt.Printf("New song added: %s(%ds)\n", song.Name, song.Duration)
	// 		p.AddSong(song)
	// 	}
	// }
}
