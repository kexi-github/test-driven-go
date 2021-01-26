package Applications

import (
	"encoding/json"
	"os"
)

type FileSystemStore struct {
	Database *os.File
	league []Player
}

func NewFileSystemStore(database *os.File) *FileSystemStore {
	info, _ := database.Stat()
	if info.Size()==0 {
		database.Write([]byte("[]"))
	}

	var league []Player
	database.Seek(0,0)
	json.NewDecoder(database).Decode(&league)

	return &FileSystemStore{
		Database: database,
		league: league,
	}
}

func (f *FileSystemStore) GetPlayerScore(name string) int {

	var wins int
	for _,player := range f.league{
		if player.Name == name{
			wins = player.Wins
			break
		}
	}
	return wins
}

func (f *FileSystemStore) RecordWin(name string) {

	flag := false

	for i, player := range f.league {
		if player.Name == name {
			f.league[i].Wins++
			flag = true
		}
	}

	if !flag {
		f.league = append(f.league,Player{name,1})
	}

	f.Database.Seek(0,0)
	//os.File 文件有一个 truncate 函数，可以让我们有效地清空文件
	f.Database.Truncate(0)
	json.NewEncoder(f.Database).Encode(f.league)
}

func (f *FileSystemStore) GetLeague() []Player{
	return f.league
}
