import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import { mockApiResponses } from '../../../tests/utils';
import ResultsDisplay from './ResultsDisplay.svelte';

describe('GOTR ResultsDisplay', () => {
	const mockResult = mockApiResponses.gotr;

	it('renders all key metrics correctly', () => {
		render(ResultsDisplay, {
			apiResult: mockResult
		});

		// Check experience display
		expect(screen.getByText('2,800,000')).toBeInTheDocument();
		expect(screen.getByText('45,000/hr')).toBeInTheDocument();

		// Check games needed
		expect(screen.getByText('175')).toBeInTheDocument();

		// Check time display
		expect(screen.getByText(/62.2h/i)).toBeInTheDocument();
	});

	it('displays loot breakdown correctly', () => {
		render(ResultsDisplay, {
			apiResult: mockResult
		});

		expect(screen.getByText(/catalytic_talisman/i)).toBeInTheDocument();
		expect(screen.getByText(/elemental_talisman/i)).toBeInTheDocument();
		expect(screen.getByText(/lantern/i)).toBeInTheDocument();
		expect(screen.getByText('8')).toBeInTheDocument();
		expect(screen.getByText('12')).toBeInTheDocument();
		expect(screen.getByText('3')).toBeInTheDocument();
	});

	it('uses custom icons when provided', () => {
		render(ResultsDisplay, {
			apiResult: mockResult,
			iconSrc: '/custom/icon.png',
			experienceIconSrc: '/custom/xp.png'
		});

		const icons = screen.getAllByRole('img');
		const mainIcon = icons.find(img => img.getAttribute('alt')?.includes('GOTR'));
		const xpIcon = icons.find(img => img.getAttribute('alt')?.includes('Experience'));

		expect(mainIcon).toHaveAttribute('src', '/custom/icon.png');
		expect(xpIcon).toHaveAttribute('src', '/custom/xp.png');
	});

	it('formats large numbers correctly', () => {
		const largeResult = {
			...mockResult,
			total_experience: 15000000,
			games_needed: 2500
		};

		render(ResultsDisplay, {
			apiResult: largeResult
		});

		expect(screen.getByText('15,000,000')).toBeInTheDocument();
		expect(screen.getByText('2,500')).toBeInTheDocument();
	});
});