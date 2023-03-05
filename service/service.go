package service

import (
	"container/list"
	"context"
	"database/sql"
	"fmt"
	"music_service/core"
	"time"
	"github.com/KhovalygTaraa/music_service/api"
)

type MusicServiceServer struct {
	api.UnimplementedMusicServiceServer
	playlist *core.SimplePlaylist
	db *sql.DB
}

func getSongsFromDb(db *sql.DB) *list.List{
	songs := list.New()
	rows, err := db.Query("select * from playlist")
	if err != nil {
        panic(err)
    }
	defer rows.Close()
	for rows.Next() {
		song := new(core.Song)
		err = rows.Scan(&song.Id, &song.Duration, &song.Name, &song.Author)
		if err != nil {
			panic(err)
		}
		songs.PushBack(song)
	}
	return songs
}

func isDBAvailable(db *sql.DB) bool{
	var res bool = true

	_, err := db.Query("select 1")
	if err != nil {
		res = false
    }
	return res
}

func NewService(db *sql.DB) api.MusicServiceServer {
	service := MusicServiceServer{}

	for i := 1; !isDBAvailable(db); i++ {
        fmt.Printf("Db is unavailable(%ds)\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Db is available")
	songs := getSongsFromDb(db)
	playlist := core.CreateSimplePlaylist("My favorite playlist", songs, context.Background())
	service.playlist = playlist
	service.db = db
	return service
}

func (srv MusicServiceServer) Play(ctx context.Context, empty *api.Empty) (*api.Empty, error) {
	srv.playlist.Play()
	return new(api.Empty), nil
}

func (srv MusicServiceServer) Pause(ctx context.Context, empty *api.Empty) (*api.Empty, error) {
	srv.playlist.Pause()
	return new(api.Empty), nil
}

func (srv MusicServiceServer) AddSong(ctx context.Context, song *api.Song) (*api.Empty, error){
	s := new(core.Song)
	s.Author = song.Author
	s.Duration = int(song.Duration)
	s.Name = song.Name
	srv.playlist.AddSong(s)
	_, err := srv.db.Exec("insert into playlist (duration, songname, author) values ($1, $2, $3)", song.Duration, song.Name, song.Author)
	if err != nil {
		panic(err)
	}
	return new(api.Empty), nil
}

func (srv MusicServiceServer) Next(ctx context.Context, empty *api.Empty) (*api.Empty, error) {
	srv.playlist.Next()
	return new(api.Empty), nil
}

func (srv MusicServiceServer) Prev(ctx context.Context, empty *api.Empty) (*api.Empty, error) {
	srv.playlist.Prev()
	return new(api.Empty), nil
}