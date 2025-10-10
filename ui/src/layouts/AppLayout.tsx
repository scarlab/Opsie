import AppHeader from "@/components/navigation/AppHeader";
import AppSidebar from "@/components/navigation/AppSidebar";
import { useState } from "react";
import { Outlet } from "react-router-dom";


export default function AppLayout() {
    const [isCollapsed, setIsCollapsed] = useState(false);

    return (
        <div>
            <AppHeader />
            <section className="flex">
                <AppSidebar collapsed={isCollapsed} setCollapsed={setIsCollapsed} />
                <main
                    className="transition-all duration-300 grow"
                    style={{
                        paddingLeft: isCollapsed
                            ? "var(--sidebar-collapse-width)"
                            : "var(--sidebar-width)",
                    }}
                >
                    <Outlet />
                </main>
            </section>
        </div>
    );
}
