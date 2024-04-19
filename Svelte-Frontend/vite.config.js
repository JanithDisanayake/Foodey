import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [sveltekit()],
  optimizeDeps: {
    exclude: ["@urql/svelte"],
  },
  server: {
    proxy: {
      "/query": {
        target: "http://localhost:3000/",
        changeOrigin: true,
        pathRewrite: {
          "^/query": "/query",
        },
      },
    },
  },
});
