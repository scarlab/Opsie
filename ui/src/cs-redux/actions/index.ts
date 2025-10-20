import { OrganizationAction } from "./organization.action";
import { AuthAction } from "./auth.action";
import { GlobalAction } from "./global.action";
import { UserAction } from "./user.action";

export const Actions = {
    organization: new OrganizationAction(),
    global: new GlobalAction(),
    auth: new AuthAction(),
    user: new UserAction(),
}