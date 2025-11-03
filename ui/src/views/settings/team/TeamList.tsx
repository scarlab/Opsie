import { Actions, useCsDispatch, useCsSelector } from "@/cs-redux";
import { Tent } from "lucide-react";
import { useEffect } from "react";
import { Link } from "react-router-dom";

export default function TeamList() {
    const dispatch = useCsDispatch();
    const { teams } = useCsSelector(state => state.team);


    useEffect(() => {
        dispatch(Actions.team.GetAllTeams());
    }, [])



    return (
        <div>

            <div className="mt-3 space-y-">
                {teams ?
                    teams.map((team, i) => (
                        <Link to={`/settings/team?team=${team.id}`} key={i} className="grid grid-cols-12 hover:bg-accent/20 transition-all py-1.5 border-b px-3">

                            <div className="col-span-1">
                                <div className="border bg-accent/20 w-8 aspect-square rounded-full grid place-items-center" >
                                    <Tent size={17} />
                                </div>
                            </div>
                            <p className="col-span-5 font-medium">{team.name}</p>
                            <p className="col-span-5 text-muted-foreground text-sm">{team.description}</p>
                            {/* <div className="col-span-1 w-full flex justify-end items-center">
                                <Button size={"icon-sm"} variant={"ghost"}><EllipsisVertical /></Button>
                            </div> */}
                        </Link>
                    ))
                    :
                    <div>
                        Loading...
                    </div>
                }
            </div>
        </div>
    )
}
