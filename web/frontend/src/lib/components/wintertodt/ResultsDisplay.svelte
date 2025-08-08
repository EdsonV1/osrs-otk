<script lang="ts">
    import type { WintertodtApiResult } from '$lib/types';

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

    $: petChanceColor = apiResult.pet_chance > 50 ? 'text-green-400' : apiResult.pet_chance > 25 ? 'text-orange-400' : 'text-red-400';
    $: petChanceBg = apiResult.pet_chance > 50 ? 'bg-green-500/10' : apiResult.pet_chance > 25 ? 'bg-orange-500/10' : 'bg-red-500/10';
</script>

<!-- Animated container with glassmorphism -->
<div class="relative overflow-hidden bg-gradient-to-br from-theme-bg-secondary/80 via-theme-bg-tertiary/60 to-theme-bg-secondary/80 backdrop-blur-lg border border-theme-border-primary/30 rounded-card shadow-card-hover animate-slide-up">
    <!-- Background fire effect -->
    <div class="absolute inset-0 bg-gradient-to-br from-orange-500/5 via-red-500/3 to-yellow-500/5 animate-pulse"></div>
    
    <!-- Header with fire icon -->
    <div class="relative p-6 pb-4">
        <div class="flex items-center justify-between mb-6">
            <div class="flex items-center space-x-3">
                <div class="w-12 h-12 bg-gradient-to-br from-orange-500 to-red-600 rounded-xl flex items-center justify-center shadow-glow animate-glow">
                    <span class="text-2xl">🔥</span>
                </div>
                <div>
                    <h2 class="text-2xl font-bold text-theme-text-primary">Wintertodt Results</h2>
                    <p class="text-theme-text-tertiary text-sm">Your firemaking journey awaits</p>
                </div>
            </div>
            <div class="text-right">
                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Rounds</div>
                <div class="text-xl font-bold text-theme-accent-primary">{apiResult.total_value || 'N/A'}</div>
            </div>
        </div>
    </div>

    <!-- Key Metrics Dashboard -->
    <div class="relative px-6 pb-4">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
            <!-- XP Card -->
            <div class="group relative bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-theme-accent-primary/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5">
                <div class="flex items-center justify-between mb-3">
                    <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
                        <span class="text-white text-lg">📊</span>
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Experience</div>
                        <div class="text-2xl font-bold text-theme-accent-primary">{formatNumber(apiResult.total_experience)}</div>
                        <div class="text-xs text-theme-text-secondary">{formatNumber(apiResult.average_exp_hour)}/hr</div>
                    </div>
                </div>
                <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                    <div class="bg-gradient-to-r from-blue-500 to-purple-600 h-2 rounded-full animate-pulse" style="width: 75%"></div>
                </div>
            </div>

            <!-- Phoenix Pet Card -->
            <div class="group relative bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-orange-400/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5">
                <div class="flex items-center justify-between mb-3">
                    <div class="w-10 h-10 bg-gradient-to-br from-orange-500 to-red-600 rounded-lg flex items-center justify-center {petChanceBg}">
                        <span class="text-white text-lg">🔥</span>
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Phoenix Chance</div>
                        <div class="text-2xl font-bold {petChanceColor}">{apiResult.pet_chance.toFixed(2)}%</div>
                        <div class="text-xs text-theme-text-secondary">Good luck!</div>
                    </div>
                </div>
                <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                    <div class="bg-gradient-to-r from-orange-500 to-red-600 h-2 rounded-full" style="width: {Math.min(apiResult.pet_chance, 100)}%"></div>
                </div>
            </div>
        </div>

        <!-- Loot & Time Section -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 mb-6">
            <!-- Loot Value -->
            <div class="lg:col-span-1 bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-green-400/40 transition-all duration-300">
                <div class="flex items-center justify-between">
                    <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center">
                        <span class="text-white text-lg">💰</span>
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Loot Value</div>
                        <div class="text-xl font-bold text-green-400">{formatGP(apiResult.total_value)}</div>
                    </div>
                </div>
            </div>

            <!-- Time Investment -->
            <div class="lg:col-span-2 bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20">
                <div class="flex items-center justify-between">
                    <div class="w-10 h-10 bg-gradient-to-br from-indigo-500 to-blue-600 rounded-lg flex items-center justify-center">
                        <span class="text-white text-lg">⏰</span>
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Time Investment</div>
                        <div class="text-xl font-bold text-theme-text-primary">{formatTime(apiResult.total_time)}</div>
                        <div class="text-xs text-theme-text-secondary">Total grinding time</div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Loot Breakdown -->
        {#if sortedLoot.length > 0}
            <div class="bg-glass backdrop-blur-md rounded-card border border-theme-border-accent/20 overflow-hidden mb-6">
                <div class="p-4 border-b border-theme-border-subtle/30">
                    <div class="flex items-center justify-between">
                        <div class="flex items-center space-x-3">
                            <div class="w-8 h-8 bg-gradient-to-br from-yellow-500 to-orange-600 rounded-lg flex items-center justify-center">
                                <span class="text-white text-sm">📦</span>
                            </div>
                            <h3 class="text-lg font-semibold text-theme-text-primary">Wintertodt Loot</h3>
                        </div>
                        <div class="text-right">
                            <div class="text-xs text-theme-text-tertiary">Total Value</div>
                            <div class="text-sm font-bold text-green-400">{formatGP(apiResult.total_value)}</div>
                        </div>
                    </div>
                </div>
                
                <!-- OSRS-Style Loot Grid -->
                <div class="p-4">
                    <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-3">
                        {#each sortedLoot as [itemName, quantity], index}
                            {@const itemValue = (apiResult.total_value / sortedLoot.reduce((sum, [_, qty]) => sum + qty, 0)) * quantity}
                            <div class="group relative bg-theme-bg-tertiary/60 hover:bg-theme-bg-tertiary rounded-lg border border-theme-border-subtle/50 hover:border-theme-accent-primary/30 transition-all duration-200 p-3 animate-scale-in" style="animation-delay: {index * 50}ms">
                                <!-- Item Image Placeholder -->
                                <div class="relative mb-2">
                                    <div class="w-10 h-10 mx-auto bg-gradient-to-br from-orange-400 to-red-600 rounded-lg flex items-center justify-center shadow-sm group-hover:shadow-md transition-shadow">
                                        <!-- Item icon - different colors for different item types -->
                                        <div class="w-6 h-6 bg-orange-200/20 rounded-sm flex items-center justify-center">
                                            {#if itemName.includes('log') || itemName.includes('wood')}
                                                <div class="w-3 h-4 bg-amber-600 rounded-sm"></div>
                                            {:else if itemName.includes('ore') || itemName.includes('bar')}
                                                <div class="w-3 h-3 bg-gray-400 rounded-full"></div>
                                            {:else if itemName.includes('seed')}
                                                <div class="w-3 h-3 bg-green-400 rounded-full"></div>
                                            {:else if itemName.includes('gem') || itemName.includes('uncut')}
                                                <div class="w-3 h-3 bg-purple-400 rotate-45"></div>
                                            {:else}
                                                <div class="w-3 h-3 bg-white/80 rounded-sm"></div>
                                            {/if}
                                        </div>
                                    </div>
                                    
                                    <!-- Quantity badge -->
                                    <div class="absolute -top-1 -right-1 bg-theme-accent-primary text-white text-xs font-bold px-1.5 py-0.5 rounded-full shadow-sm min-w-6 text-center">
                                        {formatNumber(quantity)}
                                    </div>
                                </div>
                                
                                <!-- Item name -->
                                <div class="text-center">
                                    <div class="text-xs font-medium text-theme-text-primary truncate mb-1 capitalize" title={itemName}>
                                        {itemName}
                                    </div>
                                    <div class="text-xs text-orange-400 font-semibold">
                                        {formatGP(itemValue)}
                                    </div>
                                </div>
                                
                                <!-- Hover tooltip -->
                                <div class="absolute -top-2 left-1/2 transform -translate-x-1/2 -translate-y-full opacity-0 group-hover:opacity-100 transition-opacity duration-200 pointer-events-none z-10">
                                    <div class="bg-theme-bg-elevated border border-theme-border-primary rounded-lg px-3 py-2 shadow-lg whitespace-nowrap">
                                        <div class="text-sm font-medium text-theme-text-primary capitalize">{itemName}</div>
                                        <div class="text-xs text-theme-text-secondary">Quantity: {formatNumber(quantity)}</div>
                                        <div class="text-xs text-orange-400">Estimated Value: {formatGP(itemValue)}</div>
                                        <div class="text-xs text-theme-text-tertiary">Each: ~{formatGP(itemValue / quantity)}</div>
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
                                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Item Types</div>
                                <div class="text-lg font-bold text-theme-accent-primary">{sortedLoot.length}</div>
                            </div>
                            <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Items</div>
                                <div class="text-lg font-bold text-orange-400">{formatNumber(sortedLoot.reduce((sum, [_, quantity]) => sum + quantity, 0))}</div>
                            </div>
                            <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Avg Value</div>
                                <div class="text-lg font-bold text-yellow-400">{formatGP(apiResult.total_value / sortedLoot.length)}</div>
                            </div>
                            <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Value</div>
                                <div class="text-lg font-bold text-green-400">{formatGP(apiResult.total_value)}</div>
                            </div>
                        </div>
                    </div>
                    
                    <!-- Value per round info -->
                    {#if apiResult.average_exp_hour && apiResult.average_exp_hour > 0}
                        <div class="mt-3 text-center">
                            <div class="text-xs text-theme-text-tertiary">
                                Average loot value per round: <span class="text-theme-accent-primary font-medium">{formatGP(apiResult.total_value / apiResult.average_exp_hour)}</span>
                            </div>
                        </div>
                    {/if}
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
                        Results are simulated estimates based on current game mechanics. Phoenix pet and rare drops follow statistical probability. 
                        Actual results may vary due to RNG and game updates.
                    </p>
                </div>
            </div>
        </div>
    </div>
</div>