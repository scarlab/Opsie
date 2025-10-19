import { AuthAction } from "./auth.action";
import { GlobalAction } from "./global.action";
import { UserAction } from "./user.action";

export const Actions = {
    global: new GlobalAction(),
    auth: new AuthAction(),
    user: new UserAction(),
}