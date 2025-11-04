import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuShortcut,
    DropdownMenuTrigger
} from "../cn/dropdown-menu";
import { Check, ChevronsUpDown } from "lucide-react";
import { Actions, useCsDispatch, useCsSelector } from "@/cs-redux";
import { useEffect, useState } from "react";
import type { TeamModel } from "@/types/team.type";
import { getLocalTeam } from "@/helpers/team.helper";


export default function TeamSwitcher() {
    const { auth_default_team: user_default_team, auth_teams: user_teams } = useCsSelector(state => state.team);


    const dispatch = useCsDispatch();

    const [team, setTeam] = useState<TeamModel | null>(getLocalTeam());


    useEffect(() => {
        dispatch(Actions.team.GetAllTeamOfAuthUser());
        dispatch(Actions.team.GetDefaultTeamOfAuthUser());
    }, [dispatch]);


    useEffect(() => {
        if (user_default_team) setTeam(user_default_team);
    }, [user_default_team]);




    async function handleSwitchTeam(id: string) {
        if (id === user_default_team?.id) return;
        dispatch(Actions.team.SetDefaultTeamOfAuthUser({ id }))
    }



    if (!user_default_team || !user_teams) {
        return <div className="w-32 bg-accent/40 h-7 rounded animate-pulse" role="status">
        </div>
    }


    return (
        <div className="">

            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <div className="hover:bg-accent- dark:hover:bg-accent/40 border-b border-transparent hover:border-accent cursor-pointer flex items-center gap-2 transition-all rounded px-4 py-1">
                        <div className="space-x-3">
                            <span className="truncate font-medium">{team?.name}</span>
                            {/* <span className="truncate text-xs border rounded-2xl px-2.5 bg-accent/60">{team?.id}</span> */}
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
                        Year Teams
                    </DropdownMenuLabel>
                    {user_teams.map((team, i) => (
                        <DropdownMenuItem
                            key={team.name + i}
                            onClick={() => handleSwitchTeam(team.id)}
                            className="gap-2 p-2"
                        >
                            {team.name}
                            {team.is_default &&
                                <DropdownMenuShortcut>
                                    <Check />
                                </DropdownMenuShortcut>
                            }
                        </DropdownMenuItem>
                    ))}
                    {/* <DropdownMenuSeparator />
                    <DropdownMenuItem onClick={handleAddTeam} className="gap-2 p-2" >
                        <div className="flex size-6 items-center justify-center rounded-md border bg-transparent">
                            <Plus className="size-4" />
                        </div>
                        <div className="text-muted-foreground font-medium">Add team</div>
                    </DropdownMenuItem> */}
                </DropdownMenuContent>
            </DropdownMenu>

        </div>
    )
}
