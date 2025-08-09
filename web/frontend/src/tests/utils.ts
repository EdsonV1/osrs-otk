import { vi } from 'vitest';

// Re-export the render function directly
export { render } from '@testing-library/svelte';

// Mock API responses for different tools
export const mockApiResponses = {
	wintertodt: {
		total_experience: 500000,
		average_exp_hour: 180000,
		total_time: 2.78,
		total_value: 2500000,
		pet_chance: 12.5,
		estimated_loot: {
			'burnt_page': 100,
			'supply_crate': 100,
			'magic_logs': 250,
			'yew_logs': 180
		}
	},
	gotr: {
		total_experience: 2800000,
		average_exp_hour: 45000,
		games_needed: 175,
		total_time: 62.2,
		estimated_loot: {
			'catalytic_talisman': 8,
			'elemental_talisman': 12,
			'lantern': 3
		}
	},
	'ardy-knights': {
		total_experience: 1200000,
		average_exp_hour: 240000,
		total_time: 5.0,
		gp_per_hour: 180000,
		estimated_loot: {
			'coin_pouch': 1500,
			'gems': 45
		}
	},
	birdhouses: {
		total_experience: 600000,
		average_exp_hour: 150000,
		runs_needed: 200,
		total_time: 4.0,
		estimated_loot: {
			'bird_nest': 50,
			'raw_bird_meat': 800
		}
	}
};

// Mock fetch with different responses
export function mockFetch(tool: keyof typeof mockApiResponses, success: boolean = true) {
	return vi.mocked(fetch).mockResolvedValueOnce({
		ok: success,
		json: async () => success ? mockApiResponses[tool] : { error: 'API Error' }
	} as Response);
}

// Create test IDs for easier element selection
export const testIds = {
	// Forms
	currentLevelInput: 'current-level-input',
	targetLevelInput: 'target-level-input',
	calculateButton: 'calculate-button',
	
	// Results
	resultsContainer: 'results-container',
	experienceValue: 'experience-value',
	timeValue: 'time-value',
	lootValue: 'loot-value',
	petChance: 'pet-chance',
	
	// Error states
	errorMessage: 'error-message',
	
	// Loading states
	loadingSpinner: 'loading-spinner'
};

// Wait for async operations in tests
export function waitForApiCall() {
	return new Promise(resolve => setTimeout(resolve, 0));
}