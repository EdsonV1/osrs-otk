import '@testing-library/jest-dom';
import { vi } from 'vitest';

// Mock IntersectionObserver for components that use it
global.IntersectionObserver = vi.fn(() => ({
	disconnect: vi.fn(),
	observe: vi.fn(),
	unobserve: vi.fn(),
})) as any;

// Mock ResizeObserver for responsive components
global.ResizeObserver = vi.fn(() => ({
	disconnect: vi.fn(),
	observe: vi.fn(),
	unobserve: vi.fn(),
})) as any;

// Mock fetch for API calls
global.fetch = vi.fn();

// Mock window.matchMedia for responsive design tests
Object.defineProperty(window, 'matchMedia', {
	writable: true,
	value: vi.fn().mockImplementation(query => ({
		matches: false,
		media: query,
		onchange: null,
		addListener: vi.fn(), // deprecated
		removeListener: vi.fn(), // deprecated
		addEventListener: vi.fn(),
		removeEventListener: vi.fn(),
		dispatchEvent: vi.fn(),
	})),
});

// Mock scrollIntoView for smooth scrolling tests
Element.prototype.scrollIntoView = vi.fn();

// Mock $app/environment for SvelteKit
vi.mock('$app/environment', () => ({
	browser: false,
	dev: true,
	building: false,
	version: '1.0.0'
}));

// Mock $app/stores for SvelteKit
vi.mock('$app/stores', () => {
	const readable = vi.fn(() => ({
		subscribe: vi.fn(() => vi.fn())
	}));
	
	return {
		page: readable({
			url: new URL('http://localhost:3000'),
			params: {},
			route: { id: null },
			status: 200,
			error: null,
			data: {},
			form: undefined
		}),
		navigating: readable(null),
		updated: {
			subscribe: vi.fn(() => vi.fn()),
			check: vi.fn()
		}
	};
});