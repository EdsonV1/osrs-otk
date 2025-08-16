<script lang="ts">
    //import { createEventDispatcher } from 'svelte';
    import { createEventDispatcher, type EventDispatcher } from 'svelte'; 
    import type { BirdhouseApiInput, BirdhouseApiResult, BirdhouseFormState } from '$lib/types';

    interface ComponentEvents {
        calculated: { resultData: BirdhouseApiResult };
        error: { message: string };
    }

    const dispatch: EventDispatcher<ComponentEvents> = createEventDispatcher();

    const birdhouseTypes = [
        { value: "regular", label: "Regular", hunterLevel: 5, iconSrc: "/birdhouse/bird_house.png" },
        { value: "oak", label: "Oak", hunterLevel: 14, iconSrc: "/birdhouse/oak_bird_house.png" },            
        { value: "willow", label: "Willow", hunterLevel: 24, iconSrc: "/birdhouse/willow_bird_house.png" },      
        { value: "teak", label: "Teak", hunterLevel: 34, iconSrc: "/birdhouse/teak_bird_house.png" },        
        { value: "maple", label: "Maple", hunterLevel: 44, iconSrc: "/birdhouse/maple_bird_house.png" },       
        { value: "mahogany", label: "Mahogany", hunterLevel: 49, iconSrc: "/birdhouse/mahogany_bird_house.png" }, 
        { value: "yew", label: "Yew", hunterLevel: 59, iconSrc: "/birdhouse/yew_bird_house.png" },           
        { value: "magic", label: "Magic", hunterLevel: 74, iconSrc: "/birdhouse/magic_bird_house.png" },       
        { value: "redwood", label: "Redwood", hunterLevel: 89, iconSrc: "/birdhouse/redwood_bird_house.png" },   
    ];

    let formState: BirdhouseFormState = {
        selectedLogType: birdhouseTypes[0].value,
        totalBirdhouses: 0,
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

        if (!formState.totalBirdhouses || formState.totalBirdhouses <= 0) {
            dispatch('error', { message: 'Total birdhouses must be greater than 0.' });
            isLoading = false;
            return;
        }

        const _selectedTypeData = birdhouseTypes.find(b => b.value === formState.selectedLogType);


        const apiPayload: BirdhouseApiInput = {
            type: formState.selectedLogType,
            quantity: Number(formState.totalBirdhouses),
        };

        try {
            const response = await fetch('http://localhost:8080/api/birdhouse', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(apiPayload)
            });

            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(errorText || `API Error: ${response.status}`);
            }
            const resultData = await response.json() as BirdhouseApiResult;
            dispatch('calculated', { resultData });
        } catch (err: any) {
            dispatch('error', { message: err.message || 'Calculation failed.' });
        } finally {
            isLoading = false;
        }
    }
</script>

<form class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8" on:submit|preventDefault={handleSubmit}>
    <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Birdhouse Inputs</h2>

    <fieldset class="space-y-4">
        <legend class="{fieldsetLegendClasses} mb-3">Type of Birdhouses</legend>
        
        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 gap-3">
            {#each birdhouseTypes as typeOpt (typeOpt.value)}
                <label class="{radioLabelClasses} {formState.selectedLogType === typeOpt.value ? 'ring-2 ring-theme-accent border-theme-accent shadow-md' : 'bg-gray-700/30 hover:bg-gray-700/60'}">
                    <input
                        type="radio"
                        name="birdhouseLogType"
                        bind:group={formState.selectedLogType}
                        value={typeOpt.value}
                        class="{radioInputClasses} sr-only" 
                    />
                    <img src={typeOpt.iconSrc} alt="{typeOpt.label} icon" class="w-8 h-8 object-contain flex-shrink-0"/>
                    <span class="text-sm font-medium text-theme-text-primary">{typeOpt.label}</span>
                </label>
            {/each}
        </div>
    </fieldset>

    <fieldset class="space-y-4">
        <legend class={fieldsetLegendClasses}>Calculation Details</legend>
        <div>
            <label for="totalBirdhouses" class={labelBaseClasses}>Total Number of Birdhouses to Place/Calculate:</label>
            <input type="number" id="totalBirdhouses" bind:value={formState.totalBirdhouses} min="1" step="1" max="999999" placeholder="e.g., 100" class="{inputBaseClasses} mt-1">
        </div>
    </fieldset>
    
    <button type="submit" disabled={isLoading}
            class="w-full flex justify-center py-3 px-4 border border-transparent rounded-lg shadow-button text-sm font-semibold text-white bg-theme-accent hover:bg-theme-accent-hover focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-theme-card-bg focus:ring-theme-accent disabled:opacity-60 disabled:cursor-not-allowed transition-colors duration-150">
        {isLoading ? 'Calculating...' : 'Calculate Birdhouses'}
    </button>
</form>
