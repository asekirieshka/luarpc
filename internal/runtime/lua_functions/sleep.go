package luafunctions

import (
	"time"

	lua "github.com/yuin/gopher-lua"
)

func Sleep(l *lua.LState) *lua.LFunction {
	return l.NewFunction(func(l *lua.LState) int {
		ms := l.CheckInt(1) // get first argument

		time.Sleep(time.Duration(ms) * time.Millisecond)

		return 0
	})
}
