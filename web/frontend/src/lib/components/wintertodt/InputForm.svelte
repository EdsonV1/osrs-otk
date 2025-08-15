<script lang="ts">
    import { createEventDispatcher, type EventDispatcher } from 'svelte';
    import type { WintertodtFormState, WintertodtApiInput, WintertodtApiResult, PlayerStats } from '$lib/types';
    import PlayerLookup from '../shared/PlayerLookup.svelte';

    interface ComponentEvents {
        calculated: { resultData: WintertodtApiResult };
        error: { message: string };
    }

    const dispatch: EventDispatcher<ComponentEvents> = createEventDispatcher();

    let formState: WintertodtFormState = {
        currentLevel: 50,
        targetLevel: 99,
        strategy: 'large_group',
        customPointsPerRound: undefined,
        customMinutesPerRound: undefined,
        skillLevels: {
            herblore: 1,
            mining: 1,
            fishing: 1,
            crafting: 1,
            farming: 1,
            woodcutting: 1
        }
    };

    let playerUsername = '';
    let useLivePrices = false;
    let playerLookupLoading = false;

    let useCustomStrategy = false;

    let isLoading: boolean = false;

    const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
    const fieldsetLegendClasses = "text-base font-semibold leading-7 text-theme-text-primary mb-1";
    const labelBaseClasses = "block text-xs font-medium text-theme-text-secondary mb-1";

    function handlePlayerStatsLoaded(event: CustomEvent<PlayerStats>) {
        const stats = event.detail;
        playerUsername = stats.username;
        
        // Update form state with player stats
        formState.currentLevel = stats.firemaking;
        formState.skillLevels = {
            herblore: stats.herblore,
            mining: stats.mining,
            fishing: stats.fishing,
            crafting: stats.crafting,
            farming: stats.farming,
            woodcutting: stats.woodcutting
        };

        // Handle level 99 case - if current level is already 99, keep target at 99
        if (stats.firemaking >= 99) {
            formState.targetLevel = 99;
            dispatch('error', { message: '' }); // Clear any errors
        } else {
            // Ensure target is higher than current, but cap at 99
            if (formState.targetLevel <= stats.firemaking) {
                formState.targetLevel = Math.min(stats.firemaking + 1, 99);
            }
        }
    }

    function handlePlayerLookupError(event: CustomEvent<string>) {
        dispatch('error', { message: `Player lookup failed: ${event.detail}` });
    }

    async function handleSubmit() {
        isLoading = true;
        dispatch('error', { message: '' });

        if (formState.currentLevel < 50) {
            dispatch('error', { message: 'Current Firemaking level must be at least 50' });
            isLoading = false;
            return;
        }
        
        if (formState.currentLevel >= 99 && formState.targetLevel >= 99) {
            // Allow calculation when both current and target are 99 (for loot simulation)
        } else if (formState.targetLevel < formState.currentLevel) {
            dispatch('error', { message: 'Target level must be greater than or equal to current level' });
            isLoading = false;
            return;
        }

        try {
            const apiInput = {
                current_level: formState.currentLevel,
                target_level: formState.targetLevel,
                strategy: formState.strategy,
                custom_points_per_round: useCustomStrategy ? formState.customPointsPerRound : undefined,
                custom_minutes_per_round: useCustomStrategy ? formState.customMinutesPerRound : undefined,
                skill_levels: formState.skillLevels,
                use_live_prices: useLivePrices,
                username: playerUsername.trim() || undefined
            };

            // Use live endpoint if live prices are enabled, otherwise use legacy
            const endpoint = useLivePrices ? 'http://localhost:8080/api/wintertodt/live' : 'http://localhost:8080/api/wintertodt';
            
            const response = await fetch(endpoint, {
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
    <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Wintertodt Calculator</h2>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Firemaking Levels</legend>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
                <label for="current-level" class={labelBaseClasses}>Current Level (minimum 50):</label>
                <input 
                    type="number" 
                    id="current-level" 
                    bind:value={formState.currentLevel} 
                    min="50" 
                    max="99" 
                    step="1" 
                    placeholder="e.g., 75" 
                    class="{inputBaseClasses} mt-1"
                    required
                />
            </div>
            <div>
                <label for="target-level" class={labelBaseClasses}>Target Level:</label>
                <input 
                    type="number" 
                    id="target-level" 
                    bind:value={formState.targetLevel} 
                    min="50" 
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
        <legend class={fieldsetLegendClasses}>Player Lookup (Optional)</legend>
        <div class="space-y-3">
            <div>
                <PlayerLookup 
                    bind:username={playerUsername}
                    bind:loading={playerLookupLoading}
                    placeholder="Enter OSRS username to auto-fill stats"
                    buttonText="Load Stats"
                    on:statsLoaded={handlePlayerStatsLoaded}
                    on:error={handlePlayerLookupError}
                />
                <p class="text-theme-text-tertiary text-xs mt-1">
                    Automatically loads your current skill levels from OSRS hiscores
                </p>
            </div>
            
            <div class="flex items-center">
                <input 
                    type="checkbox" 
                    id="use-live-prices" 
                    bind:checked={useLivePrices}
                    class="h-4 w-4 rounded border-theme-border bg-gray-700/50 text-theme-accent focus:ring-theme-accent focus:ring-2 focus:ring-offset-0"
                />
                <label for="use-live-prices" class="ml-2 text-sm text-theme-text-secondary">
                    Use live Grand Exchange prices (updates daily)
                </label>
            </div>
        </div>
    </fieldset>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Strategy</legend>
        <div>
            <label for="strategy" class={labelBaseClasses}>Strategy:</label>
            <select 
                id="strategy" 
                bind:value={formState.strategy} 
                class="{inputBaseClasses} mt-1"
                required
            >
                <option value="large_group">Large Group (4 min/round, 750 points)</option>
                <option value="solo">Solo (15 min/round, 1000 points)</option>
                <option value="efficient">Efficient Team (3.5 min/round, 600 points)</option>
            </select>
        </div>
        
        <div class="flex items-center space-x-2">
            <input 
                type="checkbox" 
                id="use-custom" 
                bind:checked={useCustomStrategy} 
                class="w-4 h-4 text-theme-accent bg-gray-700 border-gray-600 rounded focus:ring-theme-accent"
            />
            <label for="use-custom" class="text-sm text-theme-text-secondary">Use custom points/time per round</label>
        </div>

        {#if useCustomStrategy}
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 p-4 bg-theme-bg-primary/30 rounded-lg">
                <div>
                    <label for="custom-points" class={labelBaseClasses}>Custom Points per Round:</label>
                    <input 
                        type="number" 
                        id="custom-points" 
                        bind:value={formState.customPointsPerRound} 
                        min="1" 
                        max="2000" 
                        step="1" 
                        placeholder="e.g., 750" 
                        class="{inputBaseClasses} mt-1"
                    />
                </div>
                <div>
                    <label for="custom-minutes" class={labelBaseClasses}>Custom Minutes per Round:</label>
                    <input 
                        type="number" 
                        id="custom-minutes" 
                        bind:value={formState.customMinutesPerRound} 
                        min="0.1" 
                        max="60" 
                        step="0.1" 
                        placeholder="e.g., 4.0" 
                        class="{inputBaseClasses} mt-1"
                    />
                </div>
            </div>
        {/if}
    </fieldset>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Other Skill Levels (for drops)</legend>
        <p class="text-xs text-theme-text-tertiary mb-4">Higher skill levels improve loot quality and quantity from supply crates</p>
        <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
            <div>
                <label for="herblore" class={labelBaseClasses}>Herblore:</label>
                <input 
                    type="number" 
                    id="herblore" 
                    bind:value={formState.skillLevels.herblore} 
                    min="1" 
                    max="99" 
                    step="1" 
                    class="{inputBaseClasses} mt-1"
                />
            </div>
            <div>
                <label for="mining" class={labelBaseClasses}>Mining:</label>
                <input 
                    type="number" 
                    id="mining" 
                    bind:value={formState.skillLevels.mining} 
                    min="1" 
                    max="99" 
                    step="1" 
                    class="{inputBaseClasses} mt-1"
                />
            </div>
            <div>
                <label for="fishing" class={labelBaseClasses}>Fishing:</label>
                <input 
                    type="number" 
                    id="fishing" 
                    bind:value={formState.skillLevels.fishing} 
                    min="1" 
                    max="99" 
                    step="1" 
                    class="{inputBaseClasses} mt-1"
                />
            </div>
            <div>
                <label for="crafting" class={labelBaseClasses}>Crafting:</label>
                <input 
                    type="number" 
                    id="crafting" 
                    bind:value={formState.skillLevels.crafting} 
                    min="1" 
                    max="99" 
                    step="1" 
                    class="{inputBaseClasses} mt-1"
                />
            </div>
            <div>
                <label for="farming" class={labelBaseClasses}>Farming:</label>
                <input 
                    type="number" 
                    id="farming" 
                    bind:value={formState.skillLevels.farming} 
                    min="1" 
                    max="99" 
                    step="1" 
                    class="{inputBaseClasses} mt-1"
                />
            </div>
            <div>
                <label for="woodcutting" class={labelBaseClasses}>Woodcutting:</label>
                <input 
                    type="number" 
                    id="woodcutting" 
                    bind:value={formState.skillLevels.woodcutting} 
                    min="1" 
                    max="99" 
                    step="1" 
                    class="{inputBaseClasses} mt-1"
                />
            </div>
        </div>
    </fieldset>
    
    <button type="submit" disabled={isLoading}
            class="w-full flex justify-center py-3 px-4 border border-transparent rounded-lg shadow-button text-sm font-semibold text-white bg-theme-accent hover:bg-theme-accent-hover focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-theme-card-bg focus:ring-theme-accent disabled:opacity-60 disabled:cursor-not-allowed transition-colors duration-150">
        {isLoading ? 'Calculating...' : 'Calculate Wintertodt'}
    </button>
</form>