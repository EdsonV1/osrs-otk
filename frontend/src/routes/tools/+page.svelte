<script lang="ts">
    const availableTools = [
        {
            name: 'Ardougne Knight Calculator',
            description: 'Optimize your Thieving training with Ardougne Knights: XP, GP, and efficiency.',
            href: '/tools/ardy-knights',
            iconSrc: "/images/tools/knight_of_ardougne.png"
        },
        {
            name: 'Birdhouse Run Calculator',
            description: 'Plan your birdhouse runs for maximum Hunter XP and valuable bird nests.',
            href: '/tools/birdhouses',
            iconSrc: "/images/birdhouse/redwood_bird_house.png"
        },
        {
            name: 'Wintertodt Calculator',
            description: 'Calculate Firemaking XP, Phoenix pet chances, and loot from Wintertodt rounds.',
            href: '/tools/wintertodt',
            iconSrc: "/images/skills/firemaking.png"
        },
        {
            name: 'Guardians of the Rift Calculator',
            description: 'Calculate Runecrafting XP rates, training time, rewards, and Abyssal Protector pet chances.',
            href: '/tools/gotr',
            iconSrc: "/images/skills/runecraft.png"
        }
    ];
</script>

<div class="max-w-5xl mx-auto">
    <header class="mb-10 text-center">
        <h1 class="text-h1 text-theme-text-primary tracking-tight">
            OSRS Tools & Calculators
        </h1>
        <p class="mt-3 text-lg text-theme-text-secondary max-w-2xl mx-auto">
            A collection of utility tools and calculators to help with your Old School RuneScape journey.
        </p>
    </header>

    {#if availableTools.length > 0}
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-6">
            {#each availableTools as tool (tool.href)}
                <a href={tool.href}
                   class="group block p-6 bg-theme-card-bg rounded-xl shadow-card hover:shadow-inner-border transition-all duration-200 ease-in-out border border-theme-border hover:border-theme-accent focus:outline-none focus:ring-2 focus:ring-theme-accent focus:ring-offset-2 focus:ring-offset-theme-bg">
                    <div class="flex items-center space-x-4 mb-3">
                            {#if tool.iconSrc}
                                <img src={tool.iconSrc} alt="{tool.name} icon" 
                                class="w-12 h-12 object-contain flex-shrink-0 mt-1 rounded-md group-hover:scale-105 transition-transform" />
                            {/if}
                        <h2 class="text-lg font-semibold text-theme-text-primary group-hover:text-theme-accent transition-colors">
                            {tool.name}
                        </h2>
                    </div>
                    <p class="text-sm text-theme-text-secondary leading-relaxed">
                        {tool.description}
                    </p>
                </a>
            {/each}
        </div>
    {:else}
        <div class="text-center py-10">
            <p class="text-theme-text-secondary text-lg">More tools coming soon!</p>
        </div>
    {/if}
</div>