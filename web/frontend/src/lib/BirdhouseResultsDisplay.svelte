<script lang="ts">
    import type { BirdhouseApiResult, SeedDropInfo } from '$lib/types'; 
    
    export let apiResult: BirdhouseApiResult | null;

    function formatNum(num: number | null | undefined, decimals: number = 0): string {
        if (num === null || num === undefined || typeof num !== 'number' || isNaN(num)) return 'N/A';
        return num.toLocaleString(undefined, { 
            minimumFractionDigits: decimals, 
            maximumFractionDigits: decimals 
        });
    }

    $: processedSeedDrops = apiResult?.seed_drops 
        ? Object.entries(apiResult.seed_drops)
            .map(([name, details]) => ({ 
                name: name.replace(/_/g, ' ').replace(/\b\w/g, l => l.toUpperCase()),
                quantity: details.quantity, 
                value: details.value 
            }))
            .filter(seed => seed.quantity > 0)
            .sort((a,b) => b.value - a.value || b.quantity - a.quantity)
        : [];
</script>

<div class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-6">
    <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Birdhouse Run Results</h2>

    {#if apiResult}
        <div class="space-y-6">
            <section>
                <h3 class="text-base font-semibold leading-6 text-theme-text-primary mb-3">
                    Calculated Summary 
                </h3>
                <dl class="space-y-2.5 text-sm">
                    {#each [
                        {label: 'Est. Nests (total):', value: formatNum(apiResult.estimated_nests, 1)},
                        {label: 'Total Seed Loot Value (GP):', value: formatNum(apiResult.total_loot)},
                        {label: 'Total Hunter XP:', value: formatNum(apiResult.hunter_xp)},
                        {label: 'Total Crafting XP:', value: formatNum(apiResult.crafting_xp)}
                    ] as item}
                        <div class="flex justify-between">
                            <dt class="text-theme-text-secondary">{item.label}</dt>
                            <dd class="font-semibold text-theme-text-primary">{item.value}</dd>
                        </div>
                        {/each}
                    </dl>
                    <p class="text-xs text-theme-text-tertiary mt-1.5 ml-2">*If crafting the birdhouses.</p>
            </section>

            <hr class="border-theme-border opacity-30">

            <section>
                 <h3 class="text-base font-semibold leading-6 text-theme-text-primary mb-3">Time Estimation (Days)</h3>
                 <p class="text-xs text-theme-text-tertiary mb-2.5">Backend calculates days assuming 4 houses per "run cycle" for this time estimation.</p>
                 <dl class="space-y-2.5 text-sm">
                    {#each [
                        {label: 'Low Efficiency (2 run cycles/day):', value: `${apiResult.days_low_efficiency} days`},
                        {label: 'Medium Efficiency (7 run cycles/day):', value: `${apiResult.days_medium_efficiency} days`},
                        {label: 'High Efficiency (14 run cycles/day):', value: `${apiResult.days_high_efficiency} days`}
                    ] as item}
                        <div class="flex justify-between">
                            <dt class="text-theme-text-secondary">{item.label}</dt>
                            <dd class="font-semibold text-theme-text-primary">{item.value}</dd>
                        </div>
                    {/each}
                 </dl>
            </section>
            
            {#if processedSeedDrops.length > 0}
                <hr class="border-theme-border opacity-30">
                <section>
                    <h3 class="text-base font-semibold leading-6 text-theme-text-primary mb-3">Simulated Seed Drops (Total)</h3>
                    <div class="max-h-60 overflow-y-auto pr-2 text-sm space-y-1.5">
                        {#each processedSeedDrops as seed}
                            <div class="flex justify-between">
                                <span class="text-theme-text-secondary">{seed.name} Seed:</span>
                                <span class="font-medium text-theme-text-primary">
                                    {formatNum(seed.quantity)} <span class="text-xs text-theme-text-tertiary">(Value: {formatNum(seed.value)} GP)</span>
                                </span>
                            </div>
                        {/each}
                    </div>
                    <p class="text-xs text-theme-text-tertiary mt-2">Total value from above seeds: {formatNum(apiResult.total_loot)} GP</p>
                </section>
            {/if}
        </div>
    {:else}
        <p class="text-center text-theme-text-secondary py-6">No results to display yet.</p>
    {/if}
</div>