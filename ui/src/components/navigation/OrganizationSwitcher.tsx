import { useState } from "react";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "../cn/dropdown-menu";
import { AudioWaveform, ChevronsUpDown, Command, GalleryVerticalEnd, Plus } from "lucide-react";

const organizations = [
    {
        name: "Acme Inc",
        logo: GalleryVerticalEnd,
        plan: "Free",
    },
    {
        name: "Acme Corp.",
        logo: AudioWaveform,
        plan: "Free",
    },
    {
        name: "Evil Corp.",
        logo: Command,
        plan: "Free",
    },
]

export default function OrganizationSwitcher() {

    const [activeOrganization, setActiveOrganization] = useState(organizations[0])
    if (!activeOrganization) {
        return null
    }
    return (
        <div className="">

            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <div className="hover:bg-accent dark:hover:bg-accent/40 cursor-pointer flex items-center gap-2 transition-all rounded px-2 py-">
                        <div className="space-x-3">
                            <span className="truncate font-medium">{activeOrganization.name}</span>
                            <span className="truncate text-xs border rounded-2xl px-2.5 bg-accent/60">{activeOrganization.plan}</span>
                        </div>
                        <ChevronsUpDown size={17} />
                    </div>
                </DropdownMenuTrigger>
                <DropdownMenuContent
                    className="w-(--radix-dropdown-menu-trigger-width) min-w-56 rounded-lg"
                    align="start"
                    side={"right"}
                    sideOffset={4}
                >
                    <DropdownMenuLabel className="text-muted-foreground text-xs">
                        Organizations
                    </DropdownMenuLabel>
                    {organizations.map((organization, i) => (
                        <DropdownMenuItem
                            key={organization.name + i}
                            onClick={() => setActiveOrganization(organization)}
                            className="gap-2 p-2"
                        >
                            {organization.name}

                        </DropdownMenuItem>
                    ))}
                    <DropdownMenuSeparator />
                    <DropdownMenuItem className="gap-2 p-2">
                        <div className="flex size-6 items-center justify-center rounded-md border bg-transparent">
                            <Plus className="size-4" />
                        </div>
                        <div className="text-muted-foreground font-medium">Add organization</div>
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    )
}
