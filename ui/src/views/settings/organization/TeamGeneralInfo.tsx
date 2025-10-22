import { Button } from "@/components/cn/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/cn/card";
import { Input } from "@/components/cn/input";
import { Label } from "@/components/cn/label";
import { Clipboard, ClipboardCheck } from "lucide-react";
import { useState } from "react";
import copy from 'copy-to-clipboard';
import { Spinner } from "@/components/cn/spinner";

export default function TeamGeneralInfo() {
    const [name, setName] = useState<string>("ScarLab");
    const [copied, setCopied] = useState<boolean>(false);

    const team_id = "red-dragon-scarlab-7628345";

    function onCopy() {
        try {
            copy(team_id);
            setCopied(true);
            setTimeout(() => setCopied(false), 2000);
        } catch (err) {
            console.error('Failed to copy: ', err);
        }
    }
    return (
        <Card>
            <CardHeader>
                <CardTitle>General Information</CardTitle>
                <CardDescription>This will be displayed on team section.
                </CardDescription>
                <CardContent className="mt-3">
                    <div className="grid grid-cols-2 gap-x-3 w-full">
                        <div>
                            <Label className="mb-2" htmlFor="name">Team Name</Label>
                            <Input value={name} onChange={e => setName(e.target.value)} id="name" type="text" name="name" placeholder="Team Name" />
                        </div>
                        <div className="relative">
                            <Label className="mb-2" htmlFor="team_id">Team ID</Label>
                            <Input value={team_id} disabled id="team_id" type="text" name="team_id" />
                            <Button onClick={() => { !copied && onCopy() }} className="absolute bottom-px right-0" size={'icon-sm'} variant={'ghost'}>
                                {copied ?
                                    <ClipboardCheck className="text-green-500" />
                                    :
                                    <Clipboard />
                                }
                            </Button>
                        </div>
                    </div>
                </CardContent>
                <CardFooter className="border-t flex justify-end">
                    <Button size={'sm'} >Save</Button>
                    <Button size={'sm'} variant={'outline'}> <Spinner />Loading...</Button>
                </CardFooter>
            </CardHeader>
        </Card>
    )
}
