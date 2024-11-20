import { defineConfig } from "vitest/config";
import { sveltekit } from "@sveltejs/kit/vite";

export default defineConfig({
	plugins: [sveltekit()],

	server: {
		host: true,
		port: 3000,
		proxy: {
			"/api": process.env.VITE_API_URL || "http://localhost:8080",
		},
	},

	test: {
		include: ["src/**/*.{test,spec}.{js,ts}"],
	},
});
