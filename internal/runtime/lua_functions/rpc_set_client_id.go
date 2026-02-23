package luafunctions

import (
	"github.com/asekirieshka/luarpc/internal/runtime/data"

	discordrpc "github.com/rikkuness/discord-rpc"
	lua "github.com/yuin/gopher-lua"
)

func RpcSetClientId(l *lua.LState) *lua.LFunction {
	return l.NewFunction(func(l *lua.LState) int {
		clientId := l.CheckString(1)

		client, err := discordrpc.New(clientId)
		if err != nil {
			l.RaiseError("Failed to initialize Discord RPC client\nCheck validity of client_id provided to rpc.set_client_id and ensure that discord is open\n%v", err)
		}

		data.RpcDataStore.Drpc = client

		return 0
	})
}
