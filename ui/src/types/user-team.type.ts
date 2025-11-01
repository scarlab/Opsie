


export interface AddUserToTeamPayload {
    user_id: string;
    team_id: string;
    is_default: boolean;
}


export interface RemoveUserToTeamPayload {
    user_id: string;
    team_id: string;
}

export interface TeamMember {
    id: string;
    team_id: string;
    display_name: string;
    email: string;
    avatar: string;
    system_role: string;
    is_active: boolean;
    joined_at: Date;
}

