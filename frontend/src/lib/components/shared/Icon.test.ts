import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import Icon from './Icon.svelte';

describe('Icon Component', () => {
	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('renders with correct src and alt attributes', () => {
		render(Icon, {
			src: '/skills/firemaking.png',
			alt: 'Firemaking skill icon'
		});

		const img = screen.getByRole('img', { name: 'Firemaking skill icon' });
		expect(img).toBeInTheDocument();
		expect(img).toHaveAttribute('src', '/skills/firemaking.png');
	});

	it('applies correct size classes', () => {
		render(Icon, {
			src: '/test.png',
			alt: 'test',
			size: 'lg'
		});

		const img = screen.getByRole('img');
		expect(img).toHaveClass('w-10', 'h-10');
	});

	it('applies custom classes', () => {
		render(Icon, {
			src: '/test.png',
			alt: 'test',
			classes: 'custom-class'
		});

		const img = screen.getByRole('img');
		expect(img).toHaveClass('custom-class');
	});

	it('handles lazy loading by default', () => {
		render(Icon, {
			src: '/test.png',
			alt: 'test'
		});

		const img = screen.getByRole('img');
		expect(img).toHaveAttribute('loading', 'lazy');
	});

	it('handles eager loading when specified', () => {
		render(Icon, {
			src: '/test.png',
			alt: 'test',
			loading: 'eager'
		});

		const img = screen.getByRole('img');
		expect(img).toHaveAttribute('loading', 'eager');
	});

	it('sets correct fetchpriority for preloaded icons', () => {
		render(Icon, {
			src: '/test.png',
			alt: 'test',
			preload: true
		});

		const img = screen.getByRole('img');
		expect(img).toHaveAttribute('fetchpriority', 'high');
	});

	it('shows fallback image on error when provided', async () => {
		render(Icon, {
			src: '/broken.png',
			alt: 'test',
			fallback: '/fallback.png'
		});

		const img = screen.getByRole('img');
		
		// Simulate image load error
		await img.dispatchEvent(new Event('error'));
		
		expect(img).toHaveAttribute('src', '/fallback.png');
	});
});