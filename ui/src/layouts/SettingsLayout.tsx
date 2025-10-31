
import ViewHeader from "@/components/utils/ViewHeader";
import { cn } from "@/lib/utils";
import { Link, Outlet, useLocation } from "react-router-dom";

const data = [
    { label: 'General', to: '/settings/general' },
    { label: 'Account', to: '/settings/account' },
    { label: 'Team', to: '/settings/team' },
    { label: 'Integrations', to: '/settings/integrations' },
    { label: 'Billing', to: '/settings/billing' },
]

export default function SettingsLayout() {

    const { pathname } = useLocation();

    return (
        <div>
            <ViewHeader title="Settings" />

            <main className="mx-auto max-w-4xl w-full flex min-h-screen ">

                <div className="w-[calc(100%-10rem)] px-3 mt-3">
                    <Outlet />
                </div>

                <div className="w-[10rem] sticky top-[calc(var(--header-height))] h-[calc(100svh-var(--header-height))] border-l ps-8 flex flex-col  justify-center  px-3 py-4 gap-2">
                    {
                        data.map((item, i) => (
                            <Link to={item.to} key={`setting_m_${i}`} className={cn("hover:underline", pathname.startsWith(item.to) ? "font-bold underline text-primary" : "")}>
                                {item.label}
                            </Link>
                        ))
                    }
                </div>
            </main>
        </div>
    );
}
