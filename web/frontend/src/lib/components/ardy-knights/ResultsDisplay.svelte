<script lang="ts">
    import type { ArdyKnightResult } from '$lib/types';

    export let results: ArdyKnightResult;

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
</script>

<div class="bg-theme-card-bg p-6 rounded-lg shadow-card border border-theme-border">
    <h2 class="text-h2 text-theme-text-primary mb-6">Calculation Results</h2>
    
    <!-- Success Rate & Efficiency -->
    <div class="grid grid-cols-2 gap-4 mb-6">
        <div class="bg-theme-surface p-4 rounded-md">
            <h3 class="text-sm font-medium text-theme-text-secondary mb-2">Success Rate</h3>
            <p class="text-2xl font-bold text-theme-accent">{(results.calculated_success_rate * 100).toFixed(1)}%</p>
        </div>
        <div class="bg-theme-surface p-4 rounded-md">
            <h3 class="text-sm font-medium text-theme-text-secondary mb-2">XP Per Hour</h3>
            <p class="text-2xl font-bold text-theme-accent">{formatNumber(results.xp_hour)}</p>
        </div>
    </div>

    <!-- GP & Profit -->
    <div class="grid grid-cols-2 gap-4 mb-6">
        <div class="bg-theme-surface p-4 rounded-md">
            <h3 class="text-sm font-medium text-theme-text-secondary mb-2">GP Per Hour</h3>
            <p class="text-2xl font-bold text-green-400">{formatGP(results.gp_hour)}</p>
        </div>
        <div class="bg-theme-surface p-4 rounded-md">
            <h3 class="text-sm font-medium text-theme-text-secondary mb-2">Profit Per Hour</h3>
            <p class="text-2xl font-bold {results.profit_per_hour >= 0 ? 'text-green-400' : 'text-red-400'}">
                {formatGP(results.profit_per_hour)}
            </p>
        </div>
    </div>

    <!-- Goal Progress -->
    <div class="bg-theme-surface p-4 rounded-md mb-6">
        <h3 class="text-sm font-medium text-theme-text-secondary mb-3">Goal Progress</h3>
        <div class="space-y-2">
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">Current Level:</span>
                <span class="text-theme-accent font-medium">{results.current_thieving_level}</span>
            </div>
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">Target Level:</span>
                <span class="text-theme-accent font-medium">{results.target_thieving_level}</span>
            </div>
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">XP Remaining:</span>
                <span class="text-theme-accent font-medium">{formatNumber(results.xp_to_target)}</span>
            </div>
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">Time to Target:</span>
                <span class="text-theme-accent font-medium">{formatTime(results.hours_to_target)}</span>
            </div>
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">Pickpockets Needed:</span>
                <span class="text-theme-accent font-medium">{formatNumber(results.pickpockets_to_target)}</span>
            </div>
        </div>
    </div>

    <!-- Additional Stats -->
    <div class="bg-theme-surface p-4 rounded-md">
        <h3 class="text-sm font-medium text-theme-text-secondary mb-3">Efficiency Details</h3>
        <div class="space-y-2">
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">Effective XP Per Attempt:</span>
                <span class="text-theme-accent font-medium">{results.effective_xp_per_attempt.toFixed(2)}</span>
            </div>
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">Effective GP Per Attempt:</span>
                <span class="text-theme-accent font-medium">{results.effective_gp_per_attempt.toFixed(2)}</span>
            </div>
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">Damage Per Hour:</span>
                <span class="text-red-400 font-medium">{formatNumber(results.damage_per_hour)}</span>
            </div>
            <div class="flex justify-between text-sm">
                <span class="text-theme-text-primary">Food Needed Per Hour:</span>
                <span class="text-theme-accent font-medium">{Math.ceil(results.food_needed_per_hour)}</span>
            </div>
        </div>
    </div>
</div>