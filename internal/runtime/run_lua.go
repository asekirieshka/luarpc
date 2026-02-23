package runtime

import (
	luafunctions "github.com/asekirieshka/luarpc/internal/runtime/lua_functions"

	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

func RunLua(code string, scriptPath string) error {
	l := lua.NewState()
	defer l.Close()

	absPath, err := filepath.Abs(scriptPath)
	if err != nil {
		l.RaiseError("file not found")
		return err
	}

	scriptDir := filepath.Dir(absPath)
	l.SetGlobal("__dirname", lua.LString(scriptDir))

	l.SetGlobal("sleep", luafunctions.Sleep(l))

	httpTable := l.NewTable()
	l.SetGlobal("http", httpTable)
	httpTable.RawSetString("get", luafunctions.HttpGet(l))

	jsonTable := l.NewTable()
	l.SetGlobal("json", jsonTable)
	jsonTable.RawSetString("parse", luafunctions.JsonParse(l))

	rpcTable := l.NewTable()
	l.SetGlobal("rpc", rpcTable)
	rpcTable.RawSetString("update", luafunctions.RpcUpdate(l))
	rpcTable.RawSetString("set_client_id", luafunctions.RpcSetClientId(l))

	if err := l.DoString(code); err != nil {
		return err
	}

	return nil
}
