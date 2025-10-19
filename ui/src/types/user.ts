

export interface User {
    id: number;
    display_name: string;
    email: string;
    password: string;
    avatar: string;
    system_role: string;
    preference: Object;
    is_active: boolean;
    updated_at: Date;
    created_at: Date;
}

export interface NewOwnerPayload {
    display_name: string;
    email: string;
    password: string;
    confirmPassword: string;
}