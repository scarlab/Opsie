import TeamGeneralInfo from "./TeamGeneralInfo";
import TeamMembers from "./TeamMembers";
import NewTeam from "./NewTeam";
import { useSearchParams } from "react-router-dom";
import TeamList from "./TeamList";
import { useEffect } from "react";
import { Actions, useCsDispatch, useCsSelector } from "@/cs-redux";
import { ArrowLeft } from "lucide-react";
import TeamSlice from "@/cs-redux/slices/team.slice";

export default function TeamSettingsView() {
    const dispatch = useCsDispatch();
    const [searchParams, setSearchParams] = useSearchParams();
    const teamId = searchParams.get("team");
    const { team } = useCsSelector(state => state.team);

    useEffect(() => {
        if (teamId) {
            dispatch(Actions.team.GetTeamById({ id: teamId! }));
        }
    }, [teamId])
    return (
        <div className="space-y-5 pb-20">

            <div className="flex items-center justify-between ">

                <h1 className="text-2xl font-bold flex items-center gap-3">
                    {teamId && <ArrowLeft
                        onClick={() => {
                            searchParams.delete("team");
                            setSearchParams(searchParams);
                            dispatch(TeamSlice.actions.removeTeamMembers({}));
                            dispatch(TeamSlice.actions.removeTeam({}));
                        }} />
                    }

                    {teamId ? team?.name : "Teams"}
                </h1>
                <NewTeam />
            </div>


            {teamId ?
                <>
                    <TeamGeneralInfo />
                    <TeamMembers />
                </>
                : <TeamList />
            }

        </div>
    )
}
