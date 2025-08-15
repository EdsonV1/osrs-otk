<script lang="ts">
    import { createEventDispatcher, type EventDispatcher } from 'svelte';
    import type { ArdyKnightInput, ArdyKnightResult, PlayerStats } from '$lib/types';
    import PlayerLookup from '../shared/PlayerLookup.svelte';

    interface ComponentEvents {
        calculated: { data: ArdyKnightResult };
        error: { error: string };
    }

    const dispatch: EventDispatcher<ComponentEvents> = createEventDispatcher();

    let formState: ArdyKnightInput = {
        current_thieving_level: 55,
        target_thieving_level: 99,
        has_ardy_med: false,
        has_thieving_cape: false,
        has_rogues_outfit: false,
        has_shadow_veil: false,
        hourly_pickpockets: 1000,
        food_heal_amount: 20,
        food_cost: 100
    };

    let playerUsername = '';
    let playerLookupLoading = false;
    let isLoading: boolean = false;

    const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
    const fieldsetLegendClasses = "text-base font-semibold leading-7 text-theme-text-primary mb-1";
    const labelBaseClasses = "block text-xs font-medium text-theme-text-secondary mb-1";
    const checkboxClasses = "h-4 w-4 text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg border-theme-border-input rounded";

    function handlePlayerStatsLoaded(event: CustomEvent<PlayerStats>) {
        const stats = event.detail;
        playerUsername = stats.username;
        
        // Update form state with player thieving level
        formState.current_thieving_level = stats.thieving;

        // Handle level 99 case
        if (stats.thieving >= 99) {
            formState.target_thieving_level = 99;
        } else {
            // Ensure target is higher than current, but cap at 99
            if (formState.target_thieving_level <= stats.thieving) {
                formState.target_thieving_level = Math.min(stats.thieving + 1, 99);
            }
        }
    }

    function handlePlayerLookupError(event: CustomEvent<string>) {
        dispatch('error', { error: `Player lookup failed: ${event.detail}` });
    }

    async function handleSubmit() {
        isLoading = true;
        dispatch('error', { error: '' });

        if (!formState.current_thieving_level || !formState.target_thieving_level) {
            dispatch('error', { error: 'Current and target thieving levels are required' });
            isLoading = false;
            return;
        }

        if (formState.current_thieving_level >= 99 && formState.target_thieving_level >= 99) {
            // Allow calculation when both current and target are 99 (for profit calculations)
        } else if (formState.current_thieving_level >= formState.target_thieving_level) {
            dispatch('error', { error: 'Target level must be higher than current level' });
            isLoading = false;
            return;
        }

        try {
            const response = await fetch('http://localhost:8080/api/ardyknights', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(formState)
            });

            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(errorText || `API Error: ${response.status}`);
            }
            
            const resultData = await response.json() as ArdyKnightResult;
            dispatch('calculated', { data: resultData });
        } catch (err: any) {
            dispatch('error', { error: err.message || 'Calculation failed.' });
        } finally {
            isLoading = false;
        }
    }
</script>

<form class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8" on:submit|preventDefault={handleSubmit}>
    <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Ardougne Knight Inputs</h2>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Player Lookup (Optional)</legend>
        <div>
            <PlayerLookup 
                bind:username={playerUsername}
                bind:loading={playerLookupLoading}
                placeholder="Enter OSRS username to auto-fill thieving level"
                buttonText="Load Stats"
                on:statsLoaded={handlePlayerStatsLoaded}
                on:error={handlePlayerLookupError}
            />
            <p class="text-theme-text-tertiary text-xs mt-1">
                Automatically loads your current thieving level from OSRS hiscores
            </p>
        </div>
    </fieldset>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Character Stats</legend>
        <div class="grid grid-cols-2 gap-4">
            <div>
                <label for="current-level" class={labelBaseClasses}>Current Thieving Level:</label>
                <input 
                    type="number" 
                    id="current-level" 
                    bind:value={formState.current_thieving_level} 
                    min="1" 
                    max="98" 
                    step="1" 
                    placeholder="e.g., 55" 
                    class="{inputBaseClasses} mt-1"
                    required
                />
            </div>
            <div>
                <label for="target-level" class={labelBaseClasses}>Target Thieving Level:</label>
                <input 
                    type="number" 
                    id="target-level" 
                    bind:value={formState.target_thieving_level} 
                    min="2" 
                    max="99" 
                    step="1" 
                    placeholder="e.g., 99" 
                    class="{inputBaseClasses} mt-1"
                    required
                />
            </div>
        </div>
    </fieldset>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Equipment & Bonuses</legend>
        <div class="grid grid-cols-2 gap-4">
            <label class="flex items-center space-x-3">
                <input type="checkbox" bind:checked={formState.has_ardy_med} class={checkboxClasses} />
                <span class="text-sm text-theme-text-primary">Ardougne Medium Diary</span>
            </label>
            <label class="flex items-center space-x-3">
                <input type="checkbox" bind:checked={formState.has_thieving_cape} class={checkboxClasses} />
                <span class="text-sm text-theme-text-primary">Thieving Cape</span>
            </label>
            <label class="flex items-center space-x-3">
                <input type="checkbox" bind:checked={formState.has_rogues_outfit} class={checkboxClasses} />
                <span class="text-sm text-theme-text-primary">Full Rogue's Outfit</span>
            </label>
            <label class="flex items-center space-x-3">
                <input type="checkbox" bind:checked={formState.has_shadow_veil} class={checkboxClasses} />
                <span class="text-sm text-theme-text-primary">Shadow Veil Spell</span>
            </label>
        </div>
    </fieldset>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Training Settings</legend>
        <div class="space-y-4">
            <div>
                <label for="hourly-pickpockets" class={labelBaseClasses}>Hourly Pickpockets:</label>
                <input 
                    type="number" 
                    id="hourly-pickpockets" 
                    bind:value={formState.hourly_pickpockets} 
                    min="100" 
                    max="5000" 
                    step="50" 
                    placeholder="e.g., 1000" 
                    class="{inputBaseClasses} mt-1"
                    required
                />
            </div>
            <div class="grid grid-cols-2 gap-4">
                <div>
                    <label for="food-heal" class={labelBaseClasses}>Food Heal Amount:</label>
                    <input 
                        type="number" 
                        id="food-heal" 
                        bind:value={formState.food_heal_amount} 
                        min="1" 
                        max="99" 
                        step="1" 
                        placeholder="e.g., 20" 
                        class="{inputBaseClasses} mt-1"
                        required
                    />
                </div>
                <div>
                    <label for="food-cost" class={labelBaseClasses}>Food Cost (GP):</label>
                    <input 
                        type="number" 
                        id="food-cost" 
                        bind:value={formState.food_cost} 
                        min="1" 
                        max="10000" 
                        step="1" 
                        placeholder="e.g., 100" 
                        class="{inputBaseClasses} mt-1"
                        required
                    />
                </div>
            </div>
        </div>
    </fieldset>
    
    <button type="submit" disabled={isLoading}
            class="w-full flex justify-center py-3 px-4 border border-transparent rounded-lg shadow-button text-sm font-semibold text-white bg-theme-accent hover:bg-theme-accent-hover focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-theme-card-bg focus:ring-theme-accent disabled:opacity-60 disabled:cursor-not-allowed transition-colors duration-150">
        {isLoading ? 'Calculating...' : 'Calculate Ardougne Knights'}
    </button>
</form>