
import { Button } from "@/components/cn/button";
import { Card, CardAction, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/cn/card";
import { Actions, useCsDispatch, useCsSelector } from "@/cs-redux";
import { formattedDate } from "@/lib/time";
import { EllipsisVertical, Plus } from "lucide-react";
import { useEffect } from "react";


export default function TeamMembers() {
    const { team_members, user_default_team } = useCsSelector(state => state.team);
    const dispatch = useCsDispatch();

    useEffect(() => {
        if (user_default_team) {
            dispatch(Actions.team.GetTeamMembers({ id: user_default_team.id }))
        }
    }, [user_default_team, dispatch])
    return (
        <Card>
            <CardHeader>
                <CardTitle>Team Members ({team_members?.length})</CardTitle>
                <CardDescription>This is general info of default team.</CardDescription>
                <CardAction >
                    <Button size={"sm"} variant={"outline"} ><Plus />Add Member</Button>
                </CardAction>
            </CardHeader>

            <CardContent className="">
                <div className="grid grid-cols-12 gap-2 text-muted-foreground text-sm border-b">
                    <p className="col-span-1"></p>
                    <p className="col-span-4">Name</p>
                    <p className="col-span-4">Email</p>
                    <p className="col-span-2 text-center">System Role</p>
                </div>
                {team_members && team_members.map((member, i) =>
                    <div key={i} className="grid grid-cols-12 gap-2 items-center justify-between  border-b py-1">

                        <div className="col-span-1  flex items-center justify-center">
                            <img className="rounded-full w-9" width={200} height={200} src={member.avatar} alt={member.display_name} />
                        </div>

                        <div className="col-span-4  ">
                            <p className="text-lg font-medium leading-5">{member.display_name}</p>
                            <p className="text-xs text-muted-foreground leading-4">Member since: {formattedDate(member.joined_at)}</p>
                        </div>

                        <div className="col-span-4  ">
                            <span className="text-sm">{member.email}</span>
                        </div>


                        <div className="col-span-2 h-full flex items-center justify-center  ">
                            <p className="capitalize">{member.system_role}</p>
                        </div>


                        <div className="col-span-1 aspect-square flex items-center justify-end">
                            <Button variant={"ghost"} size={"icon-sm"}><EllipsisVertical /></Button>
                        </div>
                    </div>
                )}
            </CardContent>
        </Card>
    )
}
