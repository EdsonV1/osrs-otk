import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import userEvent from '@testing-library/user-event';
import { mockFetch, mockApiResponses } from '../../../tests/utils';
import InputForm from './InputForm.svelte';

describe('Ardy Knights InputForm', () => {
	const user = userEvent.setup();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it('renders form elements correctly', () => {
		render(InputForm);

		expect(screen.getByLabelText(/current thieving level/i)).toBeInTheDocument();
		expect(screen.getByLabelText(/target thieving level/i)).toBeInTheDocument();
		expect(screen.getByRole('button', { name: /calculate training plan/i })).toBeInTheDocument();
	});

	it('shows validation error for level below minimum', async () => {
		render(InputForm);

		const currentLevelInput = screen.getByLabelText(/current thieving level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '54'); // Below minimum level
		await user.click(calculateButton);

		expect(screen.getByText(/current level must be at least 55/i)).toBeInTheDocument();
	});

	it('shows validation error for target level lower than current', async () => {
		render(InputForm);

		const currentLevelInput = screen.getByLabelText(/current thieving level/i);
		const targetLevelInput = screen.getByLabelText(/target thieving level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '80');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '75');
		await user.click(calculateButton);

		expect(screen.getByText(/target level must be higher than current level/i)).toBeInTheDocument();
	});

	it('submits valid form data successfully', async () => {
		mockFetch('ardy-knights', true);

		const { component } = render(InputForm);
		const calculatedSpy = vi.fn();
		component.$on('calculated', calculatedSpy);

		const currentLevelInput = screen.getByLabelText(/current thieving level/i);
		const targetLevelInput = screen.getByLabelText(/target thieving level/i);
		const calculateButton = screen.getByRole('button', { name: /calculate training plan/i });

		await user.clear(currentLevelInput);
		await user.type(currentLevelInput, '75');
		await user.clear(targetLevelInput);
		await user.type(targetLevelInput, '99');
		await user.click(calculateButton);

		await vi.waitFor(() => {
			expect(calculatedSpy).toHaveBeenCalledWith(
				expect.objectContaining({
					detail: {
						resultData: mockApiResponses['ardy-knights']
					}
				})
			);
		});
	});

	it('handles API errors correctly', async () => {
		mockFetch('ardy-knights', false);

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