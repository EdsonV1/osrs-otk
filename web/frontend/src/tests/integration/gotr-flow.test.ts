import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import { mockFetch, mockApiResponses } from '../utils';
import GotrPage from '../../routes/tools/gotr/+page.svelte';

describe('GOTR Integration Flow', () => {
	const user = userEvent.setup();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('completes full calculation flow successfully', async () => {
		mockFetch('gotr', true);
		render(GotrPage);

		// Verify page loads with correct title and form
		expect(screen.getByRole('heading', { name: /guardians of the rift calculator/i })).toBeInTheDocument();
		expect(screen.getByLabelText(/current runecrafting level/i)).toBeInTheDocument();

		// Fill out form
		const currentLevelInput = screen.getByLabelText(/current runecrafting level/i);
		const targetLevelInput = screen.getByLabelText(/target runecrafting level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '77');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '99');

		// Submit form
		await user.click(calculateButton);

		// Wait for results to appear
		await vi.waitFor(() => {
			expect(screen.getByText(/gotr training results/i)).toBeInTheDocument();
		});

		// Verify key metrics are displayed
		expect(screen.getByText('2,800,000')).toBeInTheDocument(); // Experience
		expect(screen.getByText('175')).toBeInTheDocument(); // Games needed
		expect(screen.getByText(/62.2h/i)).toBeInTheDocument(); // Time

		// Verify page scrolls to results
		expect(Element.prototype.scrollIntoView).toHaveBeenCalled();
	});

	it('handles validation errors gracefully', async () => {
		render(GotrPage);

		const currentLevelInput = screen.getByLabelText(/current runecrafting level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		// Enter invalid level
		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '26');
		await user.click(calculateButton);

		// Verify error message appears
		expect(screen.getByText(/current level must be at least 27/i)).toBeInTheDocument();
		
		// Verify results don't appear
		expect(screen.queryByText(/gotr training results/i)).not.toBeInTheDocument();
	});

	it('handles API errors with proper error display', async () => {
		mockFetch('gotr', false);
		render(GotrPage);

		const currentLevelInput = screen.getByLabelText(/current runecrafting level/i);
		const targetLevelInput = screen.getByLabelText(/target runecrafting level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '77');
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