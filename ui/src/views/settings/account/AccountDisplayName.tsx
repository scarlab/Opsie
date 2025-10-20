import { Button } from "@/components/cn/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/cn/card";
import { Input } from "@/components/cn/input";
import { Label } from "@/components/cn/label";
import { Actions, useCsDispatch, useCsSelector } from "@/cs-redux";
import AuthSlice from "@/cs-redux/slices/auth.slice";
import { useState } from "react";
import { toast } from "sonner";


export default function AccountDisplayName() {
    const dispatch = useCsDispatch();

    const { authUser } = useCsSelector(state => state.auth);
    const [display_name, setDisplayName] = useState<string>(authUser?.display_name ?? "");



    async function onSave() {
        if (display_name === authUser?.display_name) return

        if (!display_name) return toast.error("Enter your display name")

        const res = await dispatch(Actions.user.UpdateAccountDisplayName({ display_name }))
        if (res.payload.message) {
            toast.success(res.payload.message)
            dispatch(AuthSlice.actions.updateAuthUser(res.payload.auth_user))
        }
        else if (res.payload.error) {
            toast.error(res.payload.error)
        }

    }

    return (
        <Card>
            <CardHeader>
                <CardTitle>Display Name</CardTitle>
                <CardDescription>This will be displayed on your profile and in the workspace.
                </CardDescription>
                <CardContent className="mt-3">
                    <Label className="mb-2" htmlFor="display_name">Display Name</Label>
                    <Input value={display_name} onChange={e => setDisplayName(e.target.value)} id="display_name" type="text" name="display_name" placeholder="Display Name" />
                </CardContent>
                <CardFooter className="border-t flex justify-end">
                    <Button onClick={onSave} size={'sm'}>Save</Button>
                </CardFooter>
            </CardHeader>
        </Card>
    )
}
