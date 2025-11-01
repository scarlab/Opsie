export function getInitials(name: string): string {
    if (!name) return "";

    // Split by spaces, remove empty parts (for extra spaces)
    const parts = name.trim().split(/\s+/);

    // Take first letter of first two words (first and last name typically)
    const initials = parts
        .slice(0, 2)
        .map(word => word.charAt(0).toUpperCase())
        .join("");

    return initials;
}
