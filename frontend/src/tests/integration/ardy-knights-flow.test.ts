import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import { mockFetch } from '../utils';
import ArdyKnightsPage from '../../routes/tools/ardy-knights/+page.svelte';

describe('Ardy Knights Integration Flow', () => {
	const user = userEvent.setup();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('completes full calculation flow successfully', async () => {
		mockFetch('ardy-knights', true);
		render(ArdyKnightsPage);

		// Verify page loads with correct title and form
		expect(screen.getByRole('heading', { name: /ardougne knights calculator/i })).toBeInTheDocument();
		expect(screen.getByLabelText(/current thieving level/i)).toBeInTheDocument();

		// Fill out form
		const currentLevelInput = screen.getByLabelText(/current thieving level/i);
		const targetLevelInput = screen.getByLabelText(/target thieving level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '75');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '99');

		// Submit form
		await user.click(calculateButton);

		// Wait for results to appear
		await vi.waitFor(() => {
			expect(screen.getByText(/ardy knights training results/i)).toBeInTheDocument();
		});

		// Verify key metrics are displayed
		expect(screen.getByText('1,200,000')).toBeInTheDocument(); // Experience
		expect(screen.getByText(/180,000 gp\/hr/i)).toBeInTheDocument(); // GP per hour
		expect(screen.getByText(/5.0h/i)).toBeInTheDocument(); // Time

		// Verify page scrolls to results
		expect(Element.prototype.scrollIntoView).toHaveBeenCalled();
	});

	it('handles validation errors gracefully', async () => {
		render(ArdyKnightsPage);

		const currentLevelInput = screen.getByLabelText(/current thieving level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		// Enter invalid level
		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '54');
		await user.click(calculateButton);

		// Verify error message appears
		expect(screen.getByText(/current level must be at least 55/i)).toBeInTheDocument();
		
		// Verify results don't appear
		expect(screen.queryByText(/ardy knights training results/i)).not.toBeInTheDocument();
	});

	it('handles API errors with proper error display', async () => {
		mockFetch('ardy-knights', false);
		render(ArdyKnightsPage);

		const currentLevelInput = screen.getByLabelText(/current thieving level/i);
		const targetLevelInput = screen.getByLabelText(/target thieving level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '75');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '99');
		await user.click(calculateButton);

		// Wait for error to appear
		await vi.waitFor(() => {
			expect(screen.getByText(/calculation error/i)).toBeInTheDocument();
		});

		// Verify error styling
		const errorContainer = screen.getByText(/calculation error/i).closest('div');
		expect(errorContainer).toHaveClass('bg-gradient-to-r', 'from-red-500/10');
	});
});