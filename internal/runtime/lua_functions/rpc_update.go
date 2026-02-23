package luafunctions

import (
	"github.com/asekirieshka/luarpc/internal/runtime/data"
	"github.com/asekirieshka/luarpc/internal/utils"

	discordrpc "github.com/rikkuness/discord-rpc"
	lua "github.com/yuin/gopher-lua"
)

func RpcUpdate(l *lua.LState) *lua.LFunction {
	return l.NewFunction(func(l *lua.LState) int {
		if data.RpcDataStore.Drpc == nil {
			l.RaiseError("Unable to update RPC: rpc.set_client_id is not called")
		}

		dataTable := l.CheckTable(1)

		activity := discordrpc.Activity{}

		activityType := dataTable.RawGetString("type")
		utils.SetValidLuaNumber(l, activityType, &activity.Type)

		name := dataTable.RawGetString("name")
		utils.SetValidLuaString(l, name, &activity.Name)

		details := dataTable.RawGetString("details")
		utils.SetValidLuaString(l, details, &activity.Details)

		state := dataTable.RawGetString("state")
		utils.SetValidLuaString(l, state, &activity.State)

		activity.Assets = parseAssets(l, dataTable)
		activity.Timestamps = parseTimestamps(l, dataTable)

		if err := data.RpcDataStore.Drpc.SetActivity(activity); err != nil {
			l.RaiseError("Error setting activity: %v", err)
		}

		return 0
	})
}

func parseAssets(l *lua.LState, dataTable *lua.LTable) *discordrpc.Assets {
	largeImg := dataTable.RawGetString("largeImage")
	smallImg := dataTable.RawGetString("smallImage")

	if largeImg.Type() == lua.LTNil && smallImg.Type() == lua.LTNil {
		return nil
	}

	assets := &discordrpc.Assets{}

	utils.SetValidLuaString(l, largeImg, &assets.LargeImage)
	utils.SetValidLuaString(l, smallImg, &assets.SmallImage)

	return assets
}

func parseTimestamps(l *lua.LState, dataTable *lua.LTable) *discordrpc.Timestamps {
	startTime := dataTable.RawGetString("start_time")
	endTime := dataTable.RawGetString("end_time")

	if startTime.Type() == lua.LTNil && endTime.Type() == lua.LTNil {
		return nil
	}

	timestamps := &discordrpc.Timestamps{}

	if startTime.Type() != lua.LTNil {
		timestamps.Start = &discordrpc.Epoch{}
		utils.SetValidLuaNumberEpoch(l, startTime, timestamps.Start)
	}

	if endTime.Type() != lua.LTNil {
		timestamps.End = &discordrpc.Epoch{}
		utils.SetValidLuaNumberEpoch(l, endTime, timestamps.End)
	}

	return timestamps
}
