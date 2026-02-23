package utils

import lua "github.com/yuin/gopher-lua"

func GoToLua(L *lua.LState, value any) lua.LValue {
	switch v := value.(type) {
	case float64:
		return lua.LNumber(v)

	case string:
		return lua.LString(v)

	case bool:
		return lua.LBool(v)

	case map[string]any:
		table := L.NewTable()
		for k, val := range v {
			table.RawSetString(k, GoToLua(L, val))
		}
		return table

	case []any:
		table := L.NewTable()
		for _, val := range v {
			table.Append(GoToLua(L, val))
		}
		return table

	default:
		return lua.LNil
	}
}
