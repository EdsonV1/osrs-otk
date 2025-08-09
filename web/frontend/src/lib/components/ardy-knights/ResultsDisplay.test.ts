import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import { mockApiResponses } from '../../../tests/utils';
import ResultsDisplay from './ResultsDisplay.svelte';

describe('Ardy Knights ResultsDisplay', () => {
	const mockResult = mockApiResponses['ardy-knights'];

	it('renders all key metrics correctly', () => {
		render(ResultsDisplay, {
			apiResult: mockResult
		});

		// Check experience display
		expect(screen.getByText('1,200,000')).toBeInTheDocument();
		expect(screen.getByText('240,000/hr')).toBeInTheDocument();

		// Check gp/hr
		expect(screen.getByText(/180,000 gp\/hr/i)).toBeInTheDocument();

		// Check time display
		expect(screen.getByText(/5.0h/i)).toBeInTheDocument();
	});

	it('displays loot breakdown correctly', () => {
		render(ResultsDisplay, {
			apiResult: mockResult
		});

		expect(screen.getByText(/coin_pouch/i)).toBeInTheDocument();
		expect(screen.getByText(/gems/i)).toBeInTheDocument();
		expect(screen.getByText('1500')).toBeInTheDocument();
		expect(screen.getByText('45')).toBeInTheDocument();
	});

	it('uses custom icons when provided', () => {
		render(ResultsDisplay, {
			apiResult: mockResult,
			iconSrc: '/custom/icon.png',
			experienceIconSrc: '/custom/xp.png'
		});

		const icons = screen.getAllByRole('img');
		const mainIcon = icons.find(img => img.getAttribute('alt')?.includes('Ardy Knights'));
		const xpIcon = icons.find(img => img.getAttribute('alt')?.includes('Experience'));

		expect(mainIcon).toHaveAttribute('src', '/custom/icon.png');
		expect(xpIcon).toHaveAttribute('src', '/custom/xp.png');
	});

	it('formats large numbers correctly', () => {
		const largeResult = {
			...mockResult,
			total_experience: 25000000,
			total_gp: 50000000
		};

		render(ResultsDisplay, {
			apiResult: largeResult
		});

		expect(screen.getByText('25,000,000')).toBeInTheDocument();
		expect(screen.getByText(/50.0M gp/i)).toBeInTheDocument();
	});
});