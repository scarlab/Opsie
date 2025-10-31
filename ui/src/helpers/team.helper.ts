import { LocalStorageKey } from "@/constants";
import type { TeamModel } from "@/types/team.type";



/**
 * Save team to localStorage
 */
export function setLocalTeam(data: TeamModel) {
    if (!data) return;
    try {
        localStorage.setItem(LocalStorageKey.DefaultTeam, JSON.stringify(data));
    } catch (err) {
        console.error("Failed to store team:", err);
    }
}

/**
 * Retrieve cached team
 */
export function getLocalTeam(): TeamModel | null {
    try {
        const teamRaw = localStorage.getItem(LocalStorageKey.DefaultTeam);
        if (!teamRaw) return null;

        return JSON.parse(teamRaw) as TeamModel;
    } catch (err) {
        console.error("Failed to read cached team:", err);
        return null;
    }
}

