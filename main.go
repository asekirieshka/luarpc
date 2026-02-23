package main

import (
	"fmt"
	"os"

	"github.com/asekirieshka/luarpc/internal/runtime"
	"github.com/asekirieshka/luarpc/internal/runtime/data"
	"github.com/asekirieshka/luarpc/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No arguments specified")
		utils.ShowHelp()
		return
	}

	defer Cleanup()

	args := os.Args[1:]

	switch args[0] {
	case "run":
		if len(args) < 2 {
			fmt.Println("File not specified")
			return
		}

		fileName := args[1]

		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		err = runtime.RunLua(string(content), fileName)
		if err != nil {
			panic(err)
		}

	case "init":
		if len(args) < 2 {
			fmt.Println("Folder not specified.\nRun luarpc init . to initialize project in the current folder.")
			return
		}

		folderName := args[1]

		if err := utils.InitProject(folderName); err != nil {
			panic(err)
		}

	case "help":
		utils.ShowHelp()

	default:
		fmt.Println("Unknown command")
		utils.ShowHelp()
	}
}

func Cleanup() {
	if data.RpcDataStore.Drpc != nil {
		if err := data.RpcDataStore.Drpc.Socket.Close(); err != nil {
			panic(err)
		}
	}
}
