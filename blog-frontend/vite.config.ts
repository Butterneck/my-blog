import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config.js';
import basicSSL from '@vitejs/plugin-basic-ssl';

export default defineConfig({
	plugins: [sveltekit(), basicSSL()],
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	},
	optimizeDeps: { exclude: ["bytemd"]},
});
