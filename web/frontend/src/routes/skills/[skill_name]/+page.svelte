<script lang="ts">
    import type { PageData } from './$types';
    import type { TrainingMethod, SkillCalculationOutput, CalculatedMethodResult } from '$lib/types';
    import { getXpForLevel, getLevelForXp } from '$lib/xp_table';

    export let data: PageData;

    $: skillInfo = data.skillData;
    // Sort all methods initially by level requirement, then by XP rate descending
    $: allTrainingMethods = data.skillData?.trainingMethods.sort((a,b) => a.levelReq - b.levelReq || b.xpRate - a.xpRate) || [];

    let currentInputMode: 'xp' | 'level' = 'level';
    let currentLevelState: number = 1;
    let currentXPState: number = getXpForLevel(currentLevelState);

    let targetInputMode: 'xp' | 'level' = 'level';
    let targetLevelState: number = Math.min(currentLevelState + 1 > 99 ? 99 : currentLevelState + 1, 99);
    let targetXPState: number = getXpForLevel(targetLevelState);
    
    let calculationError: string | null = null;

    function syncCurrentXP() { if (currentInputMode === 'level') currentXPState = getXpForLevel(currentLevelState); }
    function syncCurrentLevel() { if (currentInputMode === 'xp') currentLevelState = getLevelForXp(currentXPState); }
    function syncTargetXP() { if (targetInputMode === 'level') targetXPState = getXpForLevel(targetLevelState); }
    function syncTargetLevel() { if (targetInputMode === 'xp') targetLevelState = getLevelForXp(targetXPState); }

    $: currentLevelState, syncCurrentXP();
    $: currentXPState, syncCurrentLevel();
    $: targetLevelState, syncTargetXP();
    $: targetXPState, syncTargetLevel();

    $: {
        let baseLevelForTarget = (currentInputMode === 'level') ? currentLevelState : getLevelForXp(currentXPState);
        let baseXPForTarget = currentXPState;
        let nextPossibleLevel = Math.min(baseLevelForTarget + 1, 99);
        
        if (targetXPState <= currentXPState) {
            calculationError = "Target must be greater than current progress.";
        } else {
            calculationError = null;
        }
    }
    
    $: currentActualLevel = getLevelForXp(currentXPState);

    interface MethodCalcData {
        actionsNeeded?: number;
        timeToTargetHours?: number;
        timeToTargetFormatted?: string;
        marksOfGraceEarned?: number;
    }

    function calculateMethodMetrics(method: TrainingMethod, currentXP: number, targetXP: number): MethodCalcData | null {
        if (targetXP <= currentXP || !skillInfo) return null;

        const totalXPToGain = targetXP - currentXP;
        let timeToTargetHours: number | undefined = undefined;
        let actionsNeeded: number | undefined = undefined;
        let marksOfGraceEarned: number | undefined = undefined;

        if (method.xpRate > 0) {
            timeToTargetHours = totalXPToGain / method.xpRate;
        }
        if (method.xpPerAction && method.xpPerAction > 0) {
            actionsNeeded = Math.ceil(totalXPToGain / method.xpPerAction);
        }
        
        if (skillInfo.skillNameCanonical === 'agility') {
            if (method.marksPerHour !== undefined && timeToTargetHours !== undefined) {
                marksOfGraceEarned = method.marksPerHour * timeToTargetHours;
            }
        }

        return {
            actionsNeeded,
            timeToTargetHours,
            timeToTargetFormatted: formatHours(timeToTargetHours),
            marksOfGraceEarned,
        };
    }

    function formatHours(hours: number | undefined): string | undefined {
        if (hours === undefined || isNaN(hours) || !isFinite(hours)) return undefined;
        if (hours === 0 && !(hours > 0)) return '0h 0m';
        const totalMinutes = Math.round(hours * 60);
        if (totalMinutes === 0 && hours > 0) return '<1m';
        const h = Math.floor(totalMinutes / 60);
        const m = totalMinutes % 60;
        return `${h}h ${m}m`;
    }
    function formatNum(num: number | null | undefined, decimals: number = 0): string {
        if (num === null || num === undefined || typeof num !== 'number' || isNaN(num)) return 'N/A';
        return num.toLocaleString(undefined, { minimumFractionDigits: decimals, maximumFractionDigits: decimals });
    }

    const inputBaseClasses = "block w-full rounded-md border-0 py-2 px-3.5 bg-gray-700/50 text-theme-text-primary shadow-sm ring-1 ring-inset ring-theme-border-input placeholder:text-theme-text-tertiary focus:ring-2 focus:ring-inset focus:ring-theme-accent sm:text-sm sm:leading-6 shadow-inner-border transition-colors duration-150";
    const radioLabelClasses = "flex items-center text-sm text-theme-text-secondary cursor-pointer";
    const radioInputClasses = "h-4 w-4 border-theme-border-input text-theme-accent focus:ring-theme-accent focus:ring-offset-theme-card-bg";
    const fieldsetLegendClasses = "text-base font-semibold leading-6 text-theme-text-primary mb-2";
</script>

<div class="max-w-6xl mx-auto space-y-10 py-8 px-4 sm:px-6 lg:px-8">
    <div class="mb-6"> 
        <a href="/skills" class="inline-flex items-center text-sm font-medium text-theme-accent hover:text-theme-accent-hover transition-colors group">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="w-5 h-5 mr-1.5 transform transition-transform group-hover:-translate-x-1">
                <path fill-rule="evenodd" d="M17 10a.75.75 0 0 1-.75.75H5.56l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L5.56 9.25H16.25A.75.75 0 0 1 17 10Z" clip-rule="evenodd" />
            </svg>
            Back to All Skills
        </a>
    </div>
    <header class="text-center">
        {#if skillInfo}
            <h1 class="text-h1 text-theme-text-primary tracking-tight">{skillInfo.skillNameDisplay} Calculator</h1>
            {#if skillInfo.description}
                <p class="mt-3 text-lg text-theme-text-secondary max-w-2xl mx-auto">
                    {skillInfo.description}
                </p>
            {/if}
        {:else}
             <h1 class="text-h1 text-theme-text-primary tracking-tight">Skill Calculator</h1>
            <p class="mt-3 text-lg text-theme-text-secondary max-w-2xl mx-auto">Loading skill data...</p>
        {/if}
    </header>

    {#if skillInfo}
        <section class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8">
            <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-3">Your Progress & Goals</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-6">
                <div class="space-y-3">
                    <p class="text-base font-semibold text-theme-text-primary">Current {skillInfo.skillNameDisplay}</p>
                    <div class="flex items-center space-x-4">
                        <label class={radioLabelClasses}>
                            <input type="radio" bind:group={currentInputMode} value="level" name="currentInputMode" class={radioInputClasses}>
                            <span class="ml-2">Level</span>
                        </label>
                        <label class={radioLabelClasses}>
                            <input type="radio" bind:group={currentInputMode} value="xp" name="currentInputMode" class={radioInputClasses}>
                            <span class="ml-2">XP</span>
                        </label>
                    </div>
                    {#if currentInputMode === 'level'}
                        <input type="number" placeholder="Current Level" bind:value={currentLevelState} min="1" max="99" step="1" class={inputBaseClasses}>
                        <p class="text-xs text-theme-text-tertiary pl-1">XP: {formatNum(currentXPState)}</p>
                    {:else}
                        <input type="number" placeholder="Current XP" bind:value={currentXPState} min="0" max="200000000" step="1" class={inputBaseClasses}>
                        <p class="text-xs text-theme-text-tertiary pl-1">Level: {getLevelForXp(currentXPState)}</p>
                    {/if}
                </div>

                <div class="space-y-3">
                    <p class="text-base font-semibold text-theme-text-primary">Target {skillInfo.skillNameDisplay}</p>
                     <div class="flex items-center space-x-4">
                        <label class={radioLabelClasses}>
                            <input type="radio" bind:group={targetInputMode} value="level" name="targetInputMode" class={radioInputClasses}>
                            <span class="ml-2">Level</span>
                        </label>
                        <label class={radioLabelClasses}>
                            <input type="radio" bind:group={targetInputMode} value="xp" name="targetInputMode" class={radioInputClasses}>
                            <span class="ml-2">XP</span>
                        </label>
                    </div>
                    {#if targetInputMode === 'level'}
                        <input type="number" placeholder="Target Level" bind:value={targetLevelState} min={currentLevelState + 1 > 99 ? 99 : currentLevelState + 1} max="99" step="1" class={inputBaseClasses}>
                         <p class="text-xs text-theme-text-tertiary pl-1">XP: {formatNum(targetXPState)}</p>
                    {:else}
                        <input type="number" placeholder="Target XP" bind:value={targetXPState} min={currentXPState + 1} max="200000000" step="1" class={inputBaseClasses}>
                        <p class="text-xs text-theme-text-tertiary pl-1">Level: {getLevelForXp(targetXPState)}</p>
                    {/if}
                </div>
            </div>
            </section>

        {#if calculationError}
            <div class="bg-red-900/80 border border-red-700 text-red-100 px-4 py-3 rounded-lg shadow-md text-sm" role="alert">
                <div class="flex items-center">
                    <svg class="fill-current h-5 w-5 text-red-400 mr-3" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M2.93 17.07A10 10 0 1 1 17.07 2.93 10 10 0 0 1 2.93 17.07zM11.414 10l2.829-2.829a1 1 0 0 0-1.414-1.414L10 8.586 7.172 5.757a1 1 0 0 0-1.414 1.414L8.586 10l-2.829 2.829a1 1 0 1 0 1.414 1.414L10 11.414l2.829 2.829a1 1 0 0 0 1.414-1.414L11.414 10z"/></svg>
                    <p>{calculationError}</p>
                </div>
            </div>
        {/if}

        <section class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-8 overflow-hidden">
            <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-3 mb-0 px-6 sm:px-8 pt-6 sm:pt-8">
                Training Methods for {skillInfo.skillNameDisplay}
            </h2>
            {#if allTrainingMethods.length > 0}
                <div class="overflow-x-auto m-10"> 
                    <table class="min-w-full divide-y divide-theme-border-subtle text-sm">
                        <thead class="bg-gray-700/30 sticky top-0 z-10">
                            <tr>
                                <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Method</th>
                                <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Lvl</th>
                                <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">XP/hr</th>
                                {#if skillInfo.skillNameCanonical === 'agility'}
                                    <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Marks/hr</th>
                                {/if}
                                <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Actions to Target</th>
                                <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Time to Target</th>
                                {#if skillInfo.skillNameCanonical === 'agility'}
                                    <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary">Total Marks to Target</th>
                                {/if}
                                <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary hidden md:table-cell">Type</th>
                                <th scope="col" class="px-3 py-3.5 text-left font-semibold text-theme-text-primary hidden lg:table-cell">Location</th>
                            </tr>
                        </thead>
                        <tbody class="divide-y divide-theme-border-subtle/50">
                            {#each allTrainingMethods as method (method.id)}
                                {@const isAvailable = currentActualLevel >= method.levelReq}
                                {@const methodCalcs = calculateMethodMetrics(method, currentXPState, targetXPState)}
                                {@const colspanForNotes = 3 + (skillInfo.skillNameCanonical === 'agility' ? 1 : 0) + 2 + (skillInfo.skillNameCanonical === 'agility' ? 1 : 0) + 1 + 1}
                                
                                <tr class:opacity-60={!isAvailable} class:pointer-events-none={!isAvailable} class="hover:bg-gray-700/10">
                                    <td class="whitespace-nowrap px-3 py-3 font-medium {isAvailable ? 'text-theme-text-primary' : 'text-theme-text-tertiary'}">{method.name}</td>
                                    <td class="whitespace-nowrap px-3 py-3 {isAvailable ? 'text-theme-text-secondary' : 'text-theme-text-tertiary'}">{method.levelReq}</td>
                                    <td class="whitespace-nowrap px-3 py-3 {isAvailable ? 'text-theme-text-secondary' : 'text-theme-text-tertiary'}">{formatNum(method.xpRate)}</td>
                                    {#if skillInfo.skillNameCanonical === 'agility'}
                                        <td class="whitespace-nowrap px-3 py-3 {isAvailable ? 'text-theme-text-secondary' : 'text-theme-text-tertiary'}">{method.marksPerHour ? formatNum(method.marksPerHour, 1) : 'N/A'}</td>
                                    {/if}
                                    
                                    <td class="whitespace-nowrap px-3 py-3 {isAvailable ? 'text-theme-text-secondary' : 'text-theme-text-tertiary'}">
                                        {methodCalcs && methodCalcs.actionsNeeded !== undefined ? formatNum(methodCalcs.actionsNeeded) : (targetXPState <= currentXPState ? '-' : 'N/A')}
                                    </td>
                                    <td class="whitespace-nowrap px-3 py-3 {isAvailable ? 'text-theme-text-secondary' : 'text-theme-text-tertiary'}">
                                        {methodCalcs?.timeToTargetFormatted || (targetXPState <= currentXPState ? '-' : 'N/A')}
                                    </td>
                                    {#if skillInfo.skillNameCanonical === 'agility'}
                                        <td class="whitespace-nowrap px-3 py-3 {isAvailable ? 'text-theme-text-secondary' : 'text-theme-text-tertiary'}">
                                            {methodCalcs && methodCalcs.marksOfGraceEarned !== undefined ? formatNum(methodCalcs.marksOfGraceEarned, 1) : (targetXPState <= currentXPState ? '-' : 'N/A')}
                                        </td>
                                    {/if}
                                    <td class="whitespace-nowrap px-3 py-3 {isAvailable ? 'text-theme-text-secondary' : 'text-theme-text-tertiary'} hidden md:table-cell">{method.type || 'N/A'}</td>
                                    <td class="px-3 py-3 {isAvailable ? 'text-theme-text-secondary' : 'text-theme-text-tertiary'} hidden lg:table-cell min-w-[150px]">{method.location || 'N/A'}</td>
                                </tr>
                                {#if method.notes || (method.itemsRequired && method.itemsRequired.length > 0) || (method.questsRequired && method.questsRequired.length > 0)}
                                <tr class:opacity-60={!isAvailable} class="bg-gray-700/10 hover:bg-gray-700/20">
                                    <td class="px-3 py-2 italic text-xs {isAvailable ? 'text-theme-text-tertiary' : 'text-gray-500'}" colspan={colspanForNotes}> 
                                        {#if method.itemsRequired && method.itemsRequired.length > 0}<strong>Items:</strong> {method.itemsRequired.join(', ')}{/if}
                                        {#if method.questsRequired && method.questsRequired.length > 0}
                                            <br class:mt-1={method.itemsRequired && method.itemsRequired.length > 0}><strong>Quests:</strong> {method.questsRequired.join(', ')}
                                        {/if}
                                        {#if method.notes}
                                            <br class:mt-1={(method.itemsRequired && method.itemsRequired.length > 0) || (method.questsRequired && method.questsRequired.length > 0)}>
                                            <em>{method.notes}</em>
                                        {/if}
                                    </td>
                                </tr>
                                {/if}
                            {/each}
                        </tbody>
                    </table>
                </div>
            {:else}
                <p class="text-theme-text-secondary px-6 sm:px-8">No training methods available for {skillInfo.skillNameDisplay}.</p>
            {/if}
        </section>
    {:else}
        <div class="text-center py-10">
            <p class="text-theme-text-secondary text-lg">Skill data not loaded or skill not found. Check backend and API path.</p>
        </div>
    {/if}
</div>
