import { LocalStorageKey } from "@/constants";
import type { AuthUser } from "@/types/auth.type";

// how long cache stays valid (e.g., 3 min)
const CACHE_TTL_MS = 1 * 60 * 1000;

/**
 * Save auth user + timestamp to localStorage
 */
export function setLocalAuthUser(data: AuthUser) {
    if (!data) return;

    try {
        localStorage.setItem(LocalStorageKey.User, JSON.stringify(data));
        localStorage.setItem(LocalStorageKey.UserCachedOn, Date.now().toString());
    } catch (err) {
        console.error("Failed to store auth user:", err);
    }
}

/**
 * Retrieve cached user if still valid
 */
export function getLocalAuthUser(): AuthUser | null {
    try {
        const userRaw = localStorage.getItem(LocalStorageKey.User);
        const cachedOn = localStorage.getItem(LocalStorageKey.UserCachedOn);

        if (!userRaw || !cachedOn) return null;

        const cachedAt = parseInt(cachedOn, 10);
        const isExpired = Date.now() - cachedAt > CACHE_TTL_MS;

        if (isExpired) {
            removeLocalAuthUser();
            return null;
        }

        return JSON.parse(userRaw) as AuthUser;
    } catch (err) {
        console.error("Failed to read cached auth user:", err);
        removeLocalAuthUser();
        return null;
    }
}

/**
 * Remove cached auth user (e.g., on logout or session expiry)
 */
export function removeLocalAuthUser() {
    try {
        localStorage.removeItem(LocalStorageKey.User);
        localStorage.removeItem(LocalStorageKey.UserCachedOn);
    } catch (err) {
        console.error("Failed to clear auth user cache:", err);
    }
}
