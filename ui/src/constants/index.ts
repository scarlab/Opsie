export const LocalStorageKey = {
    User: 'user',
    UserCachedOn: 'uco',
    DefaultTeam: 'team',
} as const;
export const SystemRoles = {
    Owner: 'owner',
    Admin: 'admin',
    Staff: 'staff',
} as const;

export type LocalStorageKeyType = typeof LocalStorageKey[keyof typeof LocalStorageKey];