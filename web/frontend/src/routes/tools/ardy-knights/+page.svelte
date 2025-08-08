<script lang="ts">
    import type { ArdyKnightResult } from '$lib/types';
    import InputForm from '$lib/components/ardy-knights/InputForm.svelte';
    import ResultsDisplay from '$lib/components/ardy-knights/ResultsDisplay.svelte';
    import { tick } from 'svelte';

    let results: ArdyKnightResult | null = null;
    let error: string | null = null;
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

    async function handleCalculated(event: CustomEvent<{ data: ArdyKnightResult }>) {
        results = event.detail.data;
        error = null;
        
        // Wait for the DOM to update, then scroll to results
        await tick();
        setTimeout(() => scrollToResults(), 100);
    }

    async function handleError(event: CustomEvent<{ error: string }>) {
        error = event.detail.error;
        results = null;
        
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
        <h1 class="text-h1 text-theme-text-primary tracking-tight">Ardougne Knight Calculator</h1>
        <p class="mt-3 text-lg text-theme-text-secondary">Calculate your Thieving XP, GP, and more.</p>
    </header>

    <!-- Input Form Section -->
    <div class="max-w-2xl mx-auto mb-8">
        <div class="bg-glass backdrop-blur-md rounded-card border border-theme-border-accent/20 p-6 shadow-card">
            <InputForm on:calculated={handleCalculated} on:error={handleError} />
        </div>
    </div>

    <!-- Results/Error Section -->
    <div bind:this={resultsElement}>
        <!-- Error Display -->
        {#if error && error.trim() !== ''}
            <div class="max-w-4xl mx-auto mb-8">
                <div class="bg-gradient-to-r from-red-500/10 via-red-600/5 to-transparent border border-red-500/30 rounded-card p-4 animate-slide-down">
                    <div class="flex items-start space-x-3">
                        <div class="w-6 h-6 bg-red-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                            <svg class="w-4 h-4 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.964-.833-2.732 0L3.732 16.5c-.77.833.19 2.5 1.732 2.5z"></path>
                            </svg>
                        </div>
                        <div>
                            <h4 class="text-sm font-semibold text-red-400 mb-1">Calculation Error</h4>
                            <p class="text-xs text-theme-text-secondary leading-relaxed">{error}</p>
                        </div>
                    </div>
                </div>
            </div>
        {/if}

        <!-- Results Section -->
        {#if results}
            <div class="animate-slide-up">
                <ResultsDisplay {results} />
            </div>
        {/if}
    </div>
</div>