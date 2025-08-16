<script lang="ts">
    interface Skill {
        name: string;
        iconSrc: string; 
        href: string;
        description?: string;
    }

    const skills: Skill[] = [
        { name: 'Attack', iconSrc: '/images/skills/attack.png', href: '/skills/attack', description: 'Train your Attack level.' },
        { name: 'Hitpoints', iconSrc: '/images/skills/hitpoints.png', href: '/skills/hitpoints', description: 'Increase your survivability.' },
        { name: 'Mining', iconSrc: '/images/skills/mining.png', href: '/skills/mining', description: 'Gather ores from rocks.' },
        { name: 'Strength', iconSrc: '/images/skills/strength.png', href: '/skills/strength', description: 'Hit harder in combat.' },
        { name: 'Agility', iconSrc: '/images/skills/agility.png', href: '/skills/agility', description: 'Traverse Gielinor faster.' },
        { name: 'Smithing', iconSrc: '/images/skills/smithing.png', href: '/skills/smithing', description: 'Forge armor and weapons.' },
        { name: 'Defence', iconSrc: '/images/skills/defence.png', href: '/skills/defence', description: 'Reduce damage taken.' },
        { name: 'Herblore', iconSrc: '/images/skills/herblore.png', href: '/skills/herblore', description: 'Create powerful potions.' },
        { name: 'Fishing', iconSrc: '/images/skills/fishing.png', href: '/skills/fishing', description: 'Catch a variety of fish.' },
        { name: 'Ranged', iconSrc: '/images/skills/ranged.png', href: '/skills/ranged', description: 'Master ranged combat.' },
        { name: 'Thieving', iconSrc: '/images/skills/thieving.png', href: '/skills/thieving', description: 'Pilfer goods and XP.' },
        { name: 'Cooking', iconSrc: '/images/skills/cooking.png', href: '/skills/cooking', description: 'Prepare food for healing.' },
        { name: 'Prayer', iconSrc: '/images/skills/prayer.png', href: '/skills/prayer', description: 'Invoke the power of gods.' },
        { name: 'Crafting', iconSrc: '/images/skills/crafting.png', href: '/skills/crafting', description: 'Create jewelry and armor.' },
        { name: 'Firemaking', iconSrc: '/images/skills/firemaking.png', href: '/skills/firemaking', description: 'Burn logs for warmth and XP.' },
        { name: 'Magic', iconSrc: '/images/skills/magic.png', href: '/skills/magic', description: 'Cast powerful spells.' },
        { name: 'Fletching', iconSrc: '/images/skills/fletching.png', href: '/skills/fletching', description: 'Craft bows and arrows.' },
        { name: 'Woodcutting', iconSrc: '/images/skills/woodcutting.png', href: '/skills/woodcutting', description: 'Chop down trees for logs.' },
        { name: 'Runecraft', iconSrc: '/images/skills/runecraft.png', href: '/skills/runecraft', description: 'Craft magical runes.' },
        { name: 'Slayer', iconSrc: '/images/skills/slayer.png', href: '/skills/slayer', description: 'Hunt monsters for assignments.' },
        { name: 'Farming', iconSrc: '/images/skills/farming.png', href: '/skills/farming', description: 'Grow crops and herbs.' },
        { name: 'Construction', iconSrc: '/images/skills/construction.png', href: '/skills/construction', description: 'Build your own house.' },
        { name: 'Hunter', iconSrc: '/images/skills/hunter.png', href: '/skills/hunter', description: 'Track and trap creatures.' },
    ];
</script>

<div class="max-w-6xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
    <header class="mb-10 text-center">
        <h1 class="text-h1 text-theme-text-primary tracking-tight">
            Skill Calculators & Guides
        </h1>
        <p class="mt-3 text-lg text-theme-text-secondary max-w-2xl mx-auto">
            Select a skill below to access its specific calculator, guides, and tools.
        </p>
    </header>

    {#if skills.length > 0}
        <div class="grid grid-cols-2 xs:grid-cols-3 sm:grid-cols-4 md:grid-cols-5 lg:grid-cols-6 xl:grid-cols-7 gap-4 md:gap-5">
            {#each skills as skill (skill.name)}
                <a href={skill.href}
                   title={skill.description || skill.name}
                   class="group relative flex flex-col items-center justify-center p-6 bg-theme-card-bg rounded-lg shadow-card  border-theme-border focus-within:ring-2 focus-within:ring-inset focus-within:ring-theme-accent hover:border-theme-accent/70 transition-all duration-200 ease-in-out aspect-square transform hover:-translate-y-1 hover:shadow-xl">
                    
                    <img src={skill.iconSrc} alt="{skill.name} icon" class="h-10 w-10 md:h-12 md:w-12 mb-2 object-contain group-hover:scale-110 transition-transform" />
                    
                    <span class="block text-xs sm:text-sm font-medium text-theme-text-primary group-hover:text-theme-accent transition-colors text-center truncate w-full">
                        {skill.name}
                    </span>
                </a>
            {/each}
        </div>
    {:else}
        <div class="text-center py-10">
            <p class="text-theme-text-secondary text-lg">Skill calculators are being developed!</p>
        </div>
    {/if}
</div>