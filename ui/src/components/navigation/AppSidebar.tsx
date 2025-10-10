import { cn } from "@/lib/utils";
import { ArrowLeft, ArrowRight, LayoutDashboard, Server } from "lucide-react";
import { Link } from "react-router-dom";


type AppSidebarProps = {
    collapsed: boolean;
    setCollapsed: React.Dispatch<React.SetStateAction<boolean>>;
};

const data = [
    {
        label: "Overview",
        icon: LayoutDashboard,
        link: "/",
        items: []
    },
    {
        label: "Nodes",
        icon: Server,
        link: "/nodes",
        items: []
    },
]

export default function AppSidebar({ collapsed, setCollapsed }: AppSidebarProps) {

    return (
        <aside
            className={`transition-all duration-300 h-[calc(100vh-var(--header-height))] bg-secondary border-r fixed top-[var(--header-height)] left-0 p-3 flex flex-col gap-3`}
            style={{
                width: collapsed
                    ? "var(--sidebar-collapse-width)"
                    : "var(--sidebar-width)",
            }}
        >
            <div className="grow">
                {
                    data.map((section, i) => (
                        <div key={`s-b_x_item_${i}`} className="mb-1.5">
                            <Link
                                to={section.link}
                                className={cn(
                                    "flex items-center hover:bg-muted/30 py-1.5 px-3 rounded transition-all duration-200 gap-3",
                                    section.link === location.pathname ? "bg-muted/20 font-medium" : "font-normal"
                                )}
                            >
                                <section.icon size={18} className="shrink-0" />

                                <span
                                    className={`block overflow-hidden whitespace-nowrap transition-[opacity,width] duration-300  ${collapsed ? "opacity-0 w-0" : "opacity-100 w-auto"
                                        }`}
                                >
                                    {section.label}
                                </span>
                            </Link>

                        </div>
                    ))}
            </div>

            <div>
                <div className={cn("cursor-pointer", "flex items-center hover:bg-muted/30 py-1.5 px-3 rounded transition-all duration-200 gap-3")} onClick={() => setCollapsed(!collapsed)}>
                    {
                        collapsed ?
                            <ArrowRight className="shrink-0" size={18} />
                            :
                            <ArrowLeft className="shrink-0" size={18} />
                    }

                    <span
                        className={`block overflow-hidden whitespace-nowrap transition-[opacity,width] duration-300  ${collapsed ? "opacity-0 w-0" : "opacity-100 w-auto"
                            }`}
                    >
                        Collapse Menu
                    </span>
                </div>
            </div>
        </aside>
    )
}
