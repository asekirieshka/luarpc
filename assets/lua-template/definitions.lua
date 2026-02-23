-- NOTE: This file is for IDE autocompletion only. Do not delete or edit it if you need LSP support.

---@meta

---@class RpcData
---@field name string activity name (main string)
---@field type number? activity type (0 - Playing, 2 - Listening, 3 - Watching, 5 - Competing)
---@field state string? state (bottom string)
---@field details string? details (top string)
---@field largeImage string? large image key from application assets
---@field smallImage string? small image key
---@field start_time number? Unix timestamp of the start (shows "time elapsed")
---@field end_time number? Unix timestamp of the end (shows "time remaining" via progress bar if activity type is 2 (Listening))

---@class rpc
rpc = {}

--- Updates Discord Rich Presence
---@param data RpcData table containing activity data
function rpc.update(data) end

--- Sets the Discord application client ID
---@param client_id string Discord client ID
function rpc.set_client_id(client_id) end

---@class http
http = {}

--- Performs a GET request and returns the response body as a string
---@param url string URL address
---@return string body Response body
function http.get(url) end

---@class json
json = {}

--- Parses a JSON string and returns a Lua table
---@param json_string string JSON formatted string
---@return table|any result Deserialized data
function json.parse(json_string) end

--- Pause for a specified number of milliseconds (does not block the main Go thread)
---@param ms number milliseconds
function sleep(ms) end

--- Absolute path to the directory containing the current executable .lua file.
---@type string
__dirname = ""