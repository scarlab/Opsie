import { Button } from "@/components/cn/button";
import TeamGeneralInfo from "./TeamGeneralInfo";
import TeamMembers from "./TeamMembers";
import { Plus } from "lucide-react";
import { useSearchParams } from "react-router-dom";
import NewTeam from "./NewTeam";

export default function TeamSettingsView() {

    const [_, setSearchParams] = useSearchParams();

    return (
        <div className="space-y-5 pb-20">
            <div className="flex items-center justify-between px-3">
                <h1 className="text-2xl font-bold ">Team Settings</h1>

                <NewTeam />
            </div>

            <TeamGeneralInfo />
            <TeamMembers />
        </div>
    )
}
