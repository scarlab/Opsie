export interface LoginPayload {
    email: string;
    password: string;
}

export interface AuthUser {
    id: string;
    display_name: string;
    email: string;
    avatar: string;
    system_role: string;
    preference: Object;
    is_active: boolean;
    reset_pass: boolean;
}