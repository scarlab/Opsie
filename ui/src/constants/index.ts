export const LocalStorageKey = {
    User: 'user',
    UserCachedOn: 'uco',
} as const;

export type LocalStorageKeyType = typeof LocalStorageKey[keyof typeof LocalStorageKey];