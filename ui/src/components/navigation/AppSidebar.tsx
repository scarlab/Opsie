import { cn } from "@/lib/utils";
import { ArrowLeft, ArrowRight, Blocks, Box, LayoutDashboard, Server, Shapes, User } from "lucide-react";
import { Link, useLocation } from "react-router-dom";


type AppSidebarProps = {
    collapsed: boolean;
    setCollapsed: React.Dispatch<React.SetStateAction<boolean>>;
};

const data = [
    {
        label: "",
        items: [
            {
                label: "Overview",
                icon: LayoutDashboard,
                link: "/",
            },
            {
                label: "Nodes",
                icon: Server,
                link: "/nodes",
            },
            {
                label: "Users",
                icon: User,
                link: "/users",
            },
            // {
            //     label: "Monitor",
            //     icon: ScanHeart,
            //     link: "/monitor",
            // },
        ]
    },
    {
        label: "Workspace",
        items: [
            {
                label: "Projects",
                icon: Shapes,
                link: "/projects",
            },
            {
                label: "Resources",
                icon: Box,
                link: "/resources",
            },
            {
                label: "Apps",
                icon: Blocks,
                link: "/apps",
            },
        ]
    }
]

export default function AppSidebar({ collapsed, setCollapsed }: AppSidebarProps) {
    const location = useLocation();
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
                    data.map((group, i) => (
                        <div key={`s-b_x_item_group_${i}`}>
                            {group.label && (
                                <div
                                    className={cn(
                                        "mt-5 px-3 transition-all duration-300",
                                        collapsed ? "px-0" : "px-3"
                                    )}
                                >
                                    {collapsed ? (
                                        <div className="h-5 flex flex-col justify-end">
                                            <div className="h-1 bg-muted/40  " />
                                        </div>
                                    ) : (
                                        <p className="text-sm text-muted-foreground overflow-hidden whitespace-nowrap">
                                            {group.label}
                                        </p>
                                    )}
                                </div>
                            )}


                            {
                                group.items.map((section, j) => {
                                    const isActive =
                                        location.pathname === section.link ||
                                        location.pathname.startsWith(section.link + "/");

                                    return (
                                        <div key={`s-b_x_group_${i}item_${j}`} className="mb-1.5">
                                            <Link
                                                to={section.link}
                                                className={cn(
                                                    "flex items-center hover:bg-muted/30 py-1.5 px-3 rounded transition-all duration-200 gap-3",
                                                    isActive && "bg-muted/20"
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
                                    )
                                })
                            }
                        </div>
                    ))
                }
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
