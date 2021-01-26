package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"tests/Applications"
)

type CLI struct {
	playStore Applications.PlayStore
	in io.Reader
}

func NewCLI(playStore Applications.PlayStore, in io.Reader) *CLI {
	return &CLI{playStore: playStore, in: in}
}

func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.playStore.RecordWin(extractWinner(userInput))
}

func (c *CLI) readLine() string{
	scanner := bufio.NewScanner(c.in)
	scanner.Scan()
	return scanner.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

const dbFileName = "game.db.json"

func main() {

	fmt.Println("Let's play poker")

	fmt.Println("Type {Name} wins to record a win")
	db, _ := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	store := Applications.NewFileSystemStore(db)

	game := CLI{store, os.Stdin}
	game.PlayPoker()
}
