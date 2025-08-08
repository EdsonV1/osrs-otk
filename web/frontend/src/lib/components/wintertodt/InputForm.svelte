<script lang="ts">
    import { createEventDispatcher, type EventDispatcher } from 'svelte';
    import type { WintertodtFormState, WintertodtApiInput, WintertodtApiResult } from '$lib/types';

    interface ComponentEvents {
        calculated: { resultData: WintertodtApiResult };
        error: { message: string };
    }

    const dispatch: EventDispatcher<ComponentEvents> = createEventDispatcher();

    let formState: WintertodtFormState = {
        firemakingLevel: 50,
        roundsPerHour: 4,
        totalRounds: 100
    };

    let isLoading: boolean = false;

    const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
    const fieldsetLegendClasses = "text-base font-semibold leading-7 text-theme-text-primary mb-1";
    const labelBaseClasses = "block text-xs font-medium text-theme-text-secondary mb-1";

    async function handleSubmit() {
        isLoading = true;
        dispatch('error', { message: '' });

        if (formState.firemakingLevel < 50) {
            dispatch('error', { message: 'Firemaking level must be at least 50' });
            isLoading = false;
            return;
        }
        
        if (formState.roundsPerHour <= 0 || formState.totalRounds <= 0) {
            dispatch('error', { message: 'Rounds per hour and total rounds must be positive' });
            isLoading = false;
            return;
        }

        try {
            const apiInput: WintertodtApiInput = {
                firemaking_level: formState.firemakingLevel,
                rounds_per_hour: formState.roundsPerHour,
                total_rounds: formState.totalRounds
            };

            const response = await fetch('http://localhost:8080/api/wintertodt', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(apiInput),
            });

            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(errorText || 'Calculation failed');
            }

            const resultData = await response.json() as WintertodtApiResult;
            dispatch('calculated', { resultData });
        } catch (err: any) {
            dispatch('error', { message: err.message || 'Calculation failed.' });
        } finally {
            isLoading = false;
        }
    }
</script>

<form class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8" on:submit|preventDefault={handleSubmit}>
    <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Wintertodt Inputs</h2>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Character Stats</legend>
        <div>
            <label for="firemaking-level" class={labelBaseClasses}>Firemaking Level (minimum 50):</label>
            <input 
                type="number" 
                id="firemaking-level" 
                bind:value={formState.firemakingLevel} 
                min="50" 
                max="99" 
                step="1" 
                placeholder="e.g., 75" 
                class="{inputBaseClasses} mt-1"
                required
            />
        </div>
    </fieldset>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Calculation Details</legend>
        <div>
            <label for="rounds-per-hour" class={labelBaseClasses}>Rounds Per Hour:</label>
            <input 
                type="number" 
                id="rounds-per-hour" 
                bind:value={formState.roundsPerHour} 
                min="0.1" 
                max="20" 
                step="0.1" 
                placeholder="e.g., 4.5" 
                class="{inputBaseClasses} mt-1"
                required
            />
            <p class="mt-1 text-xs text-theme-text-tertiary">Typical range: 3-6 rounds per hour</p>
        </div>
        <div>
            <label for="total-rounds" class={labelBaseClasses}>Total Number of Rounds:</label>
            <input 
                type="number" 
                id="total-rounds" 
                bind:value={formState.totalRounds} 
                min="1" 
                max="10000" 
                step="1" 
                placeholder="e.g., 100" 
                class="{inputBaseClasses} mt-1"
                required
            />
        </div>
    </fieldset>
    
    <button type="submit" disabled={isLoading}
            class="w-full flex justify-center py-3 px-4 border border-transparent rounded-lg shadow-button text-sm font-semibold text-white bg-theme-accent hover:bg-theme-accent-hover focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-theme-card-bg focus:ring-theme-accent disabled:opacity-60 disabled:cursor-not-allowed transition-colors duration-150">
        {isLoading ? 'Calculating...' : 'Calculate Wintertodt'}
    </button>
</form>