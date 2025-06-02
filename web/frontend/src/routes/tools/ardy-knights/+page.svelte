<script lang="ts">
    import type { ArdyKnightResult } from '$lib/types';
    import InputForm from '$lib/InputForm.svelte';
    import ResultsDisplay from '$lib/ResultsDisplay.svelte';

    let results: ArdyKnightResult | null = null;
    let error: string | null = null;

    function handleCalculated(event: CustomEvent<{ data: ArdyKnightResult }>) {
        results = event.detail.data;
        error = null;
    }

    function handleError(event: CustomEvent<{ error: string }>) {
        error = event.detail.error;
        results = null;
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
        <h1 class="text-h1 text-theme-text-primary tracking-tight mt-10">Ardougne Knight Calculator</h1>
        <p class="mt-3 text-lg text-theme-text-secondary">Calculate your Thieving XP, GP, and more.</p>
    </header>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 items-start">
        <div class="lg:col-span-1">
            <InputForm on:calculated={handleCalculated} on:error={handleError} toolName="Ardougne Knights" />
        </div>

        <div class="lg:col-span-1 lg:sticky lg:top-24">
            {#if error && error.trim() !== ''}
                <div class="bg-red-900 border border-red-700 text-red-100 px-4 py-3 rounded-md mb-6" role="alert">
                    <strong class="font-bold">Error:</strong>
                    <span class="block sm:inline ml-1">{error}</span>
                </div>
            {/if}

            {#if results}
                <ResultsDisplay {results} />
            {:else if !error || error.trim() === ''}
                <div class="bg-theme-card-bg p-6 rounded-lg shadow-card border border-theme-border text-center text-theme-text-secondary">
                    Results will appear here once calculated.
                </div>
            {/if}
        </div>
    </div>
</div>