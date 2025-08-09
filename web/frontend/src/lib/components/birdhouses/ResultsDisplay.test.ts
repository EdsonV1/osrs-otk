import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import { mockApiResponses } from '../../../tests/utils';
import ResultsDisplay from './ResultsDisplay.svelte';

describe('Birdhouses ResultsDisplay', () => {
	const mockResult = mockApiResponses.birdhouses;

	it('renders all key metrics correctly', () => {
		render(ResultsDisplay, {
			apiResult: mockResult
		});

		// Check experience display
		expect(screen.getByText('600,000')).toBeInTheDocument();
		expect(screen.getByText('150,000/hr')).toBeInTheDocument();

		// Check runs needed
		expect(screen.getByText('200')).toBeInTheDocument();

		// Check time display
		expect(screen.getByText(/4.0h/i)).toBeInTheDocument();
	});

	it('displays loot breakdown correctly', () => {
		render(ResultsDisplay, {
			apiResult: mockResult
		});

		expect(screen.getByText(/bird_nest/i)).toBeInTheDocument();
		expect(screen.getByText(/raw_bird_meat/i)).toBeInTheDocument();
		expect(screen.getByText('50')).toBeInTheDocument();
		expect(screen.getByText('800')).toBeInTheDocument();
	});

	it('uses custom icons when provided', () => {
		render(ResultsDisplay, {
			apiResult: mockResult,
			iconSrc: '/custom/icon.png',
			experienceIconSrc: '/custom/xp.png'
		});

		const icons = screen.getAllByRole('img');
		const mainIcon = icons.find(img => img.getAttribute('alt')?.includes('Birdhouses'));
		const xpIcon = icons.find(img => img.getAttribute('alt')?.includes('Experience'));

		expect(mainIcon).toHaveAttribute('src', '/custom/icon.png');
		expect(xpIcon).toHaveAttribute('src', '/custom/xp.png');
	});

	it('formats large numbers correctly', () => {
		const largeResult = {
			...mockResult,
			total_experience: 10000000,
			runs_needed: 5000
		};

		render(ResultsDisplay, {
			apiResult: largeResult
		});

		expect(screen.getByText('10,000,000')).toBeInTheDocument();
		expect(screen.getByText('5,000')).toBeInTheDocument();
	});
});