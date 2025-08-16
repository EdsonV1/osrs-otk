<script lang="ts">
    import type { GOTRFormState, GOTRApiInput, GOTRApiResult, PlayerStats } from '$lib/types';
    import { createEventDispatcher } from 'svelte';
    import PlayerLookup from '../shared/PlayerLookup.svelte';

    const dispatch = createEventDispatcher<{
        calculated: { resultData: GOTRApiResult };
        error: { message: string };
    }>();

    let formData: GOTRFormState = {
        currentLevel: 27,
        targetLevel: 77
    };

    let playerUsername = '';
    let playerLookupLoading = false;
    let isLoading = false;
    let validationErrors: { [key: string]: string } = {};

    function handlePlayerStatsLoaded(event: CustomEvent<PlayerStats>) {
        const stats = event.detail;
        playerUsername = stats.username;
        
        // Update form with player's runecrafting level
        formData.currentLevel = stats.runecrafting;
        
        // Handle level 99 case
        if (stats.runecrafting >= 99) {
            formData.targetLevel = 99;
        } else {
            // Ensure target is higher than current, but cap at 99
            if (formData.targetLevel <= stats.runecrafting) {
                formData.targetLevel = Math.min(stats.runecrafting + 1, 99);
            }
        }
        
        validationErrors = {}; // Clear any validation errors
    }

    function handlePlayerLookupError(event: CustomEvent<string>) {
        dispatch('error', { message: `Player lookup failed: ${event.detail}` });
    }

    function validateForm(): boolean {
        validationErrors = {};
        
        if (formData.currentLevel < 27 || formData.currentLevel > 126) {
            validationErrors.currentLevel = 'Current level must be between 27 and 126 (minimum to access GOTR)';
        }
        
        if (formData.targetLevel < 27 || formData.targetLevel > 126) {
            validationErrors.targetLevel = 'Target level must be between 27 and 126';
        }
        
        if (formData.currentLevel >= 99 && formData.targetLevel >= 99) {
            // Allow calculation when both current and target are 99 (for reward calculations)
        } else if (formData.targetLevel <= formData.currentLevel) {
            validationErrors.targetLevel = 'Target level must be higher than current level';
        }

        return Object.keys(validationErrors).length === 0;
    }

    async function handleSubmit() {
        if (!validateForm()) return;

        isLoading = true;
        
        try {
            const apiInput: GOTRApiInput = {
                current_level: formData.currentLevel,
                target_level: formData.targetLevel
            };

            const response = await fetch('http://localhost:8080/api/tools/gotr', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(apiInput)
            });

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.error || 'Failed to calculate GOTR results');
            }

            dispatch('calculated', { resultData: data as GOTRApiResult });
        } catch (error) {
            console.error('GOTR calculation error:', error);
            dispatch('error', { 
                message: error instanceof Error ? error.message : 'An unexpected error occurred' 
            });
        } finally {
            isLoading = false;
        }
    }

    // Level presets
    const levelPresets = [
        { label: "Level 27 → 77 (Early GOTR)", current: 27, target: 77 },
        { label: "Level 77 → 99 (Optimal GOTR)", current: 77, target: 99 },
        { label: "Level 50 → 99", current: 50, target: 99 },
        { label: "Level 27 → 99 (Full GOTR)", current: 27, target: 99 }
    ];

    function applyPreset(preset: typeof levelPresets[0]) {
        formData.currentLevel = preset.current;
        formData.targetLevel = preset.target;
        validationErrors = {};
    }
</script>

<form on:submit|preventDefault={handleSubmit} class="space-y-6">
    <!-- Player Lookup -->
    <div class="space-y-4">
        <div class="block text-sm font-semibold text-theme-text-primary mb-3">Player Lookup (Optional)</div>
        <div>
            <PlayerLookup 
                bind:username={playerUsername}
                bind:loading={playerLookupLoading}
                placeholder="Enter OSRS username to auto-fill current level"
                buttonText="Load Stats"
                on:statsLoaded={handlePlayerStatsLoaded}
                on:error={handlePlayerLookupError}
            />
            <p class="text-theme-text-tertiary text-xs mt-1">
                Automatically loads your current runecrafting level from OSRS hiscores
            </p>
        </div>
    </div>

    <!-- Level Presets -->
    <div class="mb-6">
        <div class="block text-sm font-semibold text-theme-text-primary mb-3">Quick Presets</div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-3">
            {#each levelPresets as preset}
                <button
                    type="button"
                    on:click={() => applyPreset(preset)}
                    class="text-left p-3 bg-theme-bg-tertiary hover:bg-theme-bg-elevated border border-theme-border-subtle hover:border-theme-accent-primary/30 rounded-lg transition-colors duration-200 text-sm"
                >
                    <div class="font-medium text-theme-text-primary">{preset.label}</div>
                    <div class="text-xs text-theme-text-tertiary mt-1">
                        {preset.current} → {preset.target}
                    </div>
                </button>
            {/each}
        </div>
    </div>

    <!-- Level Inputs -->
    <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
        <!-- Current Level -->
        <div class="space-y-2">
            <label for="currentLevel" class="block text-sm font-semibold text-theme-text-primary">
                Current Runecrafting Level
            </label>
            <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <div class="w-5 h-5 bg-gradient-to-br from-purple-500 to-blue-600 rounded flex items-center justify-center">
                        <span class="text-white text-xs font-bold">RC</span>
                    </div>
                </div>
                <input
                    type="number"
                    id="currentLevel"
                    bind:value={formData.currentLevel}
                    min="27"
                    max="126"
                    class="w-full pl-12 pr-4 py-3 bg-theme-bg-elevated border border-theme-border-primary rounded-lg focus:ring-2 focus:ring-theme-accent-primary focus:border-theme-accent-primary transition-colors text-theme-text-primary placeholder-theme-text-tertiary"
                    placeholder="Enter current level"
                    class:border-red-500={validationErrors.currentLevel}
                />
            </div>
            {#if validationErrors.currentLevel}
                <p class="text-red-400 text-xs">{validationErrors.currentLevel}</p>
            {/if}
        </div>

        <!-- Target Level -->
        <div class="space-y-2">
            <label for="targetLevel" class="block text-sm font-semibold text-theme-text-primary">
                Target Runecrafting Level
            </label>
            <div class="relative">
                <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <div class="w-5 h-5 bg-gradient-to-br from-green-500 to-emerald-600 rounded flex items-center justify-center">
                        <span class="text-white text-xs font-bold">🎯</span>
                    </div>
                </div>
                <input
                    type="number"
                    id="targetLevel"
                    bind:value={formData.targetLevel}
                    min="27"
                    max="126"
                    class="w-full pl-12 pr-4 py-3 bg-theme-bg-elevated border border-theme-border-primary rounded-lg focus:ring-2 focus:ring-theme-accent-primary focus:border-theme-accent-primary transition-colors text-theme-text-primary placeholder-theme-text-tertiary"
                    placeholder="Enter target level"
                    class:border-red-500={validationErrors.targetLevel}
                />
            </div>
            {#if validationErrors.targetLevel}
                <p class="text-red-400 text-xs">{validationErrors.targetLevel}</p>
            {/if}
        </div>
    </div>

    <!-- Info Card -->
    <div class="bg-gradient-to-r from-purple-500/10 via-purple-600/5 to-transparent border border-purple-500/20 rounded-lg p-4">
        <div class="flex items-start space-x-3">
            <div class="w-6 h-6 bg-purple-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                <span class="text-white text-xs">ℹ️</span>
            </div>
            <div>
                <h4 class="text-sm font-semibold text-purple-400 mb-1">About GOTR</h4>
                <p class="text-xs text-theme-text-secondary leading-relaxed">
                    Guardians of the Rift provides excellent Runecrafting XP rates and valuable rewards. 
                    This calculator estimates your progress, pet chances, and potential loot based on average game performance.
                </p>
            </div>
        </div>
    </div>

    <!-- Calculate Button -->
    <button
        type="submit"
        disabled={isLoading}
        class="w-full bg-theme-accent hover:bg-theme-accent-hover disabled:bg-theme-bg-tertiary disabled:cursor-not-allowed text-white font-semibold py-4 px-6 rounded-button shadow-button hover:shadow-button-hover transition-all duration-200 flex items-center justify-center"
    >
        {#if isLoading}
            <div class="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin mr-3"></div>
            Calculating...
        {:else}
            <div class="flex items-center">
                <span class="mr-2">🧙‍♂️</span>
                Calculate GOTR Training
            </div>
        {/if}
    </button>
</form>