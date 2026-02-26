import { defineConfig } from "vitepress";

export default defineConfig({
    title: "LuaRPC",
    base: "/luarpc/",
    description: "Discord Rich Presence engine driven by Lua",
    head: [["link", { rel: "icon", href: "/luarpc/favicon.ico" }]],
    appearance: "force-dark",
    themeConfig: {
        nav: [
            { text: "Home", link: "/" },
            { text: "Guide", link: "/guide/getting-started" },
            { text: "API", link: "/api/globals" },
            { text: "Examples", link: "/examples/basic" }
        ],

        sidebar: [
            {
                text: "Introduction",
                items: [
                    { text: "Getting Started", link: "/guide/getting-started" },
                    { text: "CLI Reference", link: "/guide/cli" }
                ]
            },
            {
                text: "Runtime API",
                items: [
                    { text: "Global Functions & Variables", link: "/api/globals" },
                    { text: "RPC Module", link: "/api/rpc" }
                ]
            },
            {
                text: "Examples",
                items: [
                    { text: "Basic Examples", link: "/examples/basic" },
                    { text: "Advanced Examples", link: "/examples/advanced" }
                ]
            }
        ],

        socialLinks: [{ icon: "github", link: "https://github.com/asekirieshka/luarpc" }]
    }
});
