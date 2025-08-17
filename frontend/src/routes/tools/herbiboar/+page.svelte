<script lang="ts">
    import type { HerbiboarApiResult } from '$lib/types';
    import HerbiboarForm from '$lib/components/herbiboar/InputForm.svelte';
    import HerbiboarResults from '$lib/components/herbiboar/ResultsDisplay.svelte';
    import { tick } from 'svelte';

    interface ProTipsData {
        calculation_methodology: {
            xp_rates_source: string;
            base_formula: string;
            data_points: Array<{level: number, xp_per_catch: number, note: string}>;
        };
        game_mechanics: {
            catch_time: string;
            gear_effects: string;
            herb_yields: string[];
        };
        factors_considered: string[];
        accuracy_notes: {
            rates_vary: string;
            variance_factors: string[];
            calculation_basis: string;
        };
        pro_tips: Array<{tip: string, description: string}>;
        herb_calculation: {
            drop_rates: string;
            price_sources: string;
            profit_factors: string;
            gp_calculation: string;
        };
    }

    // Tool configuration
    const toolConfig = {
        name: 'Herbiboar Calculator',
        description: 'Calculate Hunter and Herblore XP, herb profits, and pet chances from herbiboar hunting.',
        iconSrc: '/herbiboar/herbiboar.png'
    };

    let currentApiResult: HerbiboarApiResult | null = null;
    let currentError: string | null = null;
    let resultsElement: HTMLElement;
    let proTipsData: ProTipsData | null = null;
    let showProTips = false;

    function scrollToResults() {
        if (resultsElement) {
            resultsElement.scrollIntoView({ 
                behavior: 'smooth', 
                block: 'start',
                inline: 'nearest'
            });
        }
    }

    async function handleCalculation(event: CustomEvent<{ resultData: HerbiboarApiResult }>) {
        currentApiResult = event.detail.resultData;
        currentError = null;
        
        // Wait for the DOM to update, then scroll to results
        await tick();
        setTimeout(() => scrollToResults(), 100);
    }

    async function handleCalcError(event: CustomEvent<{ message: string }>) {
        currentError = event.detail.message;
        currentApiResult = null;
        
        // Wait for the DOM to update, then scroll to error
        await tick();
        setTimeout(() => scrollToResults(), 100);
    }

    async function loadProTips() {
        if (proTipsData) {
            showProTips = !showProTips;
            return;
        }

        try {
            const response = await fetch('http://localhost:8080/api/tools/herbiboar/tips');
            if (!response.ok) {
                throw new Error('Failed to load pro tips');
            }
            proTipsData = await response.json();
            showProTips = true;
        } catch (error) {
            console.error('Error loading pro tips:', error);
            currentError = 'Failed to load calculation methodology';
        }
    }
</script>

<div class="max-w-6xl mx-auto">
    <div class="mb-6"> 
        <a href="/tools" class="inline-flex items-center text-sm font-medium text-theme-accent hover:text-theme-accent-hover transition-colors group">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5 mr-1.5 transform transition-transform group-hover:-translate-x-1">
                <path fill-rule="evenodd" d="M17 10a.75.75 0 0 1-.75.75H5.56l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L5.56 9.25H16.25A.75.75 0 0 1 17 10Z" clip-rule="evenodd" />
            </svg>
            Back to All Tools
        </a>
    </div>
    <header class="mb-10 text-center">
        <div class="flex items-center justify-center mb-4">
            <div class="w-16 h-16 bg-gradient-to-r from-green-500 to-teal-600 rounded-xl flex items-center justify-center p-2 mr-4">
                <img src={toolConfig.iconSrc} alt="{toolConfig.name} icon" class="w-full h-full object-contain" />
            </div>
            <h1 class="text-h1 text-theme-text-primary tracking-tight">{toolConfig.name}</h1>
        </div>
        <p class="mt-3 text-lg text-theme-text-secondary">
            {toolConfig.description}
        </p>
    </header>

    <!-- Input Form Section -->
    <div class="max-w-2xl mx-auto mb-8">
        <div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 p-6 shadow-card">
            <HerbiboarForm on:calculated={handleCalculation} on:error={handleCalcError} />
        </div>
    </div>

    <!-- Results/Error Section -->
    <div bind:this={resultsElement}>
        <!-- Error Display -->
        {#if currentError && currentError.trim() !== ''}
            <div class="max-w-4xl mx-auto mb-8">
                <div class="bg-gradient-to-r from-red-500/10 via-red-600/5 to-transparent border border-red-500/30 rounded-card p-4 animate-slide-down">
                    <div class="flex items-start space-x-3">
                        <div class="w-6 h-6 bg-red-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                            </svg>
                        </div>
                        <div>
                            <h4 class="text-sm font-semibold text-red-400 mb-1">Calculation Error</h4>
                            <p class="text-xs text-theme-text-secondary leading-relaxed">{currentError}</p>
                        </div>
                    </div>
                </div>
            </div>
        {/if}

        <!-- Results Section -->
        {#if currentApiResult}
            <div class="animate-slide-up">
                <HerbiboarResults result={currentApiResult} />
            </div>
        {/if}
    </div>

    <!-- Pro Tips Section -->
    <div class="max-w-4xl mx-auto mt-12">
        <div class="bg-theme-bg-secondary rounded-card border border-theme-border-accent/20 shadow-card overflow-hidden">
            <button 
                on:click={loadProTips}
                class="w-full px-6 py-4 text-left bg-gradient-to-r from-green-500/10 via-teal-500/5 to-transparent hover:from-green-500/15 hover:via-teal-500/10 transition-all duration-200 flex items-center justify-between group"
            >
                <div class="flex items-center space-x-3">
                    <div class="w-8 h-8 bg-gradient-to-r from-green-500 to-teal-600 rounded-lg flex items-center justify-center flex-shrink-0">
                        <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                        </svg>
                    </div>
                    <div>
                        <h3 class="text-lg font-semibold text-theme-text-primary">Calculation Methodology</h3>
                        <p class="text-sm text-theme-text-secondary">Learn how Herbiboar XP rates and profits are calculated</p>
                    </div>
                </div>
                <svg 
                    class="w-5 h-5 text-theme-text-secondary transition-transform duration-200 {showProTips ? 'rotate-180' : ''}" 
                    fill="none" 
                    stroke="currentColor" 
                    viewBox="0 0 24 24"
                >
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
            </button>

            {#if showProTips && proTipsData}
                <div class="px-6 pb-6 animate-slide-down">
                    <div class="border-t border-theme-border-accent/20 pt-6">
                        
                        <!-- XP Rate Methodology -->
                        <div class="mb-8">
                            <h4 class="text-md font-semibold text-theme-text-primary mb-3 flex items-center">
                                <svg class="w-5 h-5 mr-2 text-yellow-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6m0 10v-3m-3 3h.01M9 17h.01M9 14h.01M12 14h.01M15 11h.01M12 11h.01M9 11h.01M7 21h10a2 2 0 002-2V5a2 2 0 00-2-2H7a2 2 0 00-2 2v14a2 2 0 002 2z"></path>
                                </svg>
                                XP Rate Calculation
                            </h4>
                            <div class="bg-theme-bg-primary/50 rounded-lg p-4 mb-4">
                                <p class="text-sm text-theme-text-secondary mb-2">
                                    <strong>Source:</strong> {proTipsData.calculation_methodology.xp_rates_source}
                                </p>
                                <p class="text-sm text-theme-text-secondary mb-3">
                                    <strong>Formula:</strong> {proTipsData.calculation_methodology.base_formula}
                                </p>
                                
                                <h5 class="text-sm font-medium text-theme-text-primary mb-2">Key Data Points:</h5>
                                <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
                                    {#each proTipsData.calculation_methodology.data_points as point}
                                        <div class="flex justify-between items-center bg-theme-bg-secondary/50 rounded px-3 py-2">
                                            <span class="text-sm text-theme-text-secondary">Level {point.level}</span>
                                            <span class="text-sm font-medium text-theme-text-primary">{point.xp_per_catch} XP</span>
                                        </div>
                                    {/each}
                                </div>
                            </div>
                        </div>

                        <!-- Game Mechanics -->
                        <div class="mb-8">
                            <h4 class="text-md font-semibold text-theme-text-primary mb-3 flex items-center">
                                <svg class="w-5 h-5 mr-2 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 100 4m0-4v2m0-6V4"></path>
                                </svg>
                                Game Mechanics
                            </h4>
                            <div class="bg-theme-bg-primary/50 rounded-lg p-4">
                                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                    <div>
                                        <p class="text-sm text-theme-text-secondary mb-1">
                                            <strong>Catch Time:</strong> {proTipsData.game_mechanics.catch_time}
                                        </p>
                                        <p class="text-sm text-theme-text-secondary">
                                            <strong>Gear Effects:</strong> {proTipsData.game_mechanics.gear_effects}
                                        </p>
                                    </div>
                                    <div>
                                        <p class="text-sm font-medium text-theme-text-primary mb-1">Herb Yields:</p>
                                        <ul class="text-xs text-theme-text-secondary space-y-1">
                                            {#each proTipsData.game_mechanics.herb_yields as herbYield}
                                                <li>• {herbYield}</li>
                                            {/each}
                                        </ul>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <!-- Pro Tips -->
                        <div class="mb-8">
                            <h4 class="text-md font-semibold text-theme-text-primary mb-3 flex items-center">
                                <svg class="w-5 h-5 mr-2 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                                </svg>
                                Pro Tips
                            </h4>
                            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                {#each proTipsData.pro_tips as tip}
                                    <div class="bg-theme-bg-primary/50 rounded-lg p-4">
                                        <h5 class="text-sm font-medium text-theme-text-primary mb-2">{tip.tip}</h5>
                                        <p class="text-sm text-theme-text-secondary">{tip.description}</p>
                                    </div>
                                {/each}
                            </div>
                        </div>

                        <!-- Herb Calculation -->
                        <div>
                            <h4 class="text-md font-semibold text-theme-text-primary mb-3 flex items-center">
                                <svg class="w-5 h-5 mr-2 text-yellow-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                                </svg>
                                Herb Profit Calculation
                            </h4>
                            <div class="bg-theme-bg-primary/50 rounded-lg p-4">
                                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                                    <div>
                                        <p class="text-sm text-theme-text-secondary mb-2">
                                            <strong>Drop Rates:</strong> {proTipsData.herb_calculation.drop_rates}
                                        </p>
                                        <p class="text-sm text-theme-text-secondary">
                                            <strong>Price Sources:</strong> {proTipsData.herb_calculation.price_sources}
                                        </p>
                                    </div>
                                    <div>
                                        <p class="text-sm text-theme-text-secondary mb-2">
                                            <strong>Profit Factors:</strong> {proTipsData.herb_calculation.profit_factors}
                                        </p>
                                        <p class="text-sm text-theme-text-secondary">
                                            <strong>GP Calculation:</strong> {proTipsData.herb_calculation.gp_calculation}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>

                    </div>
                </div>
            {/if}
        </div>
    </div>
</div>