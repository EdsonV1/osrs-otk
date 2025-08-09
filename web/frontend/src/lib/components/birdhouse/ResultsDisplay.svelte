<script lang="ts">
    import type { BirdhouseApiResult, SeedDropInfo } from '$lib/types'; 
    
    export let apiResult: BirdhouseApiResult | null;
    export let iconSrc: string = '/images/birdhouse/redwood_bird_house.png';
    
    // Metric icons configuration
    export let experienceIconSrc: string = '/images/icons/experience.png';
    export let lootIconSrc: string = '/images/icons/coins.png';
    export let timeIconSrc: string = '/images/icons/clock.png';

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

    $: totalValue = apiResult?.total_loot || 0;
    $: nestCount = apiResult?.estimated_nests || 0;
</script>

{#if apiResult}
    <!-- Hunter-themed container with nature aesthetics -->
    <div class="relative overflow-hidden bg-gradient-to-br from-green-900/20 via-theme-bg-secondary/90 to-emerald-900/20 backdrop-blur-lg border border-theme-border-primary/30 rounded-card shadow-card-hover animate-slide-up">
        <!-- Background nature effect -->
        <div class="absolute inset-0 bg-gradient-to-br from-green-500/3 via-emerald-500/2 to-teal-500/3"></div>
        
        <!-- Header -->
        <div class="relative p-6 pb-4">
            <div class="flex items-center justify-between mb-6">
                <div class="flex items-center space-x-3">
                    <div class="w-12 h-12 bg-gradient-to-br from-green-600 to-emerald-700 rounded-xl flex items-center justify-center shadow-glow p-2">
                        <img src={iconSrc} alt="Birdhouse icon" class="w-full h-full object-contain" />
                    </div>
                    <div>
                        <h2 class="text-2xl font-bold text-theme-text-primary">Birdhouse Results</h2>
                        <p class="text-theme-text-tertiary text-sm">Your hunter training breakdown</p>
                    </div>
                </div>
                <div class="text-right">
                    <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Nests</div>
                    <div class="text-xl font-bold text-green-400">{formatNum(nestCount, 1)}</div>
                </div>
            </div>
        </div>

        <!-- Key Metrics Dashboard -->
        <div class="relative px-6 pb-4">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
                <!-- Hunter XP Card -->
                <div class="group bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-green-400/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5">
                    <div class="flex items-center justify-between mb-3">
                        <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center p-2">
                            <img src={experienceIconSrc} alt="Experience icon" class="w-full h-full object-contain" />
                        </div>
                        <div class="text-right">
                            <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Hunter XP</div>
                            <div class="text-2xl font-bold text-theme-accent-primary">{formatNum(apiResult.hunter_xp)}</div>
                            <div class="text-xs text-theme-text-secondary">Experience gained</div>
                        </div>
                    </div>
                    <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                        <div class="bg-gradient-to-r from-green-500 to-emerald-600 h-2 rounded-full animate-pulse" style="width: 85%"></div>
                    </div>
                </div>

                <!-- Profit Card -->
                <div class="group bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-yellow-400/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5">
                    <div class="flex items-center justify-between mb-3">
                        <div class="w-10 h-10 bg-gradient-to-br from-yellow-500 to-orange-600 rounded-lg flex items-center justify-center p-2">
                            <img src={lootIconSrc} alt="Loot icon" class="w-full h-full object-contain" />
                        </div>
                        <div class="text-right">
                            <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Loot Value</div>
                            <div class="text-2xl font-bold text-yellow-400">{formatNum(totalValue)} GP</div>
                            <div class="text-xs text-theme-text-secondary">Seeds & rewards</div>
                        </div>
                    </div>
                    <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                        <div class="bg-gradient-to-r from-yellow-500 to-orange-600 h-2 rounded-full" style="width: {Math.min((totalValue / 100000) * 100, 100)}%"></div>
                    </div>
                </div>
            </div>

            <!-- Time Efficiency Section -->
            <div class="bg-glass backdrop-blur-md rounded-card border border-theme-border-accent/20 p-5 mb-6">
                <div class="flex items-center justify-between mb-4">
                    <div class="flex items-center space-x-3">
                        <div class="w-8 h-8 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-lg flex items-center justify-center p-1.5">
                            <img src={timeIconSrc} alt="Time icon" class="w-full h-full object-contain" />
                        </div>
                        <h3 class="text-lg font-semibold text-theme-text-primary">Time Investment</h3>
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary">Crafting XP Bonus</div>
                        <div class="text-lg font-bold text-blue-400">{formatNum(apiResult.crafting_xp)}</div>
                    </div>
                </div>
                
                <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                    {#each [
                        {label: 'Low Efficiency', value: `${apiResult.days_low_efficiency} days`, desc: '2 runs/day', color: 'text-red-400'},
                        {label: 'Medium Efficiency', value: `${apiResult.days_medium_efficiency} days`, desc: '7 runs/day', color: 'text-orange-400'},
                        {label: 'High Efficiency', value: `${apiResult.days_high_efficiency} days`, desc: '14 runs/day', color: 'text-green-400'}
                    ] as efficiency}
                        <div class="p-3 bg-theme-bg-tertiary/50 rounded-lg text-center">
                            <div class="text-xs text-theme-text-tertiary uppercase tracking-wider mb-1">{efficiency.label}</div>
                            <div class="text-lg font-bold {efficiency.color}">{efficiency.value}</div>
                            <div class="text-xs text-theme-text-secondary">{efficiency.desc}</div>
                        </div>
                    {/each}
                </div>
            </div>

            <!-- Seed Loot Breakdown -->
            {#if processedSeedDrops.length > 0}
                <div class="bg-glass backdrop-blur-md rounded-card border border-theme-border-accent/20 overflow-hidden mb-6">
                    <div class="p-4 border-b border-theme-border-subtle/30">
                        <div class="flex items-center justify-between">
                            <div class="flex items-center space-x-3">
                                <div class="w-8 h-8 bg-gradient-to-br from-emerald-500 to-teal-600 rounded-lg flex items-center justify-center">
                                    <span class="text-white text-sm">🌱</span>
                                </div>
                                <h3 class="text-lg font-semibold text-theme-text-primary">Seed Drops</h3>
                            </div>
                            <div class="text-right">
                                <div class="text-xs text-theme-text-tertiary">Total Value</div>
                                <div class="text-sm font-bold text-emerald-400">{formatNum(totalValue)} GP</div>
                            </div>
                        </div>
                    </div>
                    
                    <!-- OSRS-Style Seed Grid -->
                    <div class="p-4">
                        <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-3">
                            {#each processedSeedDrops as seed, index}
                                <div class="group relative bg-theme-bg-tertiary/60 hover:bg-theme-bg-tertiary rounded-lg border border-theme-border-subtle/50 hover:border-theme-accent-primary/30 transition-all duration-200 p-3 animate-scale-in" style="animation-delay: {index * 50}ms">
                                    <!-- Seed Image Placeholder -->
                                    <div class="relative mb-2">
                                        <div class="w-10 h-10 mx-auto bg-gradient-to-br from-green-400 to-emerald-600 rounded-lg flex items-center justify-center shadow-sm group-hover:shadow-md transition-shadow">
                                            <!-- Seed icon - you can replace with actual images -->
                                            <div class="w-6 h-6 bg-green-200/20 rounded-full flex items-center justify-center">
                                                <div class="w-3 h-3 bg-white/80 rounded-full"></div>
                                            </div>
                                        </div>
                                        
                                        <!-- Quantity badge -->
                                        <div class="absolute -top-1 -right-1 bg-theme-accent-primary text-white text-xs font-bold px-1.5 py-0.5 rounded-full shadow-sm min-w-6 text-center">
                                            {formatNum(seed.quantity)}
                                        </div>
                                    </div>
                                    
                                    <!-- Seed name -->
                                    <div class="text-center">
                                        <div class="text-xs font-medium text-theme-text-primary truncate mb-1" title={seed.name}>
                                            {seed.name}
                                        </div>
                                        <div class="text-xs text-emerald-400 font-semibold">
                                            {formatNum(seed.value)} GP
                                        </div>
                                    </div>
                                    
                                    <!-- Hover tooltip -->
                                    <div class="absolute -top-2 left-1/2 transform -translate-x-1/2 -translate-y-full opacity-0 group-hover:opacity-100 transition-opacity duration-200 pointer-events-none z-10">
                                        <div class="bg-theme-bg-elevated border border-theme-border-primary rounded-lg px-3 py-2 shadow-lg whitespace-nowrap">
                                            <div class="text-sm font-medium text-theme-text-primary">{seed.name} Seed</div>
                                            <div class="text-xs text-theme-text-secondary">Quantity: {formatNum(seed.quantity)}</div>
                                            <div class="text-xs text-emerald-400">Total Value: {formatNum(seed.value)} GP</div>
                                            <div class="text-xs text-theme-text-tertiary">Each: {formatNum(seed.value / seed.quantity)} GP</div>
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
                                    <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Seed Types</div>
                                    <div class="text-lg font-bold text-theme-accent-primary">{processedSeedDrops.length}</div>
                                </div>
                                <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                    <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Seeds</div>
                                    <div class="text-lg font-bold text-green-400">{formatNum(processedSeedDrops.reduce((sum, seed) => sum + seed.quantity, 0))}</div>
                                </div>
                                <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                    <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Avg Value</div>
                                    <div class="text-lg font-bold text-yellow-400">{formatNum(totalValue / processedSeedDrops.length)} GP</div>
                                </div>
                                <div class="p-2 bg-theme-bg-tertiary/30 rounded-lg">
                                    <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Value</div>
                                    <div class="text-lg font-bold text-emerald-400">{formatNum(totalValue)} GP</div>
                                </div>
                            </div>
                        </div>
                        
                        <!-- Value per nest info -->
                        <div class="mt-3 text-center">
                            <div class="text-xs text-theme-text-tertiary">
                                Average value per bird nest: <span class="text-theme-accent-primary font-medium">{formatNum(totalValue / nestCount)} GP</span>
                            </div>
                        </div>
                    </div>
                </div>
            {/if}

            <!-- Hunter Pro Tip -->
            <div class="relative bg-gradient-to-r from-green-600/10 via-emerald-600/5 to-transparent p-4 rounded-card border border-green-400/20">
                <div class="flex items-start space-x-3">
                    <div class="w-6 h-6 bg-green-600 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                        <span class="text-white text-xs">💡</span>
                    </div>
                    <div>
                        <h4 class="text-sm font-semibold text-green-400 mb-1">Hunter Pro Tip</h4>
                        <p class="text-xs text-theme-text-secondary leading-relaxed">
                            Use higher tier logs for better XP rates! Don't forget to check your birdhouses every 50 minutes for maximum efficiency. 
                            Bird nest contents are random - these are statistical estimates.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
{:else}
    <!-- Empty state with hunter theme -->
    <div class="relative bg-glass backdrop-blur-md rounded-card border border-theme-border-primary/30 p-12 text-center animate-fade-in">
        <div class="w-16 h-16 bg-gradient-to-br from-green-500 to-emerald-600 rounded-xl flex items-center justify-center mx-auto mb-4 opacity-50">
            <span class="text-3xl">🏠</span>
        </div>
        <h3 class="text-lg font-semibold text-theme-text-primary mb-2">Ready for Your Birdhouse Run?</h3>
        <p class="text-theme-text-secondary">Enter your details above to see your hunter training results and potential loot!</p>
    </div>
{/if}