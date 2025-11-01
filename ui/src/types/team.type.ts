
export interface TeamModel {
    id: string;
    name: string;
    slug: string;
    description?: string;
    updated_at: Date;
    created_at: Date;
}

export interface UserTeam extends TeamModel {
    is_default?: boolean;
    joined_at?: Date;
}


export interface NewTeamPayload {
    name: string;
    description: string;
}

