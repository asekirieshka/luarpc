package data

import discordrpc "github.com/rikkuness/discord-rpc"

type RpcData struct {
	Drpc *discordrpc.Client
}

var RpcDataStore = RpcData{}
