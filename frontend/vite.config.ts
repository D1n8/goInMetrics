import react from '@vitejs/plugin-react';
import { defineConfig } from 'vite';
import { readFileSync } from 'node:fs';
import { parse } from 'smol-toml';

const toml = parse(readFileSync('../config.toml', 'utf-8')) as {
  server: {
    FRONTEND_PORT: number;
  }
}
// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: toml.server.FRONTEND_PORT ?? 5173,
  }
});
