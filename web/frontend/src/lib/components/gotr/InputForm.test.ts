import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import { mockFetch, mockApiResponses } from '../../../tests/utils';
import InputForm from './InputForm.svelte';

describe('GOTR InputForm', () => {
	const user = userEvent.setup();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('renders form elements correctly', () => {
		render(InputForm);

		expect(screen.getByLabelText(/current runecrafting level/i)).toBeInTheDocument();
		expect(screen.getByLabelText(/target runecrafting level/i)).toBeInTheDocument();
		expect(screen.getByRole('button', { name: /calculate training plan/i })).toBeInTheDocument();
	});

	it('shows validation error for level below minimum', async () => {
		render(InputForm);

		const currentLevelInput = screen.getByLabelText(/current runecrafting level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '26'); // Below minimum level
		await user.click(calculateButton);

		expect(screen.getByText(/current level must be at least 27/i)).toBeInTheDocument();
	});

	it('shows validation error for target level lower than current', async () => {
		render(InputForm);

		const currentLevelInput = screen.getByLabelText(/current runecrafting level/i);
		const targetLevelInput = screen.getByLabelText(/target runecrafting level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '80');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '75');
		await user.click(calculateButton);

		expect(screen.getByText(/target level must be higher than current level/i)).toBeInTheDocument();
	});

	it('submits valid form data successfully', async () => {
		mockFetch('gotr', true);

		const { component } = render(InputForm);
		const calculatedSpy = vi.fn();
		component.$on('calculated', calculatedSpy);

		const currentLevelInput = screen.getByLabelText(/current runecrafting level/i);
		const targetLevelInput = screen.getByLabelText(/target runecrafting level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '77');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '99');
		await user.click(calculateButton);

		await vi.waitFor(() => {
			expect(calculatedSpy).toHaveBeenCalledWith(
				expect.objectContaining({
					detail: {
						resultData: mockApiResponses.gotr
					}
				})
			);
		});
	});

	it('handles API errors correctly', async () => {
		mockFetch('gotr', false);

		const { component } = render(InputForm);
		const errorSpy = vi.fn();
		component.$on('error', errorSpy);

		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });
		await user.click(calculateButton);

		await vi.waitFor(() => {
			expect(errorSpy).toHaveBeenCalled();
		});
	});
});