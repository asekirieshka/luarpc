# Global API

LuaRPC provides several built-in global functions and variables that are not part of standard Lua
5.1 but are essential for building rich presence scripts.

---

## Global Variables

### \_\_dirname

A string containing the absolute path to the directory where the currently executing Lua script is
located.

```lua
print("Scripts are located in: " .. __dirname)
-- Useful for loading local assets or files
```

---

## Control Functions

### sleep

Pauses the execution of the script for a specified number of milliseconds.

<!-- prettier-ignore -->
> [!IMPORTANT]
> Use this inside your loops to avoid hitting Discord's rate limits and consuming 100% of your CPU.

- **Arguments:** `ms` (number) — Time to sleep in milliseconds.
- **Example:**

```lua
while true do
    print("Updating...")
    -- Your RPC logic here
    sleep(15000) -- Wait for 15 seconds
end
```

---

## HTTP Module

The `http` module allows you to fetch data from external APIs (like GitHub, Steam, or Crypto
prices).

### http.get

Performs a synchronous GET request.

- **Arguments:** `url` (string)
- **Returns:** `string` (response body)
- **Example:**

```lua
local response = http.get("https://api.coinbase.com/v2/prices/BTC-USD/spot")
print("Raw BTC Price: " .. response)
```

---

## JSON Module

Since `http.get` returns a raw string, you'll need the `json` module to parse it into a Lua table.

### json.parse

Converts a JSON string into a Lua table.

- **Arguments:** `data` (string)
- **Returns:** `table`
- **Example:**

```lua
local raw_data = http.get("https://api.github.com/repos/asekirieshka/luarpc")
local repo = json.parse(raw_data)

print("Project Stars: " .. repo.stargazers_count)
```
