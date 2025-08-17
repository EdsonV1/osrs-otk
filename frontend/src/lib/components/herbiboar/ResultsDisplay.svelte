<script lang="ts">
    import type { HerbiboarApiResult } from '$lib/types';
    import Icon from '$lib/components/shared/Icon.svelte';
    import { METRIC_ICONS, SKILL_ICONS, PET_ICONS, TOOL_ICONS } from '$lib/stores/icons';

    export let result: HerbiboarApiResult;
    export let iconSrc: string = TOOL_ICONS.herbiboar;
    
    // Performance-optimized icon configuration
    export let hunterIconSrc: string = SKILL_ICONS.hunter;
    export let herbloreIconSrc: string = SKILL_ICONS.herblore; 
    export let petIconSrc: string = PET_ICONS.herbi;
    export let lootIconSrc: string = METRIC_ICONS.coins.src;
    export let timeIconSrc: string = METRIC_ICONS.clock.src;

    function formatNumber(num: number): string {
        return new Intl.NumberFormat().format(Math.round(num));
    }

    function formatTime(hours: number): string {
        if (hours < 1) {
            const minutes = Math.round(hours * 60);
            return `${minutes}m`;
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

    function formatPercent(num: number): string {
        return `${num.toFixed(2)}%`;
    }

    // Sort herbs by total value (quantity * estimated price) - performance optimized
    $: sortedHerbs = Object.entries(result.herbs_obtained)
        .filter(([_, quantity]) => quantity > 0)
        .sort(([a], [b]) => {
            // Herb tier ordering for consistent display
            const herbTiers: { [key: string]: number } = {
                'Grimy torstol': 14, 'Grimy dwarf weed': 13, 'Grimy lantadyme': 12,
                'Grimy cadantine': 11, 'Grimy snapdragon': 10, 'Grimy kwuarm': 9,
                'Grimy avantoe': 8, 'Grimy irit leaf': 7, 'Grimy toadflax': 6,
                'Grimy ranarr weed': 5, 'Grimy harralander': 4, 'Grimy tarromin': 3,
                'Grimy marrentill': 2, 'Grimy guam leaf': 1
            };
            return (herbTiers[b] || 0) - (herbTiers[a] || 0);
        });

    $: totalHerbs = Object.values(result.herbs_obtained).reduce((sum, qty) => sum + qty, 0);
    $: petChanceColor = result.pet_chance_percent > 50 ? 'text-green-400' : 
                       result.pet_chance_percent > 25 ? 'text-orange-400' : 'text-red-400';
    $: petChanceBg = result.pet_chance_percent > 50 ? 'bg-green-500/10' : 
                     result.pet_chance_percent > 25 ? 'bg-orange-500/10' : 'bg-red-500/10';
</script>

<!-- Nature-themed container with herb aesthetics -->
<div class="relative overflow-hidden bg-gradient-to-br from-green-900/20 via-theme-bg-secondary/90 to-emerald-900/20 backdrop-blur-lg border border-theme-border-primary/30 rounded-card shadow-card-hover animate-slide-up">
    <!-- Background herb/nature effect -->
    <div class="absolute inset-0 bg-gradient-to-br from-green-500/3 via-emerald-500/2 to-teal-500/3"></div>
    
    <!-- Header -->
    <div class="relative p-6 pb-4">
        <div class="flex items-center justify-between mb-6">
            <div class="flex items-center space-x-3">
                <div class="w-12 h-12 bg-gradient-to-br from-green-600 to-emerald-700 rounded-xl flex items-center justify-center shadow-glow animate-glow p-2">
                    <Icon src={iconSrc} alt="Herbiboar icon" size="lg" loading="eager" preload={true} />
                </div>
                <div>
                    <h2 class="text-2xl font-bold text-theme-text-primary">Herbiboar Results</h2>
                    <p class="text-theme-text-tertiary text-sm">Your hunter & herblore journey</p>
                </div>
            </div>
            <div class="text-right">
                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Herbiboars</div>
                <div class="text-xl font-bold text-green-400">{formatNumber(result.herbiboars_caught)}</div>
            </div>
        </div>
    </div>

    <!-- Strategy & Performance Info -->
    <div class="relative px-6 pb-4">
        <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mb-6">
            <!-- Rate Per Hour -->
            <div class="bg-theme-bg-tertiary/80 p-4 rounded-card border border-theme-border-accent/20">
                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Rate</div>
                <div class="text-lg font-bold text-theme-text-primary">{formatNumber(result.herbiboars_per_hour)}/hr</div>
                <div class="text-xs text-theme-text-secondary">Herbiboars caught</div>
            </div>
            
            <!-- Time Required -->
            <div class="bg-theme-bg-tertiary/80 p-4 rounded-card border border-theme-border-accent/20">
                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Time Needed</div>
                <div class="text-lg font-bold text-theme-accent-primary">{formatTime(result.time_required_hours)}</div>
                <div class="text-xs text-theme-text-secondary">Total hunting time</div>
            </div>
            
            <!-- Total Herbs -->
            <div class="bg-theme-bg-tertiary/80 p-4 rounded-card border border-theme-border-accent/20">
                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Herbs Gained</div>
                <div class="text-lg font-bold text-emerald-400">{formatNumber(totalHerbs)}</div>
                <div class="text-xs text-theme-text-secondary">{result.gear_effects.magic_secateurs.herbs_per_boar}/herbiboar</div>
            </div>
        </div>

        <!-- Key Metrics Dashboard -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
            <!-- Hunter XP Card -->
            <div class="group bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-orange-400/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5">
                <div class="flex items-center justify-between mb-3">
                    <div class="w-10 h-10 bg-gradient-to-br from-orange-500 to-red-600 rounded-lg flex items-center justify-center p-2">
                        <Icon src={hunterIconSrc} alt="Hunter skill icon" size="md" loading="eager" preload={true} />
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Hunter XP</div>
                        <div class="text-2xl font-bold text-orange-400">{formatNumber(result.hunter_xp)}</div>
                        <div class="text-xs text-theme-text-secondary">{formatNumber(result.hunter_xp / result.time_required_hours)}/hr</div>
                    </div>
                </div>
                <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                    <div class="bg-gradient-to-r from-orange-500 to-red-600 h-2 rounded-full animate-pulse" style="width: 75%"></div>
                </div>
            </div>

            <!-- Herblore XP Card -->
            <div class="group bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-green-400/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5">
                <div class="flex items-center justify-between mb-3">
                    <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center p-2">
                        <Icon src={herbloreIconSrc} alt="Herblore skill icon" size="md" loading="eager" preload={true} />
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Herblore XP</div>
                        <div class="text-2xl font-bold text-green-400">{formatNumber(result.herblore_xp)}</div>
                        <div class="text-xs text-theme-text-secondary">{formatNumber(result.herblore_xp / result.time_required_hours)}/hr</div>
                    </div>
                </div>
                <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                    <div class="bg-gradient-to-r from-green-500 to-emerald-600 h-2 rounded-full animate-pulse" style="width: 85%"></div>
                </div>
            </div>
        </div>

        <!-- Profit & Pet Section -->
        {#if result.total_profit_gp > 0}
            <div class="grid grid-cols-1 lg:grid-cols-3 gap-4 mb-6">
                <!-- Total Profit -->
                <div class="lg:col-span-1 bg-theme-bg-tertiary/80 p-5 rounded-card border border-theme-border-accent/20 hover:border-green-400/40 transition-colors duration-200">
                    <div class="flex items-center justify-between">
                        <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center p-2">
                            <Icon src={lootIconSrc} alt="Profit icon" size="md" loading="eager" preload={true} />
                        </div>
                        <div class="text-right">
                            <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Total Profit</div>
                            <div class="text-xl font-bold text-green-400">{formatGP(result.total_profit_gp)}</div>
                        </div>
                    </div>
                </div>

                <!-- Profit Per Hour -->
                <div class="lg:col-span-2 bg-theme-bg-tertiary/80 p-5 rounded-card border border-theme-border-accent/20">
                    <div class="flex items-center justify-between">
                        <div class="w-10 h-10 bg-gradient-to-br from-indigo-500 to-blue-600 rounded-lg flex items-center justify-center p-2">
                            <Icon src={timeIconSrc} alt="Profit rate icon" size="md" loading="eager" preload={true} />
                        </div>
                        <div class="text-right">
                            <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Profit Rate</div>
                            <div class="text-xl font-bold text-green-400">{formatGP(result.profit_per_hour_gp)}/hr</div>
                            <div class="text-xs text-theme-text-secondary">
                                {#if result.price_info}
                                    {result.price_info.source === 'live' ? 'Live prices' : 'Static estimates'}
                                {/if}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        {/if}

        <!-- Herbi Pet Card -->
        <div class="group bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-purple-400/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5 mb-6">
            <div class="flex items-center justify-between mb-3">
                <div class="w-10 h-10 bg-gradient-to-br from-purple-500 to-pink-600 rounded-lg flex items-center justify-center {petChanceBg} p-2">
                    <Icon src={petIconSrc} alt="Herbi pet icon" size="md" loading="lazy" />
                </div>
                <div class="text-right">
                    <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Herbi Pet Chance</div>
                    <div class="text-2xl font-bold {petChanceColor}">{formatPercent(result.pet_chance_percent)}</div>
                    <div class="text-xs text-theme-text-secondary">1/6,500 per catch</div>
                </div>
            </div>
            <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                <div class="bg-gradient-to-r from-purple-500 to-pink-600 h-2 rounded-full" style="width: {Math.min(result.pet_chance_percent, 100)}%"></div>
            </div>
        </div>

        <!-- Gear Effects -->
        {#if result.gear_effects.magic_secateurs.used}
            <div class="bg-theme-bg-tertiary/80 p-4 rounded-card border border-theme-border-accent/20 mb-6">
                <div class="flex items-center space-x-4">
                    <div class="w-3 h-3 bg-green-400 rounded-full animate-pulse"></div>
                    <div>
                        <p class="text-sm font-medium text-theme-text-primary">Magic Secateurs Active</p>
                        <p class="text-xs text-theme-text-secondary">
                            +{formatNumber(result.gear_effects.magic_secateurs.extra_herbs_gained)} extra herbs gained 
                            ({result.gear_effects.magic_secateurs.herbs_per_boar} herbs per herbiboar)
                        </p>
                    </div>
                </div>
            </div>
        {/if}

        <!-- Herb Drops Breakdown -->
        <div class="bg-theme-bg-tertiary/80 rounded-card border border-theme-border-accent/20 overflow-hidden">
            <div class="p-4 border-b border-theme-border-subtle/30">
                <div class="flex items-center justify-between">
                    <div class="flex items-center space-x-3">
                        <div class="w-8 h-8 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center">
                            <span class="text-white text-sm">🌿</span>
                        </div>
                        <h3 class="text-lg font-semibold text-theme-text-primary">Herb Drops</h3>
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary">Total Herbs</div>
                        <div class="text-sm font-bold text-green-400">{formatNumber(totalHerbs)}</div>
                    </div>
                </div>
            </div>
            
            <!-- OSRS-Style Herb Grid -->
            <div class="p-4">
                <div class="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-3">
                    {#each sortedHerbs as [herbName, quantity]}
                        <div class="group bg-theme-bg-tertiary/60 hover:bg-theme-bg-tertiary rounded-lg border border-theme-border-subtle/50 hover:border-theme-accent-primary/30 transition-colors duration-200 p-3">
                            <!-- Herb Icon -->
                            <div class="relative mb-2">
                                <div class="w-10 h-10 mx-auto bg-gradient-to-br from-green-400 to-emerald-600 rounded-lg flex items-center justify-center shadow-sm group-hover:shadow-md transition-shadow">
                                    <div class="w-6 h-6 bg-green-200/20 rounded-sm flex items-center justify-center">
                                        <div class="w-3 h-4 bg-green-600 rounded-sm"></div>
                                    </div>
                                </div>
                            </div>
                            
                            <!-- Herb Details -->
                            <div class="text-center">
                                <div class="text-xs font-medium text-theme-text-primary truncate" title={herbName}>
                                    {herbName.replace('Grimy ', '')}
                                </div>
                                <div class="text-xs text-theme-accent-primary font-bold">{formatNumber(quantity)}</div>
                            </div>
                        </div>
                    {/each}
                </div>
            </div>
        </div>
    </div>
</div>