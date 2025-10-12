import AppHeader from "@/components/navigation/AppHeader";
import { Outlet } from "react-router-dom";


export default function SettingsLayout() {


    return (
        <div>
            <AppHeader />
            <main className="mx-auto max-w-3xl  px-3">
                <Outlet />
            </main>
        </div>
    );
}
