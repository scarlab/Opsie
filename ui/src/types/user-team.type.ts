


export interface AddUserToTeamPayload {
    user_id: number;
    team_id: number;
    is_default: boolean;
    is_admin: boolean;
    invited_by?: number;
}


export interface RemoveUserToTeamPayload {
    user_id: number;
    team_id: number;
}

