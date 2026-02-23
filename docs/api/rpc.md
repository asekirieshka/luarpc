# RPC Module

The `rpc` module provides the interface to communicate with the Discord Desktop client.

## rpc.set_client_id

Initializes the connection to Discord. This must be the first RPC function you call.

- **Arguments:** `id` (string) — Your Discord Application Client ID.
- **Example:**

```lua
rpc.set_client_id("1234567891234567891")
```

---

## rpc.update

Updates your Discord Rich Presence status. It takes a single table as an argument.

```lua
rpc.update({
    type = 0,
    name = "My Awesome App",
    details = "Developing scripts",
    state = "In Editor",
    largeImage = "https://example.com/large.png",
    largeText = "LuaRPC Engine",
    smallImage = "https://example.com/small.png",
    start_time = os.time()
})
```

### Activity Fields Reference

| Field        | Type     | Description                                                                 |
| :----------- | :------- | :-------------------------------------------------------------------------- |
| `type`       | `number` | Activity type. `0`: Playing, `2`: Listening, `3`: Watching, `5`: Competing. |
| `name`       | `string` | The name of the application.                                                |
| `details`    | `string` | The first line of text (e.g., "Level 100 Mage").                            |
| `state`      | `string` | The second line of text (e.g., "Exploring Dungeon").                        |
| `largeImage` | `string` | URL or asset key for the large square image.                                |
| `largeText`  | `string` | Hover text for the large image.                                             |
| `smallImage` | `string` | URL or asset key for the small circular image.                              |
| `smallText`  | `string` | Hover text for the small image.                                             |
| `start_time` | `number` | Unix timestamp. Shows "elapsed" time (e.g., `01:45 elapsed`).               |
| `end_time`   | `number` | Unix timestamp. Shows "remaining" time (e.g., `12:00 remaining`).           |

---

## Important Notes

### Rate Limits

Discord has a rate limit for Rich Presence updates (usually **once every 15 seconds**). If you call
`rpc.update` too frequently, the changes might not appear or Discord may ignore them.

Always use `sleep(15000)` in your main loop:

```lua
while true do
    rpc.update({
        details = "Time is: " .. os.date("%H:%M:%S")
    })
    sleep(15000)
end
```

### Image Assets

You can use:

1. **Direct HTTPS URLs:** Easiest way, just paste a link to any image.
    <!-- prettier-ignore -->
    > [!WARNING]
    > **File Extensions Matter:** The URL must end with a valid image extension like `.png`, `.jpg`, `.jpeg`, or `.gif`. If the link doesn't have an extension (e.g., a redirect or a dynamic API link), Discord will **not** display the image.
2. **Asset Keys:** If you uploaded images to the
   [Discord Developer Portal](https://discord.com/developers/applications), use their names (keys)
   here.
