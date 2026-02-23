---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
    name: "LuaRPC"
    text: "Discord Presence on Steroids"
    tagline: "A lightweight Go engine where your status is governed by Lua magic."
    image:
        src: /luarpc-icon.svg
        alt: LuaRPC Logo
    actions:
        - theme: brand
          text: Get Started
          link: /guide/getting-started
        - theme: alt
          text: API Reference
          link: /api/rpc
        - theme: alt
          text: View on GitHub
          link: https://github.com/asekirieshka/luarpc

features:
    - icon: 🚀
      title: Instant Launch
      details: Zero complex setup. Run `luarpc init` and you're in the game.
    - icon: 📜
      title: Powered by Lua
      details:
          Full control via Lua 5.1. Dynamic statuses, loops, conditions, and absolute flexibility
          without compilation.
    - icon: 🌐
      title: Built-in HTTP & JSON
      details:
          Connect to GitHub, Steam, or even Bitcoin exchange APIs. Your status can display anything
          that exists on the web.
    - icon: 🛠️
      title: Developer Friendly
      details:
          Included definitions.lua provides perfect IDE autocomplete (LSP) in VS Code. Writing
          scripts is a breeze.
    - icon: 📦
      title: Single Binary
      details:
          The entire runtime is packed into one executable. No external dependencies required except
          Discord itself.
    - icon: 🕊️
      title: Light as Air
      details:
          Written in Go with efficiency in mind. Consumes minimal resources while you focus on what
          really matters.
---
