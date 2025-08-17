<script lang="ts">
    import { createEventDispatcher, type EventDispatcher } from 'svelte';
    import type { HerbiboarApiInput, HerbiboarApiResult, HerbiboarFormState, PlayerStats } from '$lib/types';
    import PlayerLookup from '$lib/components/shared/PlayerLookup.svelte';

    interface ComponentEvents {
        calculated: { resultData: HerbiboarApiResult };
        error: { message: string };
    }

    const dispatch: EventDispatcher<ComponentEvents> = createEventDispatcher();

    let formState: HerbiboarFormState = {
        hunterLevel: 80,
        herbloreLevel: 31,
        magicSecateurs: true,
        calculationType: 'number',
        targetLevel: null,
        numberToCatch: 100,
        useLivePrices: true,
        username: '',
    };

    let isLoading: boolean = false;

    const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
    const fieldsetLegendClasses = "text-base font-semibold leading-7 text-theme-text-primary mb-1";
    const labelBaseClasses = "block text-xs font-medium text-theme-text-secondary mb-1";
    const radioLabelClasses = "flex items-center space-x-3 p-3 border border-theme-border-subtle rounded-lg hover:border-theme-accent/70 cursor-pointer transition-all duration-150 ease-in-out";
    const radioInputClasses = "h-4 w-4 text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg border-theme-border-input";

    async function handleSubmit() {
        isLoading = true;
        dispatch('error', { message: '' });

        // Validation
        if (formState.hunterLevel < 80) {
            dispatch('error', { message: 'Hunter level must be at least 80.' });
            isLoading = false;
            return;
        }

        if (formState.herbloreLevel < 31) {
            dispatch('error', { message: 'Herblore level must be at least 31.' });
            isLoading = false;
            return;
        }

        if (formState.calculationType === 'target') {
            if (!formState.targetLevel || formState.targetLevel <= formState.hunterLevel) {
                dispatch('error', { message: 'Target level must be higher than current level.' });
                isLoading = false;
                return;
            }
        } else {
            if (!formState.numberToCatch || formState.numberToCatch <= 0) {
                dispatch('error', { message: 'Number to catch must be greater than 0.' });
                isLoading = false;
                return;
            }
        }

        const apiInput: HerbiboarApiInput = {
            hunter_level: formState.hunterLevel,
            herblore_level: formState.herbloreLevel,
            magic_secateurs: formState.magicSecateurs,
            calculation_type: formState.calculationType,
            use_live_prices: formState.useLivePrices,
            username: formState.username || undefined,
        };

        if (formState.calculationType === 'target') {
            apiInput.target_level = formState.targetLevel!;
        } else {
            apiInput.number_to_catch = formState.numberToCatch!;
        }

        try {
            const endpoint = formState.useLivePrices ? '/api/herbiboar/live' : '/api/herbiboar';
            const response = await fetch(`http://localhost:8080${endpoint}`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(apiInput),
            });

            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(errorText || `HTTP ${response.status}`);
            }

            const result: HerbiboarApiResult = await response.json();
            dispatch('calculated', { resultData: result });

        } catch (error) {
            console.error('Herbiboar calculation failed:', error);
            dispatch('error', { 
                message: error instanceof Error ? error.message : 'Unknown error occurred' 
            });
        } finally {
            isLoading = false;
        }
    }

    function handlePlayerLookup(event: CustomEvent<PlayerStats>) {
        const playerStats = event.detail;
        formState.hunterLevel = Math.max(formState.hunterLevel, playerStats.hunter);
        formState.herbloreLevel = Math.max(formState.herbloreLevel, playerStats.herblore);
    }
</script>

<form on:submit|preventDefault={handleSubmit} class="space-y-6">
    <!-- Player Lookup Section -->
    <fieldset>
        <legend class={fieldsetLegendClasses}>Player Stats (Optional)</legend>
        <PlayerLookup 
            bind:username={formState.username} 
            on:statsLoaded={handlePlayerLookup}
        />
    </fieldset>

    <!-- Skill Levels -->
    <fieldset>
        <legend class={fieldsetLegendClasses}>Skill Levels</legend>
        <div class="grid grid-cols-2 gap-4">
            <div>
                <label for="hunter-level" class={labelBaseClasses}>
                    Hunter Level (min. 80)
                </label>
                <input
                    id="hunter-level"
                    type="number"
                    min="80"
                    max="120"
                    bind:value={formState.hunterLevel}
                    class={inputBaseClasses}
                    required
                />
            </div>
            <div>
                <label for="herblore-level" class={labelBaseClasses}>
                    Herblore Level (min. 31)
                </label>
                <input
                    id="herblore-level"
                    type="number"
                    min="31"
                    max="120"
                    bind:value={formState.herbloreLevel}
                    class={inputBaseClasses}
                    required
                />
            </div>
        </div>
    </fieldset>

    <!-- Gear Options -->
    <fieldset>
        <legend class={fieldsetLegendClasses}>Gear & Equipment</legend>
        <div class="space-y-3">
            <label class="flex items-center space-x-3">
                <input
                    type="checkbox"
                    bind:checked={formState.magicSecateurs}
                    class="h-4 w-4 text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg border-theme-border-input rounded"
                />
                <span class="text-sm text-theme-text-primary">
                    Magic Secateurs (+1 herb per herbiboar)
                </span>
            </label>
        </div>
    </fieldset>

    <!-- Calculation Type -->
    <fieldset>
        <legend class={fieldsetLegendClasses}>Calculation Mode</legend>
        <div class="space-y-3">
            <label class={radioLabelClasses}>
                <input
                    type="radio"
                    name="calculation-type"
                    value="number"
                    bind:group={formState.calculationType}
                    class={radioInputClasses}
                />
                <div>
                    <div class="text-sm font-medium text-theme-text-primary">Number Mode</div>
                    <div class="text-xs text-theme-text-secondary">Calculate for a specific number of herbiboars</div>
                </div>
            </label>
            <label class={radioLabelClasses}>
                <input
                    type="radio"
                    name="calculation-type"
                    value="target"
                    bind:group={formState.calculationType}
                    class={radioInputClasses}
                />
                <div>
                    <div class="text-sm font-medium text-theme-text-primary">Target Mode</div>
                    <div class="text-xs text-theme-text-secondary">Calculate to reach a target Hunter level</div>
                </div>
            </label>
        </div>
    </fieldset>

    <!-- Target/Number Input -->
    {#if formState.calculationType === 'target'}
        <div>
            <label for="target-level" class={labelBaseClasses}>
                Target Hunter Level
            </label>
            <input
                id="target-level"
                type="number"
                min={formState.hunterLevel + 1}
                max="120"
                bind:value={formState.targetLevel}
                class={inputBaseClasses}
                required
            />
        </div>
    {:else}
        <div>
            <label for="number-to-catch" class={labelBaseClasses}>
                Number of Herbiboars to Catch
            </label>
            <input
                id="number-to-catch"
                type="number"
                min="1"
                max="10000"
                bind:value={formState.numberToCatch}
                class={inputBaseClasses}
                required
            />
        </div>
    {/if}

    <!-- Live Prices -->
    <fieldset>
        <legend class={fieldsetLegendClasses}>Pricing Options</legend>
        <div class="space-y-3">
            <label class="flex items-center space-x-3">
                <input
                    type="checkbox"
                    bind:checked={formState.useLivePrices}
                    class="h-4 w-4 text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg border-theme-border-input rounded"
                />
                <span class="text-sm text-theme-text-primary">
                    Use Live Prices (updated from OSRS Wiki)
                </span>
            </label>
        </div>
    </fieldset>

    <!-- Submit Button -->
    <button
        type="submit"
        disabled={isLoading}
        class="w-full py-3 px-6 bg-theme-accent hover:bg-theme-accent-hover text-white font-semibold rounded-lg transition-colors duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
    >
        {isLoading ? 'Calculating...' : 'Calculate Herbiboar Results'}
    </button>
</form>