package luafunctions

import (
	"encoding/json"

	"github.com/asekirieshka/luarpc/internal/utils"

	lua "github.com/yuin/gopher-lua"
)

func JsonParse(l *lua.LState) *lua.LFunction {
	return l.NewFunction(func(l *lua.LState) int {
		jsonString := l.CheckString(1)

		var result any
		if err := json.Unmarshal([]byte(jsonString), &result); err != nil {
			l.RaiseError("Failed to parse JSON: %v", err)
		}

		l.Push(utils.GoToLua(l, result))
		return 1
	})
}
