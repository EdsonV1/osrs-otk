<script lang="ts">
    import type { ArdyKnightResult } from '$lib/types';

    export let results: ArdyKnightResult;
    export let iconSrc: string = '/tools/knight_of_ardougne.png';
    
    // Metric icons configuration
    export let lootIconSrc: string = '/icons/coins.png';
    export let experienceIconSrc: string = '/icons/experience.png';

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

    $: successRateColor = results.calculated_success_rate >= 0.8 ? 'text-green-400' : results.calculated_success_rate >= 0.6 ? 'text-orange-400' : 'text-red-400';
    $: profitColor = results.profit_per_hour >= 0 ? 'text-green-400' : 'text-red-400';
    $: levelProgress = ((results.current_thieving_level - 1) / (results.target_thieving_level - 1)) * 100;
</script>

<!-- Thieving-themed container with dark aesthetics -->
<div class="relative overflow-hidden bg-gradient-to-br from-purple-900/20 via-theme-bg-secondary/90 to-indigo-900/20 backdrop-blur-lg border border-theme-border-primary/30 rounded-card shadow-card-hover animate-slide-up">
    <!-- Background shadow effect -->
    <div class="absolute inset-0 bg-gradient-to-br from-purple-500/3 via-indigo-500/2 to-gray-900/5"></div>
    
    <!-- Header -->
    <div class="relative p-6 pb-4">
        <div class="flex items-center justify-between mb-6">
            <div class="flex items-center space-x-3">
                <div class="w-12 h-12 bg-gradient-to-br from-purple-600 to-indigo-700 rounded-xl flex items-center justify-center shadow-glow p-2">
                    <img src={iconSrc} alt="Ardy Knight icon" class="w-full h-full object-contain" />
                </div>
                <div>
                    <h2 class="text-2xl font-bold text-theme-text-primary">Thieving Results</h2>
                    <p class="text-theme-text-tertiary text-sm">Master the art of pickpocketing</p>
                </div>
            </div>
            <div class="text-right">
                <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Success Rate</div>
                <div class="text-xl font-bold {successRateColor}">{(results.calculated_success_rate * 100).toFixed(1)}%</div>
            </div>
        </div>
    </div>

    <!-- Key Performance Metrics -->
    <div class="relative px-6 pb-4">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
            <!-- XP & Efficiency Card -->
            <div class="group bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-purple-400/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5">
                <div class="flex items-center justify-between mb-3">
                    <div class="w-10 h-10 bg-gradient-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
                        <span class="text-white text-lg">⚡</span>
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Experience Rate</div>
                        <div class="text-2xl font-bold text-theme-accent-primary">{formatNumber(results.xp_hour)}</div>
                        <div class="text-xs text-theme-text-secondary">XP per hour</div>
                    </div>
                </div>
                <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                    <div class="bg-gradient-to-r from-blue-500 to-purple-600 h-2 rounded-full animate-pulse" style="width: {Math.min(results.xp_hour / 1000, 100)}%"></div>
                </div>
            </div>

            <!-- Profit Card -->
            <div class="group bg-glass backdrop-blur-md p-5 rounded-card border border-theme-border-accent/20 hover:border-green-400/40 transition-all duration-300 hover:shadow-card transform hover:-translate-y-0.5">
                <div class="flex items-center justify-between mb-3">
                    <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-emerald-600 rounded-lg flex items-center justify-center p-2">
                        <img src={lootIconSrc} alt="Profit icon" class="w-full h-full object-contain" />
                    </div>
                    <div class="text-right">
                        <div class="text-xs text-theme-text-tertiary uppercase tracking-wider">Profit Rate</div>
                        <div class="text-2xl font-bold {profitColor}">{formatGP(results.profit_per_hour)}</div>
                        <div class="text-xs text-theme-text-secondary">After food costs</div>
                    </div>
                </div>
                <div class="w-full bg-theme-bg-tertiary rounded-full h-2 overflow-hidden">
                    <div class="bg-gradient-to-r from-green-500 to-emerald-600 h-2 rounded-full" style="width: {Math.max(0, Math.min((results.profit_per_hour / results.gp_hour) * 100, 100))}%"></div>
                </div>
            </div>
        </div>

        <!-- Level Progress Section -->
        <div class="bg-glass backdrop-blur-md rounded-card border border-theme-border-accent/20 p-5 mb-6">
            <div class="flex items-center justify-between mb-4">
                <div class="flex items-center space-x-3">
                    <div class="w-8 h-8 bg-gradient-to-br from-yellow-500 to-orange-600 rounded-lg flex items-center justify-center p-1.5">
                        <img src={experienceIconSrc} alt="Experience icon" class="w-full h-full object-contain" />
                    </div>
                    <h3 class="text-lg font-semibold text-theme-text-primary">Level Progress</h3>
                </div>
                <div class="text-right">
                    <div class="text-xs text-theme-text-tertiary">Level</div>
                    <div class="text-lg font-bold text-theme-accent-primary">{results.current_thieving_level} → {results.target_thieving_level}</div>
                </div>
            </div>
            
            <!-- Progress Bar -->
            <div class="w-full bg-theme-bg-tertiary rounded-full h-3 mb-4 overflow-hidden">
                <div class="bg-gradient-to-r from-yellow-500 to-orange-600 h-3 rounded-full transition-all duration-1000 ease-out" 
                     style="width: {levelProgress}%"></div>
            </div>
            
            <div class="grid grid-cols-2 gap-4 text-sm">
                <div class="space-y-2">
                    <div class="flex justify-between">
                        <span class="text-theme-text-secondary">XP Remaining:</span>
                        <span class="font-medium text-theme-text-primary">{formatNumber(results.xp_to_target)}</span>
                    </div>
                    <div class="flex justify-between">
                        <span class="text-theme-text-secondary">Time to Goal:</span>
                        <span class="font-medium text-theme-accent-primary">{formatTime(results.hours_to_target)}</span>
                    </div>
                </div>
                <div class="space-y-2">
                    <div class="flex justify-between">
                        <span class="text-theme-text-secondary">Pickpockets:</span>
                        <span class="font-medium text-theme-text-primary">{formatNumber(results.pickpockets_to_target)}</span>
                    </div>
                    <div class="flex justify-between">
                        <span class="text-theme-text-secondary">GP per Hour:</span>
                        <span class="font-medium text-green-400">{formatGP(results.gp_hour)}</span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Efficiency Breakdown -->
        <div class="bg-glass backdrop-blur-md rounded-card border border-theme-border-accent/20 overflow-hidden mb-6">
            <div class="p-4 border-b border-theme-border-subtle/30">
                <div class="flex items-center space-x-3">
                    <div class="w-8 h-8 bg-gradient-to-br from-indigo-500 to-purple-600 rounded-lg flex items-center justify-center">
                        <span class="text-white text-sm">⚙️</span>
                    </div>
                    <h3 class="text-lg font-semibold text-theme-text-primary">Efficiency Details</h3>
                </div>
            </div>
            <div class="p-4">
                <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div class="space-y-3">
                        <div class="flex justify-between items-center p-3 bg-theme-bg-tertiary/50 rounded-lg">
                            <span class="text-sm text-theme-text-secondary">XP per Attempt:</span>
                            <span class="font-bold text-theme-accent-primary">{results.effective_xp_per_attempt.toFixed(2)}</span>
                        </div>
                        <div class="flex justify-between items-center p-3 bg-theme-bg-tertiary/50 rounded-lg">
                            <span class="text-sm text-theme-text-secondary">GP per Attempt:</span>
                            <span class="font-bold text-green-400">{results.effective_gp_per_attempt.toFixed(2)}</span>
                        </div>
                    </div>
                    <div class="space-y-3">
                        <div class="flex justify-between items-center p-3 bg-theme-bg-tertiary/50 rounded-lg">
                            <span class="text-sm text-theme-text-secondary">Damage/Hour:</span>
                            <span class="font-bold text-red-400">{formatNumber(results.damage_per_hour)}</span>
                        </div>
                        <div class="flex justify-between items-center p-3 bg-theme-bg-tertiary/50 rounded-lg">
                            <span class="text-sm text-theme-text-secondary">Food/Hour:</span>
                            <span class="font-bold text-orange-400">{Math.ceil(results.food_needed_per_hour)}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Thieving Tip -->
        <div class="relative bg-gradient-to-r from-purple-600/10 via-indigo-600/5 to-transparent p-4 rounded-card border border-purple-400/20">
            <div class="flex items-start space-x-3">
                <div class="w-6 h-6 bg-purple-600 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <span class="text-white text-xs">💡</span>
                </div>
                <div>
                    <h4 class="text-sm font-semibold text-purple-400 mb-1">Thieving Pro Tip</h4>
                    <p class="text-xs text-theme-text-secondary leading-relaxed">
                        Higher thieving levels mean better success rates and less damage taken. Consider using Shadow Veil spell or Ardougne diary perks to boost your efficiency!
                    </p>
                </div>
            </div>
        </div>
    </div>
</div>