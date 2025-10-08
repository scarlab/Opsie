export interface LoginPayload {
    email: string;
    password: string;
}

export interface NewOwnerPayload {
    name: string;
    email: string;
    password: string;
    confirmPassword: string;
}