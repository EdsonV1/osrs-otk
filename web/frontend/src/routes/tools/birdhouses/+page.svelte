<script lang="ts">
    import type { BirdhouseApiResult } from '$lib/types';
    import BirdhouseForm from '$lib/BirdhouseInputForm.svelte';
    import BirdhouseResults from '$lib/BirdhouseResultsDisplay.svelte';

    let currentApiResult: BirdhouseApiResult | null = null;
    let currentError: string | null = null;

    function handleCalculation(event: CustomEvent<{ resultData: BirdhouseApiResult }>) {
        currentApiResult = event.detail.resultData;
        currentError = null;
    }

    function handleCalcError(event: CustomEvent<{ message: string }>) {
        currentError = event.detail.message;
        currentApiResult = null;
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
        <h1 class="text-h1 text-theme-text-primary tracking-tight mt-10">Birdhouse Run Calculator</h1>
        <p class="mt-3 text-lg text-theme-text-secondary">
            Estimate XP, nests, and valuable loot from your birdhouse runs based on log type and total houses.
        </p>
    </header>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 items-start">
        <div class="lg:col-span-2">
            <BirdhouseForm on:calculated={handleCalculation} on:error={handleCalcError} />
        </div>

        <div class="lg:col-span-1 lg:sticky lg:top-24">
            {#if currentError && currentError.trim() !== ''}
                <div class="bg-red-900/80 border border-red-700 text-red-100 px-4 py-3 rounded-lg mb-6 shadow-md" role="alert">
                    <div class="flex">
                        <div class="py-1">
                            <svg class="fill-current h-6 w-6 text-red-500 mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zM11.414 10l2.829-2.829a1 1 0 0 0-1.414-1.414L10 8.586 7.172 5.757a1 1 0 0 0-1.414 1.414L8.586 10l-2.829 2.829a1 1 0 1 0 1.414 1.414L10 11.414l2.829 2.829a1 1 0 0 0 1.414-1.414L11.414 10z"/></svg>
                        </div>
                        <div>
                            <p class="font-bold">Calculation Error</p>
                            <p class="text-sm">{currentError}</p>
                        </div>
                    </div>
                </div>
            {/if}

            {#if currentApiResult}
                <BirdhouseResults apiResult={currentApiResult} />
            {:else if !currentError || currentError.trim() === ''}
                <div class="bg-theme-card-bg p-6 rounded-lg shadow-card border border-theme-border text-center text-theme-text-secondary">
                    Results will appear here once calculated.
                </div>
            {/if}
        </div>
    </div>
</div>