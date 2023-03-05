package core

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Playlist interface {
	Play()
	Pause()
	AddSong()
	Next()
	Prev()
}

type Song struct {
	Name string
	Duration int
	Author string
	Id int
}

type SimplePlaylist struct {
	Name string
	Ctx context.Context
	Songs *list.List

	currentSongNode *list.Element
	currentSongPlayTime int
	pauseCtx context.CancelFunc
	coreMtx *sync.Mutex
	isPlaying bool
}

func CreateSimplePlaylist(name string, songs *list.List, ctx context.Context) *SimplePlaylist {
	p := &SimplePlaylist{Name: name, Songs: songs, Ctx: ctx}

	p.coreMtx = new(sync.Mutex)
	p.currentSongPlayTime = 0
	p.isPlaying = false
	p.pauseCtx = nil
	p.currentSongNode = p.Songs.Front()
	return p
}

func (p *SimplePlaylist) Play() {
	if p.isPlaying {
		return
	}

	if p.currentSongNode == nil {
		p.currentSongNode = p.Songs.Front()
	}
	ctx, cancel := context.WithCancel(p.Ctx)
	p.pauseCtx = cancel
	go func() {
		p.isPlaying = true
		for {
			currentSong := p.currentSongNode.Value.(*Song)
			fmt.Println("Now playing:", currentSong.Name)
			for p.currentSongPlayTime != currentSong.Duration {
				for i := 0; i != 5; i++ {
					select {
					case <- ctx.Done():
						return
					default:
						time.Sleep(200 * time.Millisecond)
					}
				}
				p.currentSongPlayTime++
				fmt.Printf("%s: la-la-la(%ds)\n", currentSong.Name, p.currentSongPlayTime)
			}
			p.currentSongPlayTime = 0
			if p.currentSongNode == p.Songs.Back() {
				p.currentSongNode = p.Songs.Front()
			} else {
				p.currentSongNode = p.currentSongNode.Next()
			}
		}
	}()
}

func (p *SimplePlaylist) Pause() {
	fmt.Println("Paused")
	if !p.isPlaying {
		return
	}
	p.isPlaying = false
	p.pauseCtx()

}
func (p *SimplePlaylist) AddSong(song *Song) {
	p.coreMtx.Lock()
	fmt.Printf("New song added. Name: %s. Author: %s. Duration: %ds.", song.Name, song.Author, song.Duration)
	p.Songs.PushBack(song)
	p.coreMtx.Unlock()
}

func (p *SimplePlaylist) Next() {
	p.coreMtx.Lock()
	p.Pause()
	p.currentSongPlayTime = 0
	if p.currentSongNode == p.Songs.Back() {
		p.currentSongNode = p.Songs.Front()
	} else {
		p.currentSongNode = p.currentSongNode.Next()
	}
	p.Play()
	p.coreMtx.Unlock()
}

func (p *SimplePlaylist) Prev() {
	p.coreMtx.Lock()
	p.Pause()
	p.currentSongPlayTime = 0
	if p.currentSongNode == p.Songs.Front() {
		p.currentSongNode = p.Songs.Back()
	} else {
		p.currentSongNode = p.currentSongNode.Prev()
	}
	p.Play()
	p.coreMtx.Unlock()
}

func (p *SimplePlaylist) DeleteSong(name string) error {
	var err error = errors.New("not found")

	for node := p.Songs.Front(); node != nil; node = node.Next() {
		if node.Value.(*Song).Name == name {
			p.Songs.Remove(node)
			err = nil
			break
		}
	}
	return err
}

func (p *SimplePlaylist) GetSongs() list.List {
	return *p.Songs
}

func (p *SimplePlaylist) GetSong(name string) (Song, error) {
	var song *Song = &Song{Author: "", Name: "", Duration: 0}

	for node := p.Songs.Front(); node != nil; node = node.Next() {
		if node.Value.(*Song).Name == name {
			song = node.Value.(*Song)
		}	
	}
	return *song, errors.New("not found")
}