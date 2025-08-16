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
        src: '/images/skills/runecraft.png', // Using runecraft as experience icon
        preload: true,
        fallback: '/images/skills/magic.png'
    },
    coins: {
        src: '/images/tools/knight_of_ardougne.png', // Using knight icon for coins temporarily
        preload: true,
        fallback: '/images/skills/thieving.png'
    },
    clock: {
        src: '/images/skills/agility.png', // Using agility icon for time/speed
        preload: true,
        fallback: '/images/skills/construction.png'
    },
    chart: {
        src: '/images/skills/slayer.png',
        preload: false,
        fallback: '/images/skills/strength.png'
    }
} as const;

// Skill icons - lazy loaded (using existing PNG files)
export const SKILL_ICONS = {
    firemaking: '/images/skills/firemaking.png',
    runecraft: '/images/skills/runecraft.png',
    thieving: '/images/skills/thieving.png',
    hunter: '/images/skills/hunter.png',
    attack: '/images/skills/attack.png',
    agility: '/images/skills/agility.png',
    mining: '/images/skills/mining.png'
} as const;

// Tool icons - lazy loaded  
export const TOOL_ICONS = {
    knight: '/images/tools/knight_of_ardougne.png',
    birdhouse: '/images/birdhouse/redwood_bird_house.png'
} as const;

// Pet icons - using skill icons as placeholders
export const PET_ICONS = {
    phoenix: '/images/skills/firemaking.png', // Phoenix relates to firemaking
    abyssal_protector: '/images/skills/runecraft.png' // Abyssal protector relates to runecraft
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