export const LocalStorageKey = {
    User: 'user',
    UserCachedOn: 'uco',
    DefaultTeam: 'team',
} as const;
export type LocalStorageKeyType = typeof LocalStorageKey[keyof typeof LocalStorageKey];


export const SystemRoles = {
    Owner: 'owner',
    Admin: 'admin',
    Staff: 'staff',
} as const;

export type SystemRolesType = typeof SystemRoles[keyof typeof SystemRoles];