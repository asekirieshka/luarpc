package utils

import (
	"time"

	discordrpc "github.com/rikkuness/discord-rpc"
	lua "github.com/yuin/gopher-lua"
)

func SetValidLuaString(l *lua.LState, luaVal lua.LValue, target *string) {
	if luaVal.Type() == lua.LTNil {
		return
	}

	val, ok := luaVal.(lua.LString)

	if !ok {
		l.RaiseError("expected string, got %s", luaVal.Type().String())
	}

	*target = string(val)
}

func SetValidLuaNumber(l *lua.LState, luaVal lua.LValue, target *int) {
	if luaVal.Type() == lua.LTNil {
		return
	}

	val, ok := luaVal.(lua.LNumber)

	if !ok {
		l.RaiseError("expected number, got %s", luaVal.Type().String())
	}

	*target = int(val)
}

func SetValidLuaNumberEpoch(l *lua.LState, luaVal lua.LValue, target *discordrpc.Epoch) {
	if luaVal.Type() == lua.LTNil {
		return
	}

	val, ok := luaVal.(lua.LNumber)

	if !ok {
		l.RaiseError("expected number, got %s", luaVal.Type().String())
	}

	t := time.Unix(int64(val), 0)
	target.Time = t
}
