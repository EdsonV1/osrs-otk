<script lang="ts">
    import type { ArdyKnightResult } from '$lib/types';
    export let results: ArdyKnightResult | null;

    function formatNumber(num: number | null | undefined): string {
        if (num === null || num === undefined || typeof num !== 'number' || isNaN(num)) return 'N/A';
        return num.toLocaleString();
    }
    function formatDecimal(num: number | null | undefined, places: number = 2): string {
        if (num === null || num === undefined || typeof num !== 'number' || isNaN(num)) return 'N/A';
        return num.toFixed(places);
    }
    function formatHours(hours: number | null | undefined): string {
        if (hours === null || hours === undefined || !isFinite(hours) || typeof hours !== 'number' || isNaN(hours)) return 'N/A';
        if (hours === 0 && !(hours > 0) ) return '0h 0m';
        const totalMinutes = Math.round(hours * 60);
        if (totalMinutes === 0 && hours > 0) return '<1m';
        const h = Math.floor(totalMinutes / 60);
        const m = totalMinutes % 60;
        return `${h}h ${m}m`;
    }
</script>

<div class="bg-theme-card-bg shadow-card rounded-lg border border-theme-border p-6 space-y-6">
    <h2 class="text-h3 text-theme-text-primary border-b border-theme-border pb-4">Results</h2>

    {#if results}
        <div class="space-y-6">
            <section aria-labelledby="general-results-title">
                <dl class="space-y-2.5">
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">Success Rate:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatDecimal(results.calculated_success_rate * 100) }%</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">XP/Attempt:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatDecimal(results.effective_xp_per_attempt) }</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">GP/Attempt:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatDecimal(results.effective_gp_per_attempt) }</dd>
                    </div>
                </dl>
            </section>

            <hr class="border-theme-border opacity-30">

            <section aria-labelledby="hourly-rates-title">
                <h3 id="hourly-rates-title" class="text-base font-semibold leading-6 text-theme-text-primary mb-3">Hourly Rates</h3>
                <dl class="space-y-2.5">
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">XP/Hour:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatNumber(results.xp_hour) }</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">GP/Hour:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatNumber(results.gp_hour) }</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">Profit/Hour:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatNumber(results.profit_per_hour) } GP</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">Damage/Hour:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatNumber(results.damage_per_hour) }</dd>
                    </div>
                     <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">Food/Hour:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatNumber(results.food_needed_per_hour) } units</dd>
                    </div>
                </dl>
            </section>

            <hr class="border-theme-border opacity-30">

            <section aria-labelledby="progression-target-title">
                <h3 id="progression-target-title" class="text-base font-semibold leading-6 text-theme-text-primary mb-3">Progression to Target</h3>
                <dl class="space-y-2.5">
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">Current:</dt>
                        <dd class="font-semibold text-theme-text-primary">Lvl { results.current_thieving_level } ({formatNumber(results.current_total_xp)} XP)</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">Target:</dt>
                        <dd class="font-semibold text-theme-text-primary">Lvl { results.target_thieving_level } ({formatNumber(results.target_total_xp)} XP)</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">XP To Target:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatNumber(results.xp_to_target) }</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">Pickpockets:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatNumber(results.pickpockets_to_target) }</dd>
                    </div>
                    <div class="flex justify-between text-sm">
                        <dt class="text-theme-text-secondary">Time To Target:</dt>
                        <dd class="font-semibold text-theme-text-primary">{ formatHours(results.hours_to_target) }</dd>
                    </div>
                </dl>
            </section>
        </div>
    {:else}
        <p class="text-center text-theme-text-secondary py-6">No results to display yet.</p>
    {/if}
</div>
