import { Theme } from "@/components/theme";
import Config from "@/config";
import { Outlet } from "react-router-dom";

export default function AuthLayout() {
    return (
        <div className="">
            <div className="px-5 pb-3 fixed bottom-0 left-0 w-full flex items-end justify-between">
                <div>
                    <small className="text-muted-foreground">&copy; scarlab.in / opsie - {Config.version}</small>
                </div>
                <div>
                    <Theme />
                </div>
            </div>
            <Outlet />
        </div>
    );
}
