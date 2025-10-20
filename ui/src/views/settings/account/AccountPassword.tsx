import { Button } from "@/components/cn/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/cn/card";
import { InputPassword } from "@/components/cn/input-password";
import { Label } from "@/components/cn/label";
import { Actions, useCsDispatch } from "@/cs-redux";
import { useState } from "react";
import { toast } from "sonner";


export default function AccountPassword() {
    const dispatch = useCsDispatch();

    const [password, setPassword] = useState<string>("");
    const [newPassword, setNewPassword] = useState<string>("");
    const [confirmPassword, setConfirmPassword] = useState<string>("");

    async function onSave() {
        if (!password) return toast.error("Enter your current password")
        else if (!newPassword) return toast.error("Enter new password")
        else if (!confirmPassword) return toast.error("Confirm new password")
        else if (newPassword !== confirmPassword) return toast.error("New password doesn't match")


        const res = await dispatch(Actions.user.UpdateAccountPassword({ password, new_password: newPassword }))
        if (res.payload.message) {
            toast.success(res.payload.message)
            setPassword("")
            setNewPassword("")
            setConfirmPassword("")
        } else if (res.payload.error) {
            toast.error(res.payload.error)
        }
    }
    return (
        <Card>
            <CardHeader>
                <CardTitle>Password</CardTitle>
                <CardDescription>Update your account password here.
                </CardDescription>
                <CardContent className="mt-3 space-y-5">
                    <div>
                        <Label className="mb-2" htmlFor="current-password">Current Password</Label>
                        <InputPassword value={password} onChange={e => setPassword(e.target.value)} id="current-password" name="current-password" placeholder="* * * * * * * * *" />
                    </div>
                    <div>
                        <Label className="mb-2" htmlFor="new-password">New Password</Label>
                        <InputPassword value={newPassword} onChange={e => setNewPassword(e.target.value)} id="new-password" name="new-password" placeholder="* * * * * * * * *" />
                    </div>
                    <div>
                        <Label className="mb-2" htmlFor="confirm-password">Re-enter New Password</Label>
                        <InputPassword value={confirmPassword} onChange={e => setConfirmPassword(e.target.value)} id="confirm-password" name="confirm-password" placeholder="* * * * * * * * *" />
                    </div>
                </CardContent>
                <CardFooter className="border-t flex justify-end">
                    <Button onClick={onSave} size={'sm'}>Save</Button>
                </CardFooter>
            </CardHeader>
        </Card>
    )
}
