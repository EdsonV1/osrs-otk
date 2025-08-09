import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import { mockFetch, mockApiResponses } from '../utils';
import BirdhousesPage from '../../routes/tools/birdhouses/+page.svelte';

describe('Birdhouses Integration Flow', () => {
	const user = userEvent.setup();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('completes full calculation flow successfully', async () => {
		mockFetch('birdhouses', true);
		render(BirdhousesPage);

		// Verify page loads with correct title and form
		expect(screen.getByRole('heading', { name: /birdhouse calculator/i })).toBeInTheDocument();
		expect(screen.getByLabelText(/current hunter level/i)).toBeInTheDocument();

		// Fill out form
		const currentLevelInput = screen.getByLabelText(/current hunter level/i);
		const targetLevelInput = screen.getByLabelText(/target hunter level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '60');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '80');

		// Submit form
		await user.click(calculateButton);

		// Wait for results to appear
		await vi.waitFor(() => {
			expect(screen.getByText(/birdhouse training results/i)).toBeInTheDocument();
		});

		// Verify key metrics are displayed
		expect(screen.getByText('600,000')).toBeInTheDocument(); // Experience
		expect(screen.getByText('200')).toBeInTheDocument(); // Runs needed
		expect(screen.getByText(/4.0h/i)).toBeInTheDocument(); // Time

		// Verify page scrolls to results
		expect(Element.prototype.scrollIntoView).toHaveBeenCalled();
	});

	it('handles validation errors gracefully', async () => {
		render(BirdhousesPage);

		const currentLevelInput = screen.getByLabelText(/current hunter level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		// Enter invalid level
		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '4');
		await user.click(calculateButton);

		// Verify error message appears
		expect(screen.getByText(/current level must be at least 5/i)).toBeInTheDocument();
		
		// Verify results don't appear
		expect(screen.queryByText(/birdhouse training results/i)).not.toBeInTheDocument();
	});

	it('handles API errors with proper error display', async () => {
		mockFetch('birdhouses', false);
		render(BirdhousesPage);

		const currentLevelInput = screen.getByLabelText(/current hunter level/i);
		const targetLevelInput = screen.getByLabelText(/target hunter level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '60');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '80');
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