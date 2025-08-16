import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import { mockFetch, mockApiResponses } from '../utils';
import WintertodtPage from '../../routes/tools/wintertodt/+page.svelte';

describe('Wintertodt Integration Flow', () => {
	const user = userEvent.setup();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('completes full calculation flow successfully', async () => {
		mockFetch('wintertodt', true);
		render(WintertodtPage);

		// Verify page loads with correct title and form
		expect(screen.getByRole('heading', { name: /wintertodt calculator/i })).toBeInTheDocument();
		expect(screen.getByLabelText(/firemaking level/i)).toBeInTheDocument();

		// Fill out form
		const levelInput = screen.getByLabelText(/firemaking level/i);
		const roundsInput = screen.getByLabelText(/total rounds/i);
		const calculateButton = screen.getByRole('button', { name: /calculate/i });

		await user.clear(levelInput);
		await user.type(levelInput, '85');
		await user.clear(roundsInput);
		await user.type(roundsInput, '100');

		// Submit form
		await user.click(calculateButton);

		// Wait for results to appear
		await vi.waitFor(() => {
			expect(screen.getByText(/wintertodt results/i)).toBeInTheDocument();
		});

		// Verify key metrics are displayed
		expect(screen.getByText('500,000')).toBeInTheDocument(); // Experience
		expect(screen.getByText('12.5%')).toBeInTheDocument(); // Pet chance
		expect(screen.getByText(/2.5M gp/i)).toBeInTheDocument(); // Loot value

		// Verify page scrolls to results
		expect(Element.prototype.scrollIntoView).toHaveBeenCalled();
	});

	it('handles validation errors gracefully', async () => {
		render(WintertodtPage);

		const levelInput = screen.getByLabelText(/firemaking level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate/i });

		// Enter invalid level
		await user.clear(levelInput);
		await user.type(levelInput, '30');
		await user.click(calculateButton);

		// Verify error message appears
		expect(screen.getByText(/firemaking level must be at least 50/i)).toBeInTheDocument();
		
		// Verify results don't appear
		expect(screen.queryByText(/wintertodt results/i)).not.toBeInTheDocument();
	});

	it('handles API errors with proper error display', async () => {
		mockFetch('wintertodt', false);
		render(WintertodtPage);

		const levelInput = screen.getByLabelText(/firemaking level/i);
		const roundsInput = screen.getByLabelText(/total rounds/i);
		const calculateButton = screen.getByRole('button', { name: /calculate/i });

		await user.clear(levelInput);
		await user.type(levelInput, '85');
		await user.clear(roundsInput);
		await user.type(roundsInput, '100');
		await user.click(calculateButton);

		// Wait for error to appear
		await vi.waitFor(() => {
			expect(screen.getByText(/calculation error/i)).toBeInTheDocument();
		});

		// Verify error styling
		const errorContainer = screen.getByText(/calculation error/i).closest('div');
		expect(errorContainer).toHaveClass('bg-gradient-to-r', 'from-red-500/10');
	});

	it('shows loading state during calculation', async () => {
		// Mock slow API response
		vi.mocked(fetch).mockImplementation(() => 
			new Promise(resolve => 
				setTimeout(() => resolve({
					ok: true,
					json: async () => mockApiResponses.wintertodt
				} as Response), 100)
			)
		);

		render(WintertodtPage);

		const calculateButton = screen.getByRole('button', { name: /calculate/i });
		await user.click(calculateButton);

		// Verify loading state
		expect(calculateButton).toBeDisabled();
		expect(screen.getByText(/calculating/i)).toBeInTheDocument();

		// Wait for completion
		await vi.waitFor(() => {
			expect(screen.getByText(/wintertodt results/i)).toBeInTheDocument();
		}, { timeout: 200 });
	});

	it('navigates back to tools page correctly', async () => {
		render(WintertodtPage);

		const backLink = screen.getByRole('link', { name: /back to all tools/i });
		expect(backLink).toHaveAttribute('href', '/tools');
	});
});