import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import { mockFetch, mockApiResponses } from '../../../tests/utils';
import InputForm from './InputForm.svelte';

describe('Birdhouses InputForm', () => {
	const user = userEvent.setup();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('renders form elements correctly', () => {
		render(InputForm);

		expect(screen.getByLabelText(/current hunter level/i)).toBeInTheDocument();
		expect(screen.getByLabelText(/target hunter level/i)).toBeInTheDocument();
		expect(screen.getByRole('button', { name: /calculate training plan/i })).toBeInTheDocument();
	});

	it('shows validation error for level below minimum', async () => {
		render(InputForm);

		const currentLevelInput = screen.getByLabelText(/current hunter level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '4'); // Below minimum level
		await user.click(calculateButton);

		expect(screen.getByText(/current level must be at least 5/i)).toBeInTheDocument();
	});

	it('shows validation error for target level lower than current', async () => {
		render(InputForm);

		const currentLevelInput = screen.getByLabelText(/current hunter level/i);
		const targetLevelInput = screen.getByLabelText(/target hunter level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '50');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '40');
		await user.click(calculateButton);

		expect(screen.getByText(/target level must be higher than current level/i)).toBeInTheDocument();
	});

	it('submits valid form data successfully', async () => {
		mockFetch('birdhouses', true);

		const { component } = render(InputForm);
		const calculatedSpy = vi.fn();
		component.$on('calculated', calculatedSpy);

		const currentLevelInput = screen.getByLabelText(/current hunter level/i);
		const targetLevelInput = screen.getByLabelText(/target hunter level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '60');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '80');
		await user.click(calculateButton);

		await vi.waitFor(() => {
			expect(calculatedSpy).toHaveBeenCalledWith(
				expect.objectContaining({
					detail: {
						resultData: mockApiResponses.birdhouses
					}
				})
			);
		});
	});

	it('handles API errors correctly', async () => {
		mockFetch('birdhouses', false);

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