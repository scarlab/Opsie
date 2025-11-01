import { Button } from "@/components/cn/button";
import { Card, CardAction, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/cn/card";
import { Input } from "@/components/cn/input";
import { Label } from "@/components/cn/label";
import { Clipboard, ClipboardCheck } from "lucide-react";
import { useEffect, useState } from "react";
import copy from 'copy-to-clipboard';
import { useCsSelector } from "@/cs-redux";
import { getInitials } from "@/lib/text";

export default function TeamGeneralInfo() {
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
    return (
        <Card>
            <CardHeader>
                <CardTitle>General Info</CardTitle>
                <CardDescription>This is general info of default team.</CardDescription>
                <CardAction className="bg-accent w-14 aspect-square rounded-full grid place-items-center">
                    <span className="font-black text-xl text-primary">{getInitials(team?.name!)}</span>
                </CardAction>
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
                <Button size={'sm'} >Save</Button>
            </CardFooter>
        </Card>
    )
}
