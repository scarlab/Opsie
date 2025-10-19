import { Button } from "@/components/cn/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/cn/card";
import { Input } from "@/components/cn/input";
import { Label } from "@/components/cn/label";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/cn/dialog"
import {
    InputOTP,
    InputOTPGroup,
    InputOTPSlot,
} from "@/components/cn/input-otp"
import Config from "@/config";
import { CheckCheck, Ellipsis, Mail, Send } from "lucide-react";
import { useState } from "react";
import { DialogClose } from "@radix-ui/react-dialog";
import { useCsSelector } from "@/cs-redux";


export default function AccountEmail() {
    const { authUser } = useCsSelector(state => state.auth);

    const [email, setEmail] = useState<string>(authUser?.email ?? "");
    const [otp, setOtp] = useState<string>("");
    const [otpSent, setOtpSent] = useState<boolean>(false);

    async function sendOtp() {
        console.log(email);
        setOtpSent(true);
    }

    async function verifyOtp() {
        console.log(otp);
        setOtp("");
        setOtpSent(false);
    }



    return (
        <Card>
            <CardHeader>
                <CardTitle>Email</CardTitle>
                <CardDescription>Enter the email address you want to use to log in with {Config.projectName}. Your primary email will be used for account-related notifications.</CardDescription>
                <CardContent className="mt-3">
                    <div className="bg-accent px-3 py-2 rounded flex items-center justify-between ">
                        <div>
                            <Mail size={15} className="inline me-2 text-muted-foreground" />
                            <span>{email}</span>
                            <span className="border border-green-300 text-green-500 text-xs rounded-2xl px-1.5 pb-px ms-3">
                                <CheckCheck size={14} className="pb-0.5 inline" />
                                Verified
                            </span>
                        </div>


                        <Dialog>
                            <DialogTrigger asChild>
                                <Button size={'icon-sm'} variant={'ghost'}><Ellipsis /></Button>
                            </DialogTrigger>
                            <DialogContent >
                                <DialogHeader>
                                    <DialogTitle>{otpSent ? "Verify OTP" : "Update Email Address"}</DialogTitle>
                                    <DialogDescription>
                                        {otpSent ? "Enter the OTP sent to your new email address to verify & complete the process." : "Enter the new email address. This action will send a verification email to the new address."}
                                    </DialogDescription>
                                </DialogHeader>
                                {otpSent ?
                                    <div>
                                        <Label className="mb-2" htmlFor="otp">OTP</Label>

                                        <InputOTP value={otp} onChange={data => setOtp(data)} maxLength={6} id="otp">
                                            <InputOTPGroup>
                                                <InputOTPSlot index={0} />
                                                <InputOTPSlot index={1} />
                                                <InputOTPSlot index={2} />
                                                <InputOTPSlot index={3} />
                                                <InputOTPSlot index={4} />
                                                <InputOTPSlot index={5} />
                                            </InputOTPGroup>
                                        </InputOTP>
                                    </div>
                                    :
                                    <div>
                                        <Label className="mb-2" htmlFor="email">Email</Label>
                                        <Input value={email} onChange={e => setEmail(e.target.value)} id="email" type="email" name="email" placeholder="Email Address" />
                                    </div>
                                }

                                <DialogFooter>
                                    {otpSent ?
                                        <DialogClose asChild>
                                            <Button onClick={verifyOtp} size={'sm'}>Verify OTP</Button>
                                        </DialogClose>
                                        :
                                        <Button onClick={sendOtp} size={'sm'}>Send OTP <Send /></Button>
                                    }
                                </DialogFooter>
                            </DialogContent>
                        </Dialog>
                    </div>
                </CardContent>
            </CardHeader>
        </Card>
    )
}
