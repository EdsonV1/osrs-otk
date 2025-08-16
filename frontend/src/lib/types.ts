export interface ArdyKnightResult {
    calculated_success_rate: number;
    effective_xp_per_attempt: number;
    effective_gp_per_attempt: number;
    xp_hour: number;
    gp_hour: number;
    damage_per_hour: number;
    food_needed_per_hour: number;
    profit_per_hour: number;
    current_thieving_level: number;
    target_thieving_level: number;
    current_total_xp: number;
    target_total_xp: number;
    xp_to_target: number;
    hours_to_target: number;
    pickpockets_to_target: number;
}

export interface ArdyKnightInput {
    current_thieving_xp?: number | null;
    current_thieving_level?: number | null;
    target_thieving_xp?: number | null;
    target_thieving_level?: number | null;
    has_ardy_med: boolean;
    has_thieving_cape: boolean;
    has_rogues_outfit: boolean;
    has_shadow_veil: boolean;
    hourly_pickpockets: number;
    food_heal_amount: number;
    food_cost: number;
}

export interface BirdhouseFormState {
    selectedLogType: string; // Lowercase value, e.g., "redwood", "magic" (maps to API 'type')
    totalBirdhouses: number; // Total number of birdhouses to calculate for (maps to API 'quantity')
    // hunterLevel, numberOfHouses, numberOfRuns are removed from this specific form state
}

export interface BirdhouseApiInput {
    type: string;     // Lowercase birdhouse type, e.g., "redwood"
    quantity: number; // Total number of individual birdhouses
}

export interface SeedDropInfo {
    quantity: number;
    value: number;
}

export interface BirdhouseApiResult {
    estimated_nests: number;
    hunter_xp: number;
    crafting_xp: number;
    days_low_efficiency: number;
    days_medium_efficiency: number;
    days_high_efficiency: number;
    seed_drops: {
        [seedName: string]: SeedDropInfo; // e.g., "acorn": { quantity: X, value: Y }
    };
    total_loot: number; // Total GP value derived from seed_drops by the backend
}

export interface TrainingMethod {
    id: string;                 // Unique identifier for this method, e.g., "canifis_rooftop"
    name: string;               // Display name, e.g., "Canifis Rooftop Course"
    levelReq: number;           // Level required
    xpRate: number;             // Primary XP per hour
    marksPerHour?: number;      // Agility specific: Marks of Grace per hour (optional)
    xpPerAction?: number;       // Optional: XP per single action (lap, log, ore)
    actionName?: string;        // Optional: Name of the action, e.g., "lap", "log"
    alternativeXpRate?: { type: string, rate: number }[]; // For methods giving XP in other skills
    location?: string;          // In-game location
    itemsRequired?: string[];   // List of item names
    questsRequired?: string[];  // List of quest names
    notes?: string;             // Additional notes or tips
    tags?: string[];            // For filtering/categorization, e.g., ["rooftop", "afk", "profitable"]
    type?: string;              // General category, e.g., "Rooftop Course", "Mining Ore", "Combat"
}

export interface SkillData {
    skillNameCanonical: string; // e.g., "agility", "mining" (lowercase, for API path)
    skillNameDisplay: string;   // e.g., "Agility", "Mining" (for display in titles)
    description?: string;       // A brief description of the skill
    trainingMethods: TrainingMethod[];
}

export interface SkillCalculatorPageData {
    skillData: SkillData;
}

export interface CalculatedMethodResult {
    methodId: string;
    methodName: string;
    actionsNeeded?: number;
    timeToCompleteHours?: number;
    xpFromThisMethod?: number;
    marksOfGraceEarned?: number;
    amylasePurchased?: number; 
}

export interface SkillCalculationOutput {
    totalXPToGain: number;
    methodsBreakdown: CalculatedMethodResult[];
    overallTimeToTarget?: string; 
    totalMarksOfGrace?: number;   
    totalAmylase?: number;       
}

export interface WintertodtFormState {
    currentLevel: number;
    targetLevel: number;
    strategy: string;
    customPointsPerRound?: number;
    customMinutesPerRound?: number;
    skillLevels: {
        herblore: number;
        mining: number;
        fishing: number;
        crafting: number;
        farming: number;
        woodcutting: number;
    };
}

export interface WintertodtApiInput {
    current_level: number;
    target_level: number;
    strategy: string;
    custom_points_per_round?: number;
    custom_minutes_per_round?: number;
    skill_levels: {
        herblore: number;
        mining: number;
        fishing: number;
        crafting: number;
        farming: number;
        woodcutting: number;
    };
}

export interface WintertodtApiResult {
    current_level: number;
    target_level: number;
    xp_needed: number;
    rounds_needed: number;
    total_experience: number;
    average_exp_hour: number;
    pet_chance: number;
    estimated_loot: { [itemName: string]: number };
    total_value: number;
    total_time: number;
    strategy: string;
    points_per_round: number;
    minutes_per_round: number;
    total_points_earned: number;
}

// GOTR (Guardians of the Rift) Types
export interface GOTRFormState {
    currentLevel: number;
    targetLevel: number;
}

export interface GOTRApiInput {
    current_level: number;
    target_level: number;
}

export interface GOTRReward {
    name: string;
    quantity: number;
    value: number;
    drop_rate?: string;
}

export interface GOTRApiResult {
    current_level: number;
    target_level: number;
    xp_needed: number;
    games_needed: number;
    hours_needed: number;
    average_xp_per_game: number;
    average_xp_per_hour: number;
    total_reward_rolls: number;
    pet_chance_percentage: number;
    estimated_rewards: GOTRReward[];
    total_reward_value: number;
    gp_per_hour: number;
}

export interface PlayerStats {
    username: string;
    overall: number;
    attack: number;
    defence: number;
    strength: number;
    hitpoints: number;
    ranged: number;
    prayer: number;
    magic: number;
    cooking: number;
    woodcutting: number;
    fletching: number;
    fishing: number;
    firemaking: number;
    crafting: number;
    smithing: number;
    mining: number;
    herblore: number;
    agility: number;
    thieving: number;
    slayer: number;
    farming: number;
    runecrafting: number;
    hunter: number;
    construction: number;
}