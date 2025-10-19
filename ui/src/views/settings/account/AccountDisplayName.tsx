import { Button } from "@/components/cn/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/cn/card";
import { Input } from "@/components/cn/input";
import { Label } from "@/components/cn/label";
import { useCsSelector } from "@/cs-redux";
import { useState } from "react";


export default function AccountDisplayName() {
    const { authUser } = useCsSelector(state => state.auth);
    const [name, setName] = useState<string>(authUser?.display_name ?? "");

    return (
        <Card>
            <CardHeader>
                <CardTitle>Display Name</CardTitle>
                <CardDescription>This will be displayed on your profile and in the workspace.
                </CardDescription>
                <CardContent className="mt-3">
                    <Label className="mb-2" htmlFor="display_name">Display Name</Label>
                    <Input value={name} onChange={e => setName(e.target.value)} id="display_name" type="text" name="display_name" placeholder="Display Name" />
                </CardContent>
                <CardFooter className="border-t flex justify-end">
                    <Button size={'sm'}>Save</Button>
                </CardFooter>
            </CardHeader>
        </Card>
    )
}
