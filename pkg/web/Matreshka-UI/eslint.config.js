import eslintPluginVue from "eslint-plugin-vue";
import eslintPluginImport from "eslint-plugin-import";
import eslintPluginPrettier from "eslint-plugin-prettier";
import typescriptParser from "@typescript-eslint/parser";
import typescriptPlugin from "@typescript-eslint/eslint-plugin";

export default [
  {
    files: ["**/*.ts", "**/*.vue"],
    languageOptions: {
      parser: typescriptParser,
      parserOptions: {
        ecmaVersion: "latest",
        sourceType: "module",
        project: "./tsconfig.json", // optional if using type-aware rules
      },
    },
    plugins: {
      vue: eslintPluginVue,
      import: eslintPluginImport,
      prettier: eslintPluginPrettier,
      "@typescript-eslint": typescriptPlugin,
    },
    rules: {
      quotes: ["error", "double"],
      "import/order": [
        "error",
        {
          groups: ["builtin", "external", "internal"],
          alphabetize: { order: "asc", caseInsensitive: true },
          "newlines-between": "always",
        }
      ],
      "prettier/prettier": "error",

      // Optional TypeScript-specific rules
      "@typescript-eslint/explicit-module-boundary-types": "off",
      "@typescript-eslint/no-unused-vars": ["warn", { argsIgnorePattern: "^_" }],
    },
  },
];
