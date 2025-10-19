import { AuthAction } from "./auth.action";
import { GlobalAction } from "./global.action";

export const Actions = {
    global: new GlobalAction(),
    auth: new AuthAction(),
}