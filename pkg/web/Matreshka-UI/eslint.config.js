import eslintPluginVue from "eslint-plugin-vue";
import eslintPluginImport from "eslint-plugin-import";
import eslintPluginPrettier from "eslint-plugin-prettier";
import tsParser from "@typescript-eslint/parser";
import typescriptPlugin from "@typescript-eslint/eslint-plugin";
import simpleImportSort from "eslint-plugin-simple-import-sort";
import vueParser from "vue-eslint-parser";

export default [
  {
    files: ["src/**/*.vue"],
    languageOptions: {
      parser: vueParser,
      parserOptions: {
        parser: tsParser,
        ecmaVersion: "latest",
        sourceType: "module",
        project: "./tsconfig.json",
        extraFileExtensions: [".vue"],
      },
    },
    plugins: {
      vue: eslintPluginVue,
      import: eslintPluginImport,
      prettier: eslintPluginPrettier,
      "@typescript-eslint": typescriptPlugin,
      "simple-import-sort": simpleImportSort,
    },
    rules: {
      quotes: ["error", "double"],
      "simple-import-sort/imports": "error",
      "simple-import-sort/exports": "error",
      "prettier/prettier": "error",
      "vue/no-unused-vars": "warn",

      "@typescript-eslint/explicit-module-boundary-types": "off",
      "@typescript-eslint/no-unused-vars": ["warn", { argsIgnorePattern: "^_" }],
    },
  },
  {
    files: ["src/**/*.ts"],
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        ecmaVersion: "latest",
        sourceType: "module",
        project: "./tsconfig.json", // optional if using type-aware rules
        extraFileExtensions: [".vue"],
      },
    },
    plugins: {
      vue: eslintPluginVue,
      import: eslintPluginImport,
      prettier: eslintPluginPrettier,
      "@typescript-eslint": typescriptPlugin,
      "simple-import-sort": simpleImportSort,
    },
    rules: {
      quotes: ["error", "double"],
      "simple-import-sort/imports": "error",
      "simple-import-sort/exports": "error",
      "prettier/prettier": "error",

      // Optional TypeScript-specific rules
      "@typescript-eslint/explicit-module-boundary-types": "off",
      "@typescript-eslint/no-unused-vars": ["warn", { argsIgnorePattern: "^_" }],
    },
  },
];
