import { defineConfig } from "vitest/config";

export default defineConfig({
  resolve: { tsconfigPaths: true },
  test: {
    exclude: ["node_modules/**"],
    passWithNoTests: true,
    environment: "jsdom",
  },
});
