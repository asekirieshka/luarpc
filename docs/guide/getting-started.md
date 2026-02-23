# Getting Started

Welcome to **LuaRPC**! This guide will help you set up your first Discord Rich Presence script using
the Go toolchain.

## 1. Installation

The easiest way to install LuaRPC is via `go install`. This will automatically compile the binary
and add it to your shell's PATH.

```bash
go install github.com/asekirieshka/luarpc@latest
```

> **Prerequisite:** Make sure you have [Go](https://go.dev/dl/) installed and your `$GOPATH/bin` (or
> `%USERPROFILE%\go\bin` on Windows) is in your system's PATH.

---

## 2. Initialization

Create your project by running the init command. We recommend providing a folder name to keep things
organized:

```bash
luarpc init my-rpc-project
```

Then navigate into the created directory:

```bash
cd my-rpc-project
```

<!-- prettier-ignore -->
> [!TIP]
> If you already created a folder and are inside it, you can use `luarpc init .` to
> initialize the current directory.

---

## 3. Project Structure

The initialization command generates the following files:

- **main.lua**: Your script entry point.
- **definitions.lua**: Helper file for IDE autocomplete (LSP). **Do not delete this!**

---

## 4. Setup Client ID

Open `main.lua` and replace the placeholder ID with your **Discord Application Client ID**:

```lua
rpc.set_client_id("1234567891234567891")
```

If you don't have one, create an app at the
[Discord Developer Portal](https://discord.com/developers/applications).

---

## 5. Run it

Start your presence engine by running:

```bash
luarpc run main.lua
```

Your Discord status should update immediately. To stop the process, press `Ctrl + C`.

<!-- prettier-ignore -->
> [!CAUTION]
> **Troubleshooting: Status not showing?**
> 1. **Discord must be running:** The engine communicates with the local Discord client. Make sure Discord is open on your machine.
> 2. **Activity Status:** Go to Discord **Settings > Activity Privacy** and ensure that "Share my activity" is toggled **ON**.
> 3. **App Status:** Ensure you haven't set your status to "Invisible" in Discord.

---

## 6. Next Steps

- Explore the **[RPC API](../api/rpc)** for advanced status customization.
- Check out **[Global Functions](../api/globals)** for HTTP requests and timers.
