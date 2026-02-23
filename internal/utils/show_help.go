package utils

import "fmt"

func ShowHelp() {
	fmt.Println(`
LuaRPC - A powerful Discord Rich Presence engine driven by Lua.

Usage:
  luarpc <command> [arguments]

Commands:
  run <file.lua>    Execute a Lua script to start your Discord presence.
                    Example: luarpc run main.lua

  init <folder>     Initialize a new project with templates and definitions.
                    Use "." to initialize in the current directory.
                    Example: luarpc init my-status

  help              Show this helpful message.

For more information, visit: https://github.com/asekirieshka/luarpc`)
}
