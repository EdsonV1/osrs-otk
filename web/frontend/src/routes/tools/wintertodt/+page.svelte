<script lang="ts">
    import type { WintertodtApiResult } from '$lib/types';
    import WintertodtForm from '$lib/components/wintertodt/InputForm.svelte';
    import WintertodtResults from '$lib/components/wintertodt/ResultsDisplay.svelte';
    import { tick } from 'svelte';

    // Tool configuration
    const toolConfig = {
        name: 'Wintertodt Calculator',
        description: 'Calculate experience, loot, and Phoenix pet chances from Wintertodt based on your Firemaking level and planned rounds.',
        iconSrc: '/images/skills/firemaking.png'
    };

    let currentApiResult: WintertodtApiResult | null = null;
    let currentError: string | null = null;
    let resultsElement: HTMLElement;

    function scrollToResults() {
        if (resultsElement) {
            resultsElement.scrollIntoView({ 
                behavior: 'smooth', 
                block: 'start',
                inline: 'nearest'
            });
        }
    }

    async function handleCalculation(event: CustomEvent<{ resultData: WintertodtApiResult }>) {
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
            <div class="w-16 h-16 bg-gradient-to-r from-orange-500 to-red-600 rounded-xl flex items-center justify-center p-2 mr-4">
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
            <WintertodtForm on:calculated={handleCalculation} on:error={handleCalcError} />
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
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
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
                <WintertodtResults apiResult={currentApiResult} iconSrc={toolConfig.iconSrc} />
            </div>
        {/if}
    </div>
</div>