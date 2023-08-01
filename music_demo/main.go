package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lib *MusicManager
var id = 1

func main() {
	fmt.Println(`Enter following commands to control the player:
                     lib list -- View the existing music lib
                     lib add <name><artist><source><type> -- Add a music to the music lib
                     lib remove <name> -- Remove the specified music from the lib
                     play <name> -- Play the specified music`)
	lib = NewMusicManager()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := reader.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommands(tokens)
		} else {
			fmt.Println("Unsupported command: ", tokens[0])
		}
	}
}

func handlePlayCommands(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	music, _ := lib.FindByName(tokens[1])
	if music == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	Play(music.Source, music.Type)
}

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			music, _ := lib.Get(i)
			fmt.Println(i+1, ":", music.Name, music.Artist, music.Source, music.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.AddMusic(Music{
				Id:     strconv.Itoa(id),
				Name:   tokens[2],
				Artist: tokens[3],
				Source: tokens[4],
				Type:   tokens[5],
			})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}
