# CLI Reference

LuaRPC is a **Zero Config** engine. You don't need complex JSON or YAML files to get
started—everything is handled through the command line and your Lua scripts.

---

## luarpc init

Initializes a new project workspace.

```bash
# Create project in a new folder
luarpc init my-rpc-project

# Initialize current directory
luarpc init .
```

**What it does:**

- Generates `main.lua` (your logic).
- Generates `definitions.lua` (IDE stubs for autocomplete).

<!-- prettier-ignore -->
> [!TIP]
> Since LuaRPC is zero-config, the `init` command is only needed to give you a starting point. You can manually create a `.lua` file and it will work just fine.

---

## luarpc run

Starts the Discord Rich Presence engine using your script.

```bash
luarpc run main.lua
```

**Key details:**

- **Hot Reloading:** Currently, there is no hot reloading and you need to restart the command
  (`Ctrl + C` and run again) to apply changes to your Lua code.
- **Execution:** The engine runs the Lua script in a continuous loop (if defined) or executes it
  once.
- **Logs:** Any `print()` calls in your Lua script will appear directly in your terminal.

---

## luarpc version

Checks your current installation.

```bash
luarpc version
```

If you see a version older than the one on GitHub, simply run
`go install github.com/asekirieshka/luarpc@latest` to update.

---

## Technical Details

### Why definitions.lua?

Even though we are "Zero Config", we care about developer experience. The `definitions.lua` file is
not a configuration file—it's a **Type Definition** file. It tells your IDE (VS Code, Zed, Neovim)
that functions like `rpc.update` exist, so you get perfect autocomplete while coding.

### Path Resolution

When you run `luarpc run script.lua`, the engine sets the global `__dirname` variable to the
absolute path of that script. This allows you to load local files reliably regardless of where you
call the command from.
