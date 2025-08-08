<script lang="ts">
    import type { WintertodtApiResult } from './types';

    export let apiResult: WintertodtApiResult;

    function formatNumber(num: number): string {
        return new Intl.NumberFormat().format(Math.round(num));
    }

    function formatTime(hours: number): string {
        if (hours < 1) {
            return `${Math.round(hours * 60)}m`;
        } else if (hours < 24) {
            return `${hours.toFixed(1)}h`;
        } else {
            const days = Math.floor(hours / 24);
            const remainingHours = Math.round(hours % 24);
            return `${days}d ${remainingHours}h`;
        }
    }

    function formatGP(value: number): string {
        if (value >= 1000000) {
            return `${(value / 1000000).toFixed(1)}M gp`;
        } else if (value >= 1000) {
            return `${(value / 1000).toFixed(0)}K gp`;
        } else {
            return `${value} gp`;
        }
    }

    $: sortedLoot = Object.entries(apiResult.estimated_loot)
        .filter(([_, quantity]) => quantity > 0)
        .sort((a, b) => b[1] - a[1]);
</script>

<div class="bg-theme-card-bg p-6 rounded-lg shadow-card border border-theme-border">
    <h2 class="text-h2 text-theme-text-primary mb-6">Wintertodt Results</h2>
    
    <!-- Experience & Time Summary -->
    <div class="grid grid-cols-2 gap-4 mb-6">
        <div class="bg-theme-surface p-4 rounded-md">
            <h3 class="text-sm font-medium text-theme-text-secondary mb-2">Total Experience</h3>
            <p class="text-2xl font-bold text-theme-accent">{formatNumber(apiResult.total_experience)}</p>
        </div>
        <div class="bg-theme-surface p-4 rounded-md">
            <h3 class="text-sm font-medium text-theme-text-secondary mb-2">Experience Rate</h3>
            <p class="text-2xl font-bold text-theme-accent">{formatNumber(apiResult.average_exp_hour)}/hr</p>
        </div>
    </div>

    <div class="grid grid-cols-2 gap-4 mb-6">
        <div class="bg-theme-surface p-4 rounded-md">
            <h3 class="text-sm font-medium text-theme-text-secondary mb-2">Total Time</h3>
            <p class="text-xl font-semibold text-theme-text-primary">{formatTime(apiResult.total_time)}</p>
        </div>
        <div class="bg-theme-surface p-4 rounded-md">
            <h3 class="text-sm font-medium text-theme-text-secondary mb-2">Phoenix Pet Chance</h3>
            <p class="text-xl font-semibold text-orange-400">{apiResult.pet_chance.toFixed(2)}%</p>
        </div>
    </div>

    <!-- Loot Value -->
    <div class="bg-theme-surface p-4 rounded-md mb-6">
        <h3 class="text-sm font-medium text-theme-text-secondary mb-2">Estimated Loot Value</h3>
        <p class="text-2xl font-bold text-green-400">{formatGP(apiResult.total_value)}</p>
    </div>

    <!-- Detailed Loot Breakdown -->
    <div class="bg-theme-surface p-4 rounded-md">
        <h3 class="text-sm font-medium text-theme-text-secondary mb-3">Estimated Loot</h3>
        <div class="space-y-2">
            {#each sortedLoot as [itemName, quantity]}
                <div class="flex justify-between items-center py-1">
                    <span class="text-sm text-theme-text-primary capitalize">{itemName}</span>
                    <span class="text-sm font-medium text-theme-accent">{formatNumber(quantity)}</span>
                </div>
            {/each}
        </div>
        
        {#if sortedLoot.length === 0}
            <p class="text-sm text-theme-text-secondary italic">No loot to display</p>
        {/if}
    </div>

    <!-- Additional Info -->
    <div class="mt-6 p-4 bg-blue-900/20 border border-blue-800 rounded-md">
        <h4 class="text-sm font-medium text-blue-200 mb-2">Note:</h4>
        <p class="text-xs text-blue-300">
            Loot quantities are simulated estimates and will vary with each calculation. 
            The Phoenix pet and rare drops are based on statistical probability.
        </p>
    </div>
</div>