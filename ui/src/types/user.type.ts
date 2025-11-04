import type { SystemRolesType } from "@/constants";


export interface UserModel {
    id: string;
    display_name: string;
    email: string;
    avatar: string;
    system_role: SystemRolesType;
    preference: Object;
    is_active: boolean;
    updated_at: Date;
    created_at: Date;
}

export interface NewOwnerPayload {
    display_name: string;
    email: string;
    password: string;
    confirm_password: string;
}


export interface NewUserPayload {
    display_name: string;
    email: string;
    password: string;
    system_role: SystemRolesType;
    avatar: string;
}


export interface UpdateUserPayload {
    display_name: string;
    email: string;
    system_role: SystemRolesType
}