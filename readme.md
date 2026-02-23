# LuaRPC

**Zero-config** Discord Rich Presence engine driven by Lua. Lightweight, fast, and infinitely
customizable.

- **🚀 Go-Powered:** Single binary, no heavy dependencies.
- **💡 Simple by Design:** Built for developers, focused on ease-of-use.
- **🔋 Batteries Included:** Built-in HTTP, JSON and sleep support.
- **⚡ Zero Config:** Just `luarpc run main.lua` and you're live.

---

## Simple as Lua

```lua
rpc.set_client_id("1234567891234567891")

while true do
    -- Fetching price from Coinbase API
    local raw_data = http.get("https://api.coinbase.com/v2/prices/BTC-USD/spot")
    local data = json.parse(raw_data)

    local price = data.data.amount

    rpc.update({
        type = 0,
        name = "Crypto Tracker",
        details = "Bitcoin (BTC)",
        state = "Price: $" .. price,
        largeImage = "https://cryptologos.cc/logos/bitcoin-btc-logo.png",
        start_time = os.time()
    })

    print("Price updated: $" .. price)
    sleep(60000) -- Update once per minute
end
```

[Get Started](https://asekirieshka.github.io/luarpc)
