<script lang="ts">
    import { createEventDispatcher, type EventDispatcher } from 'svelte';
    import type { ArdyKnightInput, ArdyKnightResult } from '$lib/types';

    export let toolName: string = "Inputs";

    interface ComponentEvents {
        calculated: { data: ArdyKnightResult };
        error: { error: string };
    }
    const dispatch: EventDispatcher<ComponentEvents> = createEventDispatcher();

    let currentProgressType: 'xp' | 'level' = 'level';
    let currentXP: number | null = null;
    let currentLevel: number | null = null;
    let targetProgressType: 'xp' | 'level' = 'level';
    let targetXP: number | null = null;
    let targetLevel: number | null = null;
    let hasArdyMed: boolean = false;
    let hasThievingCape: boolean = false;
    let hasRoguesOutfit: boolean = false;
    let hasShadowVeil: boolean = false;
    let hourlyPickpockets: number | null = 2500;
    let foodHealAmount: number | null = 20;
    let foodCost: number | null = 200;
    let isLoading: boolean = false;

    async function handleSubmit() {
        isLoading = true;
        dispatch('error', { error: '' });

        console.log('--- handleSubmit triggered ---');
        console.log('Current Input States:', {
            currentProgressType, currentXP, currentLevel,
            targetProgressType, targetXP, targetLevel,
            hasArdyMed, hasThievingCape, hasRoguesOutfit, hasShadowVeil,
            hourlyPickpockets, foodHealAmount, foodCost
        });

        const payload: Partial<ArdyKnightInput> = {
            has_ardy_med: hasArdyMed,
            has_thieving_cape: hasThievingCape,
            has_rogues_outfit: hasRoguesOutfit,
            has_shadow_veil: hasShadowVeil,
            hourly_pickpockets: Number(hourlyPickpockets) || 0,
            food_heal_amount: Number(foodHealAmount) || 0,
            food_cost: Number(foodCost) || 0,
        };

        if (currentProgressType === 'xp') {
            if (currentXP !== null && String(currentXP).trim() !== '') payload.current_thieving_xp = Number(currentXP);
        } else {
            if (currentLevel !== null && String(currentLevel).trim() !== '') payload.current_thieving_level = Number(currentLevel);
        }
        if (targetProgressType === 'xp') {
            if (targetXP !== null && String(targetXP).trim() !== '') payload.target_thieving_xp = Number(targetXP);
        } else {
            if (targetLevel !== null && String(targetLevel).trim() !== '') payload.target_thieving_level = Number(targetLevel);
        }
        
        console.log('Payload before validation:', JSON.parse(JSON.stringify(payload)));

        if (payload.current_thieving_xp === undefined && payload.current_thieving_level === undefined) {
            dispatch('error', { error: 'Provide current XP or Level.' }); isLoading = false; return;
        }
        if (payload.target_thieving_xp === undefined && payload.target_thieving_level === undefined) {
            dispatch('error', { error: 'Provide target XP or Level.' }); isLoading = false; return;
        }
        if (!payload.hourly_pickpockets || payload.hourly_pickpockets <= 0) {
            dispatch('error', { error: 'Hourly pickpockets > 0.' }); isLoading = false; return;
        }

        try {
            const apiEndpoint = 'http://localhost:8080/api/ardyknights';
            console.log("Sending final payload to API:", JSON.stringify(payload));
            const response = await fetch(apiEndpoint, { method: 'POST', headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(payload) });
            if (!response.ok) { const errorText = await response.text(); throw new Error(errorText || `HTTP error ${response.status}`); }
            const data = await response.json() as ArdyKnightResult;
            dispatch('calculated', { data });
        } catch (err: any) { 
            dispatch('error', { error: err.message || 'Unknown calculation error.' });
        } finally { 
            isLoading = false; 
        }
    }

    const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3 bg-gray-700 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border";
    const radioCheckboxBaseClasses = "h-4 w-4 border-theme-border-input text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg";
    const fieldsetLegendClasses = "text-base font-semibold leading-6 text-theme-text-primary";
    const labelTextClasses = "text-sm text-theme-text-secondary";
    const labelContainerClasses = "flex items-center cursor-pointer";
</script>

<form class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8" on:submit|preventDefault={handleSubmit}>
    <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">{toolName} Inputs</h2>

    <fieldset class="space-y-3">
        <legend class={fieldsetLegendClasses}>Current Progress</legend>
        <div class="flex items-center space-x-6">
            <label class="{labelContainerClasses}">
                <input type="radio" bind:group={currentProgressType} value="level" name="currentProgressTypeRadio" class="{radioCheckboxBaseClasses}">
                <span class="ml-2 {labelTextClasses}">Level</span>
            </label>
            <label class="{labelContainerClasses}">
                <input type="radio" bind:group={currentProgressType} value="xp" name="currentProgressTypeRadio" class="{radioCheckboxBaseClasses}">
                <span class="ml-2 {labelTextClasses}">XP</span>
            </label>
        </div>
        {#if currentProgressType === 'xp'}
            <input type="number" placeholder="Current Thieving XP" bind:value={currentXP} min="0" step="1" class="{inputBaseClasses}">
        {:else}
            <input type="number" placeholder="Current Thieving Level" bind:value={currentLevel} min="1" max="99" step="1" class="{inputBaseClasses}">
        {/if}
    </fieldset>

    <fieldset class="space-y-3">
        <legend class={fieldsetLegendClasses}>Target Progress</legend>
        <div class="flex items-center space-x-6">
            <label class="{labelContainerClasses}">
                <input type="radio" bind:group={targetProgressType} value="level" name="targetProgressTypeRadio" class="{radioCheckboxBaseClasses}">
                <span class="ml-2 {labelTextClasses}">Level</span>
            </label>
            <label class="{labelContainerClasses}">
                <input type="radio" bind:group={targetProgressType} value="xp" name="targetProgressTypeRadio" class="{radioCheckboxBaseClasses}">
                <span class="ml-2 {labelTextClasses}">XP</span>
            </label>
        </div>
        {#if targetProgressType === 'xp'}
            <input type="number" placeholder="Target Thieving XP" bind:value={targetXP} min="0" step="1" class="{inputBaseClasses}">
        {:else}
            <input type="number" placeholder="Target Thieving Level" bind:value={targetLevel} min="1" max="99" step="1" class="{inputBaseClasses}">
        {/if}
    </fieldset>

    <fieldset class="space-y-3">
        <legend class={fieldsetLegendClasses}>Boosts & Gear</legend>
        <label class="{labelContainerClasses}">
            <input type="checkbox" bind:checked={hasArdyMed} class="{radioCheckboxBaseClasses} rounded">
            <span class="ml-2 {labelTextClasses}">Ardougne Med Diary</span>
        </label>
        <label class="{labelContainerClasses}">
            <input type="checkbox" bind:checked={hasThievingCape} class="{radioCheckboxBaseClasses} rounded">
            <span class="ml-2 {labelTextClasses}">Thieving Cape</span>
        </label>
        <label class="{labelContainerClasses}">
            <input type="checkbox" bind:checked={hasRoguesOutfit} class="{radioCheckboxBaseClasses} rounded">
            <span class="ml-2 {labelTextClasses}">Rogue's Outfit</span>
        </label>
        <label class="{labelContainerClasses}">
            <input type="checkbox" bind:checked={hasShadowVeil} class="{radioCheckboxBaseClasses} rounded">
            <span class="ml-2 {labelTextClasses}">Shadow Veil</span>
        </label>
    </fieldset>
    
    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Parameters</legend>
        <div>
            <label for="hourlyPickpockets" class="block text-sm font-medium leading-6 {labelTextClasses}">Pickpockets/hour:</label>
            <input type="number" id="hourlyPickpockets" bind:value={hourlyPickpockets} min="1" step="1" class="mt-1 {inputBaseClasses}">
        </div>
        <div>
            <label for="foodHealAmount" class="block text-sm font-medium leading-6 {labelTextClasses}">Food Heal Amount:</label>
            <input type="number" id="foodHealAmount" bind:value={foodHealAmount} min="0" step="1" class="mt-1 {inputBaseClasses}">
        </div>
        <div>
            <label for="foodCost" class="block text-sm font-medium leading-6 {labelTextClasses}">Food Cost (GP):</label>
            <input type="number" id="foodCost" bind:value={foodCost} min="0" step="1" class="mt-1 {inputBaseClasses}">
        </div>
    </fieldset>

    <button type="submit" disabled={isLoading}
            class="w-full flex justify-center py-2.5 px-4 border border-transparent rounded-md shadow-button text-sm font-medium text-white bg-theme-accent hover:bg-theme-accent-hover focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-theme-card-bg focus:ring-theme-accent disabled:opacity-60 disabled:cursor-not-allowed transition-colors">
        {isLoading ? 'Calculating...' : 'Calculate'}
    </button>
</form>