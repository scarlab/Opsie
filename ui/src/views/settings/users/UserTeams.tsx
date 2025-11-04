import { Button } from "@/components/cn/button";
import { useCsSelector } from "@/cs-redux";
import { EllipsisVertical, Tent, X } from "lucide-react";
import { toast } from "sonner";

export default function UserTeams() {
    const { user_teams } = useCsSelector(state => state.team);



    return (
        <div>

            <div className="mt-3 space-y-">
                <div className="grid grid-cols-12 gap-2 pb-1 text-muted-foreground text-sm border-b">
                    <p className="col-span-1 text-center ps-1">###</p>
                    <p className="col-span-5 ps-3">Name</p>
                    <p className="col-span-5">Description</p>
                    <p className="col-span-1">Action</p>
                </div>

                {user_teams ?
                    user_teams.length === 0 ?
                        <div className="w-full h-96 flex flex-col justify-center items-center space-y-3">
                            <p className="text-muted-foreground">This user is not part of any team.</p>
                        </div>
                        :
                        user_teams.map((team, i) => (
                            <div key={i} className="grid grid-cols-12 hover:bg-accent/20 transition-all py-1.5 border-b px-3">
                                <div className="col-span-1">
                                    <div className="border bg-accent/20 w-8 aspect-square rounded-full grid place-items-center" >
                                        <Tent size={17} />
                                    </div>
                                </div>
                                <p className="col-span-5 font-medium">{team.name}</p>
                                <p className="col-span-5 text-muted-foreground text-sm">{team.description}</p>
                                <div className="col-span-1 w-full flex justify-end items-center">
                                    <button className="cursor-pointer hover:bg-accent p-1 rounded transition-all" onClick={() => toast.info("Done...")}><X size={18} color="red" /></button>
                                </div>
                            </div>
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
