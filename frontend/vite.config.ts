import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
			'/api': "http://backend-app-1:8082"
		}
	}
});
