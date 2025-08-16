import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import { mockFetch, mockApiResponses } from '../../../tests/utils';
import InputForm from './InputForm.svelte';

describe('Wintertodt InputForm', () => {
	const user = userEvent.setup();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('renders form elements correctly', () => {
		render(InputForm);

		expect(screen.getByLabelText(/firemaking level/i)).toBeInTheDocument();
		expect(screen.getByLabelText(/total rounds/i)).toBeInTheDocument();
		expect(screen.getByRole('button', { name: /calculate/i })).toBeInTheDocument();
	});

	it('shows validation error for invalid firemaking level', async () => {
		render(InputForm);

		const levelInput = screen.getByLabelText(/firemaking level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate/i });

		await user.clear(levelInput);
		await user.type(levelInput, '49'); // Below minimum level
		await user.click(calculateButton);

		expect(screen.getByText(/firemaking level must be at least 50/i)).toBeInTheDocument();
	});

	it('shows validation error for invalid rounds', async () => {
		render(InputForm);

		const roundsInput = screen.getByLabelText(/total rounds/i);
		const calculateButton = screen.getByRole('button', { name: /calculate/i });

		await user.clear(roundsInput);
		await user.type(roundsInput, '0');
		await user.click(calculateButton);

		expect(screen.getByText(/total rounds must be at least 1/i)).toBeInTheDocument();
	});

	it('submits form with valid data and emits calculated event', async () => {
		mockFetch('wintertodt', true);

		const { component } = render(InputForm);
		const calculatedSpy = vi.fn();
		component.$on('calculated', calculatedSpy);

		const levelInput = screen.getByLabelText(/firemaking level/i);
		const roundsInput = screen.getByLabelText(/total rounds/i);
		const calculateButton = screen.getByRole('button', { name: /calculate/i });

		await user.clear(levelInput);
		await user.type(levelInput, '85');
		await user.clear(roundsInput);
		await user.type(roundsInput, '100');
		await user.click(calculateButton);

		// Wait for async API call
		await vi.waitFor(() => {
			expect(calculatedSpy).toHaveBeenCalledWith(
				expect.objectContaining({
					detail: {
						resultData: mockApiResponses.wintertodt
					}
				})
			);
		});
	});

	it('handles API errors and emits error event', async () => {
		mockFetch('wintertodt', false);

		const { component } = render(InputForm);
		const errorSpy = vi.fn();
		component.$on('error', errorSpy);

		const levelInput = screen.getByLabelText(/firemaking level/i);
		const roundsInput = screen.getByLabelText(/total rounds/i);
		const calculateButton = screen.getByRole('button', { name: /calculate/i });

		await user.clear(levelInput);
		await user.type(levelInput, '85');
		await user.clear(roundsInput);
		await user.type(roundsInput, '100');
		await user.click(calculateButton);

		await vi.waitFor(() => {
			expect(errorSpy).toHaveBeenCalledWith(
				expect.objectContaining({
					detail: {
						message: expect.stringContaining('error')
					}
				})
			);
		});
	});

	it('disables calculate button while loading', async () => {
		// Mock a slow API response
		vi.mocked(fetch).mockImplementation(() => 
			new Promise(resolve => 
				setTimeout(() => resolve({
					ok: true,
					json: async () => mockApiResponses.wintertodt
				} as Response), 100)
			)
		);

		render(InputForm);

		const calculateButton = screen.getByRole('button', { name: /calculate/i });
		
		await user.click(calculateButton);
		
		expect(calculateButton).toBeDisabled();
		expect(screen.getByText(/calculating/i)).toBeInTheDocument();
	});
});