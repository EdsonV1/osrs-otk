import { describe, it, expect } from 'vitest';
import { mockApiResponses, mockFetch } from './utils';

describe('Test Suite Setup', () => {
	it('should have correct mock API responses', () => {
		expect(mockApiResponses.wintertodt).toBeDefined();
		expect(mockApiResponses.gotr).toBeDefined();
		expect(mockApiResponses['ardy-knights']).toBeDefined();
		expect(mockApiResponses.birdhouses).toBeDefined();

		// Check wintertodt response structure
		expect(mockApiResponses.wintertodt.total_experience).toBe(500000);
		expect(mockApiResponses.wintertodt.pet_chance).toBe(12.5);
		expect(mockApiResponses.wintertodt.estimated_loot).toBeDefined();
	});

	it('should mock fetch correctly', () => {
		mockFetch('wintertodt', true);
		expect(fetch).toBeDefined();
	});

	it('should have all utilities available', () => {
		expect(typeof mockFetch).toBe('function');
		expect(typeof mockApiResponses).toBe('object');
	});
});

describe('API Mock Responses', () => {
	it('wintertodt mock has expected structure', () => {
		const response = mockApiResponses.wintertodt;
		expect(response).toMatchObject({
			total_experience: expect.any(Number),
			average_exp_hour: expect.any(Number),
			total_time: expect.any(Number),
			total_value: expect.any(Number),
			pet_chance: expect.any(Number),
			estimated_loot: expect.any(Object)
		});
	});

	it('gotr mock has expected structure', () => {
		const response = mockApiResponses.gotr;
		expect(response).toMatchObject({
			total_experience: expect.any(Number),
			average_exp_hour: expect.any(Number),
			games_needed: expect.any(Number),
			total_time: expect.any(Number),
			estimated_loot: expect.any(Object)
		});
	});

	it('ardy-knights mock has expected structure', () => {
		const response = mockApiResponses['ardy-knights'];
		expect(response).toMatchObject({
			total_experience: expect.any(Number),
			average_exp_hour: expect.any(Number),
			total_time: expect.any(Number),
			gp_per_hour: expect.any(Number),
			estimated_loot: expect.any(Object)
		});
	});

	it('birdhouses mock has expected structure', () => {
		const response = mockApiResponses.birdhouses;
		expect(response).toMatchObject({
			total_experience: expect.any(Number),
			average_exp_hour: expect.any(Number),
			runs_needed: expect.any(Number),
			total_time: expect.any(Number),
			estimated_loot: expect.any(Object)
		});
	});
});