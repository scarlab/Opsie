import { Button } from "@/components/cn/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/cn/card";
import { Input } from "@/components/cn/input";
import { Label } from "@/components/cn/label";
import { Clipboard, ClipboardCheck } from "lucide-react";
import { useEffect, useState } from "react";
import copy from 'copy-to-clipboard';
import { Actions, useCsDispatch, useCsSelector } from "@/cs-redux";
import { toast } from "sonner";

export default function TeamGeneralInfo() {
    const dispatch = useCsDispatch();
    const [teamId, setTeamId] = useState<string>("");
    const [name, setName] = useState<string>("");
    const [description, setDescription] = useState<string>("");
    const [copied, setCopied] = useState<boolean>(false);

    const { team } = useCsSelector(state => state.team);

    useEffect(() => {
        if (team) {
            setName(team.name);
            setDescription(team?.description! || '');
            setTeamId(team.id);
        }
    }, [team])

    function onCopy() {
        try {
            copy(team?.id || '');
            setCopied(true);
            setTimeout(() => setCopied(false), 2000);
        } catch (err) {
            console.error('Failed to copy: ', err);
        }
    }

    async function onSave() {
        if (!teamId || !team) return;
        if (team.name === name && team.description === description) return;

        if (!name) return toast.error("Team name is required");

        const res = await dispatch(Actions.team.Update({ id: teamId, data: { name, description } }));
        if (res.meta.requestStatus === "fulfilled") {
            toast.success(res.payload.message);
        }
        else if (res.meta.requestStatus === "rejected") {
            toast.error(res.payload.error);
        }
    }

    return (
        <Card>
            <CardHeader>
                <CardTitle>General Info</CardTitle>
                <CardDescription>This is general info of default team.</CardDescription>
            </CardHeader>

            <CardContent className="">
                <div className="grid grid-cols-2 gap-x-3 w-full">
                    <div>
                        <Label className="mb-2" htmlFor="name">Name</Label>
                        <Input value={name} onChange={e => setName(e.target.value)} id="name" type="text" name="name" placeholder="Team name" />
                    </div>
                    <div className="relative">
                        <Label className="mb-2" htmlFor="team_id">Team ID</Label>
                        <Input value={teamId} disabled id="team_id" type="text" name="team_id" />
                        <Button onClick={() => { !copied && onCopy() }} className="absolute bottom-px right-0" size={'icon-sm'} variant={'ghost'}>
                            {copied ?
                                <ClipboardCheck className="text-green-500" />
                                :
                                <Clipboard />
                            }
                        </Button>
                    </div>
                </div>
                <div className="mt-3">
                    <Label className="mb-2" htmlFor="name">description</Label>
                    <Input value={description} onChange={e => setDescription(e.target.value)} id="description" type="text" name="description" placeholder="Team description" />
                </div>
            </CardContent>
            <CardFooter className="border-t flex justify-end items-center">
                <Button onClick={onSave} size={'sm'} >Save</Button>
            </CardFooter>
        </Card>
    )
}
