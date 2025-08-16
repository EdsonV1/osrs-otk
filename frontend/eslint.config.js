import js from '@eslint/js';
import ts from '@typescript-eslint/eslint-plugin';
import tsParser from '@typescript-eslint/parser';
import svelte from 'eslint-plugin-svelte';
import svelteParser from 'svelte-eslint-parser';
import globals from 'globals';

export default [
	js.configs.recommended,
	{
		files: ['**/*.{js,mjs,cjs,ts}'],
		languageOptions: {
			globals: {
				...globals.browser,
				...globals.node,
			},
		},
	},
	{
		files: ['**/*.{ts,tsx}'],
		ignores: ['**/*.test.ts', '**/*.spec.ts', 'src/tests/**'],
		languageOptions: {
			parser: tsParser,
			parserOptions: {
				ecmaVersion: 2020,
				sourceType: 'module',
				project: './tsconfig.json',
			},
		},
		plugins: {
			'@typescript-eslint': ts,
		},
		rules: {
			...ts.configs.recommended.rules,
			'@typescript-eslint/no-unused-vars': ['error', { argsIgnorePattern: '^_' }],
			'@typescript-eslint/no-explicit-any': 'warn',
		},
	},
	{
		files: ['**/*.test.{ts,tsx}', '**/*.spec.{ts,tsx}', 'src/tests/**/*.ts'],
		languageOptions: {
			parser: tsParser,
			parserOptions: {
				ecmaVersion: 2020,
				sourceType: 'module',
				// Don't use project for test files
			},
			globals: {
				...globals.browser,
				...globals.node,
				describe: 'readonly',
				it: 'readonly',
				expect: 'readonly',
				beforeEach: 'readonly',
				afterEach: 'readonly',
				vi: 'readonly',
			},
		},
		plugins: {
			'@typescript-eslint': ts,
		},
		rules: {
			...ts.configs.recommended.rules,
			'@typescript-eslint/no-unused-vars': ['error', { argsIgnorePattern: '^_' }],
			'@typescript-eslint/no-explicit-any': 'off',
		},
	},
	{
		files: ['**/*.svelte'],
		languageOptions: {
			parser: svelteParser,
			parserOptions: {
				parser: tsParser,
				extraFileExtensions: ['.svelte'],
				ecmaVersion: 2020,
				sourceType: 'module',
			},
			globals: {
				...globals.browser,
				fetch: 'readonly',
				console: 'readonly',
				document: 'readonly',
				window: 'readonly',
				HTMLElement: 'readonly',
				HTMLImageElement: 'readonly',
				CustomEvent: 'readonly',
				setTimeout: 'readonly',
				clearTimeout: 'readonly',
			},
		},
		plugins: {
			svelte,
			'@typescript-eslint': ts,
		},
		rules: {
			...svelte.configs.recommended.rules,
			// Svelte-specific rules
			'svelte/no-at-debug-tags': 'warn',
			'svelte/no-at-html-tags': 'error',
			// Allow unused variables that start with underscore
			'no-unused-vars': ['error', { argsIgnorePattern: '^_', varsIgnorePattern: '^_' }],
			'@typescript-eslint/no-unused-vars': ['error', { argsIgnorePattern: '^_', varsIgnorePattern: '^_' }],
		},
	},
	{
		ignores: [
			'build/',
			'.svelte-kit/',
			'dist/',
			'node_modules/',
			'coverage/',
			'*.config.js',
			'*.config.ts',
			// Temporarily ignore work-in-progress skill pages
			'src/routes/skills/**',
		],
	},
];