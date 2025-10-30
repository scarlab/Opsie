import AppHeader from "@/components/navigation/AppHeader";
import { Link, Outlet } from "react-router-dom";

const data = [
    { label: 'General', to: '/settings/general' },
    { label: 'Account', to: '/settings/account' },
    { label: 'Team', to: '/settings/team' },
    { label: 'Integrations', to: '/settings/integrations' },
    { label: 'Billing', to: '/settings/billing' },
]

export default function SettingsLayout() {


    return (
        <div className="">
            <AppHeader />
            {/* <div className=" sticky top-[var(--header-height)] h-[var(--header-height)] z-10 border-b bg-background">
                <div className="w-full max-w-4xl mx-auto flex items-center h-full px-3">
                    <h1 className="text-2xl font-bold ">Account Settings</h1>
                </div>
            </div> */}

            <main className="mx-auto max-w-4xl w-full flex min-h-screen ">
                <div className="w-[10rem] sticky top-[calc(var(--header-height))] h-[calc(100svh-var(--header-height))] border-r flex flex-col  px-3 py-4 gap-2">
                    {
                        data.map((item, i) => (
                            <Link to={item.to} key={`setting_m_${i}`}>
                                {item.label}
                            </Link>
                        ))
                    }
                </div>
                <div className="w-[calc(100%-10rem)] px-3 mt-3">
                    <Outlet />
                </div>
            </main>
        </div>
    );
}
