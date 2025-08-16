import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import { mockApiResponses } from '../../../tests/utils';
import ResultsDisplay from './ResultsDisplay.svelte';

describe('Wintertodt ResultsDisplay', () => {
	const mockResult = mockApiResponses.wintertodt;

	it('renders all key metrics correctly', () => {
		render(ResultsDisplay, {
			apiResult: mockResult
		});

		// Check experience display
		expect(screen.getByText('500,000')).toBeInTheDocument();
		expect(screen.getByText('180,000/hr')).toBeInTheDocument();

		// Check pet chance
		expect(screen.getByText('12.5%')).toBeInTheDocument();

		// Check loot value
		expect(screen.getByText(/2.5M gp/i)).toBeInTheDocument();

		// Check time display
		expect(screen.getByText(/2.8h/i)).toBeInTheDocument();
	});

	it('displays loot breakdown correctly', () => {
		render(ResultsDisplay, {
			apiResult: mockResult
		});

		expect(screen.getByText(/burnt_page/i)).toBeInTheDocument();
		expect(screen.getByText(/supply_crate/i)).toBeInTheDocument();
		expect(screen.getByText(/magic_logs/i)).toBeInTheDocument();
		expect(screen.getByText('100')).toBeInTheDocument();
		expect(screen.getByText('250')).toBeInTheDocument();
	});

	it('uses custom icons when provided', () => {
		render(ResultsDisplay, {
			apiResult: mockResult,
			iconSrc: '/custom/icon.png',
			experienceIconSrc: '/custom/xp.png'
		});

		const icons = screen.getAllByRole('img');
		const mainIcon = icons.find(img => img.getAttribute('alt')?.includes('Wintertodt'));
		const xpIcon = icons.find(img => img.getAttribute('alt')?.includes('Experience'));

		expect(mainIcon).toHaveAttribute('src', '/custom/icon.png');
		expect(xpIcon).toHaveAttribute('src', '/custom/xp.png');
	});

	it('formats large numbers correctly', () => {
		const largeResult = {
			...mockResult,
			total_experience: 15000000,
			total_value: 125000000
		};

		render(ResultsDisplay, {
			apiResult: largeResult
		});

		expect(screen.getByText('15,000,000')).toBeInTheDocument();
		expect(screen.getByText(/125.0M gp/i)).toBeInTheDocument();
	});

	it('shows correct pet chance color coding', () => {
		// Test high pet chance (green)
		const highChanceResult = { ...mockResult, pet_chance: 75 };
		const { rerender } = render(ResultsDisplay, {
			apiResult: highChanceResult
		});

		let petChanceElement = screen.getByText('75.00%');
		expect(petChanceElement).toHaveClass('text-green-400');

		// Test medium pet chance (orange)
		const mediumChanceResult = { ...mockResult, pet_chance: 30 };
		rerender({ apiResult: mediumChanceResult });

		petChanceElement = screen.getByText('30.00%');
		expect(petChanceElement).toHaveClass('text-orange-400');

		// Test low pet chance (red)
		const lowChanceResult = { ...mockResult, pet_chance: 10 };
		rerender({ apiResult: lowChanceResult });

		petChanceElement = screen.getByText('10.00%');
		expect(petChanceElement).toHaveClass('text-red-400');
	});
});