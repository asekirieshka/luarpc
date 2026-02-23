package luafunctions

import (
	"context"
	"io"
	"net/http"
	"time"

	lua "github.com/yuin/gopher-lua"
)

func HttpGet(l *lua.LState) *lua.LFunction {
	return l.NewFunction(func(l *lua.LState) int {
		url := l.CheckString(1)

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			l.RaiseError("Failed to create request: %v", err)
		}

		client := http.DefaultClient

		resp, err := client.Do(req)
		if err != nil {
			l.RaiseError("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			l.RaiseError("Failed to read response body: %v", err)
		}

		l.Push(lua.LString(string(bodyBytes)))
		return 1
	})
}
