math.randomseed(os.time())

local sentences = {
    "Fixing a bug, created 5 mins ago",
    "It works on my machine",
    "Writing code to fix the code",
    "Powered by coffee and Goroutines",
    "The code is fine, it's the logic that's wrong",
}

local function get_random_sentence()
    local randomIndex = math.random(#sentences)
    return sentences[randomIndex]
end

rpc.set_client_id("1234567891234567891")

while true do
    local status = get_random_sentence()

    rpc.update({
        type = 0,
        name = "LuaRPC",
        details = status,
        state = "Developing in Go",
        largeImage = "https://www.lua.org/images/lua-logo.gif",
        start_time = os.time()
    })

    print("Status updated: " .. status)
    sleep(15000)
end