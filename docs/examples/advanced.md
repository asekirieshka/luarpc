# Advanced Examples

Take your presence to the next level by integrating external data and configurations.

## Loading Configuration from JSON

Avoid hardcoding your Client ID by using a separate `config.json` file.

**config.json:**

```json
{
    "client_id": "1234567891234567891"
}
```

**main.lua:**

```lua
local config_path = __dirname .. "/config.json"
local config_file = io.open(config_path, "r")

if not config_file then
    error("Config file not found at " .. config_path)
end

local config = json.parse(config_file:read("*all"))
config_file:close()

rpc.set_client_id(config.client_id)

rpc.update({
    type = 2,
    name = "LuaRPC Project",
    details = "Coding",
    state = "Building a cool project",
    largeImage = "https://go.dev/doc/gopher/fifteen.gif",
    start_time = os.time()
})

while true do
    sleep(99999)
end
```

---

## Real-time Bitcoin Price

This example uses the `http` and `json` modules to fetch and display the current BTC price every
minute.

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
