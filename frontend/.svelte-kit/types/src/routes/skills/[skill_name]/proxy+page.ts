// @ts-nocheck
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types'; 
import type { SkillData } from '$lib/types'; 

export const load = async ({ params, fetch }: Parameters<PageLoad>[0]) => {
    const skillNameParam = params.skill_name?.toLowerCase();

    if (!skillNameParam) {
        throw error(400, 'Skill name parameter is missing.');
    }

    try {
        const response = await fetch(`http://localhost:8080/api/skill-data/${skillNameParam}`);

        if (!response.ok) {
            if (response.status === 404) {
                throw error(404, `Skill data not found for: ${skillNameParam}. Ensure backend endpoint and data file exist.`);
            }
            const errorText = await response.text();
            throw error(response.status, `Failed to load data for '${skillNameParam}': ${errorText || response.statusText}`);
        }

        const skillData = await response.json() as SkillData;

        if (!skillData || typeof skillData.skillNameDisplay !== 'string' || !Array.isArray(skillData.trainingMethods)) {
             throw error(500, `Invalid or incomplete data structure received from backend for skill: ${skillNameParam}`);
        }
        
        return {
            skillData: skillData 
        };
    } catch (e: any) {
        console.error(`Error in load function for ${skillNameParam}:`, e);

        if (e.status && e.body?.message) { 
            throw error(e.status, e.body.message);
        }
        throw error(500, `Could not load skill data for '${skillNameParam}'. Error: ${e.message || 'Unknown server error'}`);
    }
};