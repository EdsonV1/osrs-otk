<script lang="ts">
    const featuredTools = [
        {
            name: 'Wintertodt Calculator',
            description: 'Optimize your Firemaking training and rewards from the Wintertodt boss.',
            href: '/tools/wintertodt',
            iconSrc: '/images/skills/firemaking.png',
            gradient: 'from-orange-500 to-red-600',
            stats: ['XP/Hour', 'Loot Value', 'Pet Chance']
        },
        {
            name: 'Ardougne Knights',
            description: 'Master Thieving with optimized Ardougne Knight calculations.',
            href: '/tools/ardy-knights',
            iconSrc: '/images/tools/knight_of_ardougne.png',
            gradient: 'from-purple-500 to-indigo-600',
            stats: ['GP/Hour', 'XP Rates', 'Success Rate']
        },
        {
            name: 'Birdhouse Runs',
            description: 'Plan efficient Hunter training with birdhouse run optimization.',
            href: '/tools/birdhouses',
            iconSrc: '/images/birdhouse/redwood_bird_house.png',
            gradient: 'from-green-500 to-teal-600',
            stats: ['XP/Run', 'Profit', 'Time Efficiency']
        },
        {
            name: 'Guardians of the Rift',
            description: 'Calculate Runecrafting XP rates, training time, and valuable rewards.',
            href: '/tools/gotr',
            iconSrc: '/images/skills/runecraft.png',
            gradient: 'from-purple-500 to-blue-600',
            stats: ['XP/Hour', 'Rewards', 'Pet Chance']
        }
    ];

    const skills = [
        { name: 'Attack', icon: '⚔️', href: '/skills/attack' },
        { name: 'Hunter', icon: '🏹', href: '/skills/hunter' },
        { name: 'Firemaking', icon: '🔥', href: '/skills/firemaking' },
        { name: 'Thieving', icon: '💰', href: '/skills/thieving' },
        { name: 'Agility', icon: '🏃', href: '/skills/agility' },
        { name: 'Mining', icon: '⛏️', href: '/skills/mining' }
    ];

    const stats = [
        { label: 'Calculators', value: '10+' },
        { label: 'Skills Covered', value: '23' },
        { label: 'Active Users', value: '2.5K+' }
    ];
</script>

<!-- Hero Section -->
<div class="relative bg-gradient-to-br from-theme-bg-primary via-theme-bg-secondary to-theme-bg-tertiary">
    <!-- Simplified background -->
    <div class="absolute inset-0 bg-hero-gradient opacity-60"></div>
    
    <div class="relative max-w-7xl mx-auto py-24 px-4 sm:px-6 lg:px-8">
        <div class="text-center">
            <!-- Main Heading -->
            <h1 class="text-5xl sm:text-6xl lg:text-7xl font-bold text-theme-text-primary mb-6 tracking-tight">
                <span class="block">OSRS</span>
                <span class="block bg-accent-gradient bg-clip-text text-transparent">
                    XP Kits & Tools!
                </span>
            </h1>
            
            <!-- Subtitle -->
            <p class="max-w-3xl mx-auto text-xl sm:text-2xl text-theme-text-secondary mb-12 leading-relaxed">
                The ultimate toolkit for Old School RuneScape players. Optimize your training, maximize your efficiency, and level up faster.
            </p>
            
            <!-- CTA Buttons -->
            <div class="flex flex-col sm:flex-row gap-4 justify-center items-center">
                <a href="/tools"
                   class="group px-8 py-4 bg-theme-accent hover:bg-theme-accent-hover text-white font-semibold rounded-button shadow-button hover:shadow-button-hover transition-colors duration-200">
                    <span class="flex items-center">
                        Explore Tools
                        <svg class="w-5 h-5 ml-2 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"></path>
                        </svg>
                    </span>
                </a>
                
                <a href="/skills"
                   class="px-8 py-4 bg-theme-bg-secondary hover:bg-theme-bg-tertiary border border-theme-border-primary hover:border-theme-accent-primary text-theme-text-primary font-semibold rounded-button transition-colors duration-200">
                    View Skills
                </a>
            </div>
        </div>
        
        <!-- Stats -->
        <div class="mt-20 grid grid-cols-1 sm:grid-cols-3 gap-8">
            {#each stats as stat}
                <div class="text-center">
                    <div class="text-3xl sm:text-4xl font-bold text-theme-accent-primary mb-2">{stat.value}</div>
                    <div class="text-theme-text-tertiary uppercase tracking-wider text-sm font-medium">{stat.label}</div>
                </div>
            {/each}
        </div>
    </div>
</div>

<!-- Featured Tools Section -->
<div class="py-24 px-4 sm:px-6 lg:px-8 bg-theme-bg-primary">
    <div class="max-w-7xl mx-auto">
        <div class="text-center mb-16">
            <h2 class="text-4xl font-bold text-theme-text-primary mb-4">Featured Calculators</h2>
            <p class="text-xl text-theme-text-secondary max-w-2xl mx-auto">
                Powerful tools designed to help you optimize your Old School RuneScape experience
            </p>
        </div>
        
        <div class="grid lg:grid-cols-3 gap-8">
            {#each featuredTools as tool}
                <div class="group relative">
                    <!-- Simplified Card -->
                    <a href={tool.href} class="block bg-theme-bg-secondary p-8 rounded-card border border-theme-border-primary group-hover:border-theme-accent-primary/50 transition-colors duration-200 hover:shadow-card">
                        <!-- Icon -->
                        <div class="w-16 h-16 bg-gradient-to-r {tool.gradient} rounded-xl flex items-center justify-center mb-6 p-2">
                            <img src={tool.iconSrc} alt="{tool.name} icon" class="w-full h-full object-contain" />
                        </div>
                        
                        <!-- Content -->
                        <h3 class="text-2xl font-bold text-theme-text-primary mb-3 group-hover:text-theme-accent-primary transition-colors duration-200">
                            {tool.name}
                        </h3>
                        
                        <p class="text-theme-text-secondary mb-6 leading-relaxed">
                            {tool.description}
                        </p>
                        
                        <!-- Features -->
                        <div class="flex flex-wrap gap-2 mb-6">
                            {#each tool.stats as stat}
                                <span class="px-3 py-1.5 bg-theme-bg-tertiary text-theme-text-tertiary text-sm rounded-pill border border-theme-border-subtle">
                                    {stat}
                                </span>
                            {/each}
                        </div>
                        
                        <!-- Arrow -->
                        <div class="flex items-center text-theme-accent-primary font-medium">
                            Try it now
                            <svg class="w-5 h-5 ml-2 group-hover:translate-x-1 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8l4 4m0 0l-4 4m4-4H3"></path>
                            </svg>
                        </div>
                    </a>
                </div>
            {/each}
        </div>

        <div class="text-center mt-12">
            <a href="/tools" 
               class="inline-flex items-center px-8 py-3 bg-theme-bg-tertiary hover:bg-theme-bg-elevated border border-theme-border-primary hover:border-theme-accent-primary text-theme-text-primary hover:text-theme-accent-primary font-medium rounded-button shadow-button hover:shadow-button-hover transition-all duration-200">
                View All Tools
                <svg class="w-4 h-4 ml-2 group-hover:translate-x-0.5 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                </svg>
            </a>
        </div>
    </div>
</div>

<!-- Skills Preview Section -->
<div class="py-24 px-4 sm:px-6 lg:px-8 bg-theme-bg-secondary">
    <div class="max-w-7xl mx-auto">
        <div class="text-center mb-16">
            <h2 class="text-4xl font-bold text-theme-text-primary mb-4">Skill Calculators</h2>
            <p class="text-xl text-theme-text-secondary">
                Comprehensive training guides and XP calculators for every skill
            </p>
        </div>
        
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
            {#each skills as skill}
                <a href={skill.href} 
                   class="group bg-theme-bg-tertiary hover:bg-theme-bg-tertiary p-6 rounded-card border border-theme-border-primary hover:border-theme-accent-primary/50 transition-colors duration-200 hover:shadow-card">
                    <div class="text-center">
                        <div class="text-3xl mb-3">
                            {skill.icon}
                        </div>
                        <div class="text-sm font-medium text-theme-text-primary group-hover:text-theme-accent-primary transition-colors duration-200">
                            {skill.name}
                        </div>
                    </div>
                </a>
            {/each}
        </div>
        
        <div class="text-center mt-12">
            <a href="/skills" 
               class="inline-flex items-center px-8 py-3 bg-theme-bg-tertiary hover:bg-theme-bg-elevated border border-theme-border-primary hover:border-theme-accent-primary text-theme-text-primary hover:text-theme-accent-primary font-medium rounded-button shadow-button hover:shadow-button-hover transition-all duration-200">
                View All Skills
                <svg class="w-4 h-4 ml-2 group-hover:translate-x-0.5 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                </svg>
            </a>
        </div>
    </div>
</div>