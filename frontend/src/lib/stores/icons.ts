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
        src: '/icons/experience.png', // Dedicated experience icon
        preload: true,
        fallback: '/skills/magic.png'
    },
    coins: {
        src: '/icons/coins.png', // Dedicated coins icon
        preload: true,
        fallback: '/skills/thieving.png'
    },
    clock: {
        src: '/icons/clock.png', // Dedicated clock icon
        preload: true,
        fallback: '/skills/construction.png'
    },
    chart: {
        src: '/skills/slayer.png',
        preload: false,
        fallback: '/skills/strength.png'
    }
} as const;

// Skill icons - lazy loaded (using existing PNG files)
export const SKILL_ICONS = {
    firemaking: '/skills/firemaking.png',
    runecraft: '/skills/runecraft.png',
    thieving: '/skills/thieving.png',
    hunter: '/skills/hunter.png',
    herblore: '/skills/herblore.png',
    attack: '/skills/attack.png',
    agility: '/skills/agility.png',
    mining: '/skills/mining.png'
} as const;

// Tool icons - lazy loaded  
export const TOOL_ICONS = {
    knight: '/tools/knight_of_ardougne.png',
    birdhouse: '/birdhouse/redwood_bird_house.png',
    herbiboar: '/herbiboar/herbiboar.png'
} as const;

// Pet icons - using skill icons as placeholders
export const PET_ICONS = {
    phoenix: '/skills/firemaking.png', // Phoenix relates to firemaking
    abyssal_protector: '/skills/runecraft.png', // Abyssal protector relates to runecraft
    herbi: '/herbiboar/herbiboar.png' // Herbi pet relates to herbiboar hunting
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