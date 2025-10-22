import { TeamAction } from "./team.action";
import { AuthAction } from "./auth.action";
import { GlobalAction } from "./global.action";
import { UserAction } from "./user.action";

export const Actions = {
    team: new TeamAction(),
    global: new GlobalAction(),
    auth: new AuthAction(),
    user: new UserAction(),
}