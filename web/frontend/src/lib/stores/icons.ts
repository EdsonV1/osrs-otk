// Performance-optimized icon store
// Centralized icon paths for better caching and maintainability

export interface IconConfig {
    src: string;
    preload?: boolean;
    fallback?: string;
}

// Common metric icons - these will be preloaded for performance
export const METRIC_ICONS = {
    experience: {
        src: '/images/icons/experience.svg',
        preload: true,
        fallback: '/images/icons/experience.png'
    },
    coins: {
        src: '/images/icons/coins.svg', 
        preload: true,
        fallback: '/images/icons/coins.png'
    },
    clock: {
        src: '/images/icons/clock.svg',
        preload: true,
        fallback: '/images/icons/clock.png'
    },
    chart: {
        src: '/images/icons/chart.svg',
        preload: false,
        fallback: '/images/icons/chart.png'
    }
} as const;

// Skill icons - lazy loaded
export const SKILL_ICONS = {
    firemaking: '/images/skills/firemaking.svg',
    runecraft: '/images/skills/runecraft.svg',
    thieving: '/images/skills/thieving.svg',
    hunter: '/images/skills/hunter.svg',
    attack: '/images/skills/attack.svg',
    agility: '/images/skills/agility.svg',
    mining: '/images/skills/mining.svg'
} as const;

// Tool icons - lazy loaded  
export const TOOL_ICONS = {
    knight: '/images/tools/knight_of_ardougne.svg',
    birdhouse: '/images/birdhouse/redwood_bird_house.svg'
} as const;

// Pet icons - lazy loaded
export const PET_ICONS = {
    phoenix: '/images/pets/phoenix.svg',
    abyssal_protector: '/images/pets/abyssal_protector.svg'
} as const;

// Preload critical icons for better performance
export function preloadCriticalIcons() {
    Object.values(METRIC_ICONS).forEach(icon => {
        if (icon.preload) {
            const link = document.createElement('link');
            link.rel = 'preload';
            link.href = icon.src;
            link.as = 'image';
            document.head.appendChild(link);
        }
    });
}

// Get icon with fallback
export function getIcon(category: keyof typeof METRIC_ICONS, key: string): IconConfig {
    return METRIC_ICONS[category] || { src: key, preload: false };
}