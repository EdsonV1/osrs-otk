<script lang="ts">
    import type { GOTRApiResult } from '$lib/types';

    export let apiResult: GOTRApiResult;
    export let iconSrc: string = '/images/skills/runecraft.png';
    
    // Metric icons configuration
    export let timeIconSrc: string = '/images/icons/clock.png';
    export let experienceIconSrc: string = '/images/icons/experience.png';
    export let lootIconSrc: string = '/images/icons/coins.png';

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

    $: sortedRewards = apiResult.estimated_rewards
        .filter(reward => reward.quantity > 0)
        .sort((a, b) => (b.quantity * b.value) - (a.quantity * a.value));

    $: petChanceColor = apiResult.pet_chance_percentage > 50 ? 'text-green-400' : apiResult.pet_chance_percentage > 25 ? 'text-orange-400' : 'text-red-400';
    $: petChanceBg = apiResult.pet_chance_percentage > 50 ? 'bg-green-500/10' : apiResult.pet_chance_percentage > 25 ? 'bg-orange-500/10' : 'bg-red-500/10';
</script>

<!-- Main Results Container -->
<div class="relative bg-theme-bg-secondary border border-theme-border-primary/30 rounded-card shadow-card-hover">
    <!-- Mystic background effect -->
    <div class="absolute inset-0 bg-gradient-to-br from-purple-500/3 to-blue-500/2 rounded-card"></div>
    
    <!-- Header with GOTR icon -->
    <div class="relative p-6 pb-4">
        <div class="flex items-center justify-between mb-6">
            <div class="flex items-center space-x-3">
                <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-blue-600 rounded-xl flex items-center justify-center shadow-glow animate-glow p-2">
                    <img src={iconSrc} alt="GOTR icon" class="w-full h-full object-contain" />
                </div>
                <div>
                    <h2 class="text-2xl font-bold text-theme-text-primary">GOTR Training Plan</h2>
                    <p class="text-theme-text-tertiary text-sm">Level {apiResult.current_level} → {apiResult.target_level} Runecrafting</p>
                </div>
            </div>
            <div class="text-right">
                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">XP Needed</div>
                <div class="text-xl font-bold text-theme-accent-primary">{formatNumber(apiResult.xp_needed)}</div>
            </div>
        </div>
    </div>

    <!-- Key Metrics Dashboard -->
    <div class="relative px-6 pb-4">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
            <!-- Games & Time Card -->
            <div class="group bg-theme-bg-tertiary/80 p-5 rounded-card border border-theme-border-accent/20 hover:border-theme-accent-primary/40 transition-colors duration-200 hover:shadow-card">
                <div class="flex items-center justify-between mb-3">
                    <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center p-2">
                        <img src={timeIconSrc} alt="Time icon" class="w-full h-full object-contain" />
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Training Time</div>
                        <div class="text-2xl font-bold text-theme-accent-primary">{formatTime(apiResult.hours_needed)}</div>
                        <div class="text-xs text-theme-text-secondary">{formatNumber(apiResult.games_needed)} games</div>
                    </div>
                </div>
                <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                    <div class="bg-gradient-to-r from-blue-500 to-purple-600 h-2 rounded-full animate-pulse" style="width: 85%"></div>
                </div>
            </div>

            <!-- XP Rates Card -->
            <div class="group bg-theme-bg-tertiary/80 p-5 rounded-card border border-theme-border-accent/20 hover:border-green-400/40 transition-colors duration-200 hover:shadow-card">
                <div class="flex items-center justify-between mb-3">
                    <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center p-2">
                        <img src={experienceIconSrc} alt="Experience icon" class="w-full h-full object-contain" />
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">XP Rates</div>
                        <div class="text-2xl font-bold text-green-400">{formatNumber(apiResult.average_xp_per_hour)}/hr</div>
                        <div class="text-xs text-theme-text-secondary">{formatNumber(apiResult.average_xp_per_game)} per game</div>
                    </div>
                </div>
                <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                    <div class="bg-gradient-to-r from-green-500 to-emerald-600 h-2 rounded-full" style="width: 90%"></div>
                </div>
            </div>
        </div>

        <!-- Profit & Pet Section -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 mb-6">
            <!-- Profit Per Hour -->
            <div class="lg:col-span-1 bg-theme-bg-tertiary/80 p-5 rounded-card border border-theme-border-accent/20 hover:border-yellow-400/40 transition-colors duration-200">
                <div class="flex items-center justify-between">
                    <div class="w-10 h-10 bg-gradient-to-br from-yellow-500 to-orange-600 rounded-lg flex items-center justify-center p-2">
                        <img src={lootIconSrc} alt="Loot icon" class="w-full h-full object-contain" />
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">GP/Hour</div>
                        <div class="text-xl font-bold text-yellow-400">{formatGP(apiResult.gp_per_hour)}</div>
                    </div>
                </div>
            </div>

            <!-- Pet Chance -->
            <div class="lg:col-span-2 bg-theme-bg-tertiary/80 p-5 rounded-card border border-theme-border-accent/20 hover:border-purple-400/40 transition-colors duration-200">
                <div class="flex items-center justify-between">
                    <div class="w-10 h-10 bg-gradient-to-br from-purple-500 to-pink-600 rounded-lg flex items-center justify-center {petChanceBg}">
                        <span class="text-white text-lg">🐾</span>
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Abyssal Protector Chance</div>
                        <div class="text-xl font-bold {petChanceColor}">{apiResult.pet_chance_percentage.toFixed(2)}%</div>
                        <div class="text-xs text-theme-text-secondary">{formatNumber(apiResult.total_reward_rolls)} reward rolls</div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Rewards Breakdown -->
        {#if sortedRewards.length > 0}
            <div class="bg-theme-bg-tertiary/80 rounded-card border border-theme-border-accent/20 overflow-hidden mb-6">
                <div class="p-4 border-b border-theme-border-subtle/30">
                    <div class="flex items-center justify-between">
                        <div class="flex items-center space-x-3">
                            <div class="w-8 h-8 bg-gradient-to-br from-purple-500 to-blue-600 rounded-lg flex items-center justify-center">
                                <span class="text-white text-sm">🎁</span>
                            </div>
                            <h3 class="text-lg font-semibold text-theme-text-primary">Estimated Rewards</h3>
                        </div>
                        <div class="text-right">
                            <div class="text-xs text-theme-text-tertiary">Total Value</div>
                            <div class="text-sm font-bold text-green-400">{formatGP(apiResult.total_reward_value)}</div>
                        </div>
                    </div>
                </div>
                
                <!-- OSRS-Style Reward Grid -->
                <div class="p-4">
                    <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-3">
                        {#each sortedRewards as reward}
                            {@const totalValue = reward.quantity * reward.value}
                            <div class="group bg-theme-bg-tertiary/60 hover:bg-theme-bg-tertiary rounded-lg border border-theme-border-subtle/50 hover:border-theme-accent-primary/30 transition-colors duration-200 p-3">
                                <!-- Item Image Placeholder -->
                                <div class="relative mb-2">
                                    <div class="w-10 h-10 mx-auto bg-gradient-to-br from-purple-400 to-blue-600 rounded-lg flex items-center justify-center shadow-sm group-hover:shadow-md transition-shadow">
                                        <!-- Item icon based on reward type -->
                                        <div class="w-6 h-6 bg-purple-200/20 rounded-sm flex items-center justify-center">
                                            {#if reward.name.includes('rune') || reward.name.includes('essence')}
                                                <div class="w-3 h-3 bg-blue-400 rounded-sm"></div>
                                            {:else if reward.name.includes('catalyst')}
                                                <div class="w-3 h-3 bg-purple-400 rounded-full"></div>
                                            {:else if reward.name.includes('guardian')}
                                                <div class="w-3 h-3 bg-green-400 rounded-full"></div>
                                            {:else if reward.name.includes('lantern')}
                                                <div class="w-3 h-3 bg-yellow-400 rotate-45"></div>
                                            {:else}
                                                <div class="w-3 h-3 bg-white/80 rounded-sm"></div>
                                            {/if}
                                        </div>
                                    </div>
                                    
                                    <!-- Quantity badge -->
                                    <div class="absolute -top-1 -right-1 bg-theme-accent-primary text-white text-xs font-bold px-1.5 py-0.5 rounded-full shadow-sm min-w-6 text-center">
                                        {formatNumber(reward.quantity)}
                                    </div>
                                </div>
                                
                                <!-- Item name and value -->
                                <div class="text-center">
                                    <div class="text-xs font-medium text-theme-text-primary truncate mb-1 capitalize" title={reward.name}>
                                        {reward.name}
                                    </div>
                                    <div class="text-xs text-purple-400 font-semibold">
                                        {formatGP(totalValue)}
                                    </div>
                                </div>
                                
                                <!-- Hover tooltip -->
                                <div class="absolute -top-2 left-1/2 transform -translate-x-1/2 -translate-y-full opacity-0 group-hover:opacity-100 transition-opacity duration-200 pointer-events-none z-10">
                                    <div class="bg-theme-bg-elevated border border-theme-border-primary rounded-lg px-3 py-2 shadow-lg whitespace-nowrap">
                                        <div class="text-sm font-medium text-theme-text-primary capitalize">{reward.name}</div>
                                        <div class="text-xs text-theme-text-secondary">Quantity: {formatNumber(reward.quantity)}</div>
                                        <div class="text-xs text-purple-400">Total Value: {formatGP(totalValue)}</div>
                                        <div class="text-xs text-theme-text-tertiary">Each: {formatGP(reward.value)}</div>
                                        {#if reward.drop_rate}
                                            <div class="text-xs text-orange-400">Drop Rate: {reward.drop_rate}</div>
                                        {/if}
                                    </div>
                                    <!-- Tooltip arrow -->
                                    <div class="absolute top-full left-1/2 transform -translate-x-1/2 w-2 h-2 bg-theme-bg-elevated border-r border-b border-theme-border-primary rotate-45"></div>
                                </div>
                            </div>
                        {/each}
                    </div>
                    
                    <!-- Summary stats -->
                    <div class="mt-4 pt-4 border-t border-theme-border-subtle/30">
                        <div class="grid grid-cols-2 md:grid-cols-4 gap-4 text-center">
                            <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Reward Types</div>
                                <div class="text-lg font-bold text-theme-accent-primary">{sortedRewards.length}</div>
                            </div>
                            <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Reward Rolls</div>
                                <div class="text-lg font-bold text-purple-400">{formatNumber(apiResult.total_reward_rolls)}</div>
                            </div>
                            <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Profit</div>
                                <div class="text-lg font-bold text-green-400">{formatGP(apiResult.total_reward_value)}</div>
                            </div>
                            <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Per Game Avg</div>
                                <div class="text-lg font-bold text-yellow-400">{formatGP(apiResult.total_reward_value / apiResult.games_needed)}</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        {/if}

        <!-- Pro Tip -->
        <div class="mt-6 relative bg-gradient-to-r from-theme-accent-primary/10 via-theme-accent-primary/5 to-transparent p-4 rounded-card border border-theme-accent-primary/20">
            <div class="flex items-start space-x-3">
                <div class="w-6 h-6 bg-theme-accent-primary rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <span class="text-white text-xs">💡</span>
                </div>
                <div>
                    <h4 class="text-sm font-semibold text-theme-accent-primary mb-1">Pro Tip</h4>
                    <p class="text-xs text-theme-text-secondary leading-relaxed">
                        GOTR provides the best Runecrafting XP/hour at level 77+. Rewards scale with your Runecrafting level. 
                        Consider wearing the Runecrafting outfit for additional XP bonuses!
                    </p>
                </div>
            </div>
        </div>
    </div>
</div>