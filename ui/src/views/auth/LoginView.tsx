import { Button } from "@/components/cn/button";
import { Input } from "@/components/cn/input";
import { Label } from "@/components/cn/label";
import { InputPassword } from "@/components/cn/input-password";
import Config from "@/config";
import { Lock } from "lucide-react";
import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import type { LoginPayload } from "@/types/auth.type";
import { Actions, useCsDispatch } from "@/cs-redux";
import { toast } from "sonner";

export default function LoginView() {
    const dispatch = useCsDispatch();
    const navigate = useNavigate()


    const [payload, setPayload] = useState<LoginPayload>({
        email: '',
        password: '',
    });

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const { name, value } = e.target;
        setPayload(prev => ({
            ...prev,
            [name]: value
        }));
    }


    async function onLogin() {
        if (!payload.email) return toast.error("Enter Email")
        else if (!payload.password) return toast.error("Enter Password")

        const res = await dispatch(Actions.auth.Login({ email: payload.email, password: payload.password }))

        if (res.payload.message) {
            toast.success(res.payload.message)
            navigate("/", { replace: true })

        }
        else if (res.payload.error) {
            toast.error(res.payload.error)
        }
    }

    return (
        <div className="grid place-items-center min-h-svh">
            <div className=" max-w-sm w-full px-3 py-7 space-y-10">
                <div className="flex items-center justify-center text-muted-foreground">
                    <Lock size={50} />
                </div>
                <div className="text-center">
                    <h1 className="text-primary text-2xl font-black">Log in to {Config.projectName}</h1>
                    <p className="text-muted-foreground">Secure access to your {Config.projectName} account.</p>
                </div>


                <div className="max-w-sm mx-auto flex flex-col gap-y-5">
                    <div className="space-y-2">
                        <Label>Email</Label>
                        <Input name="email" value={payload.email} onChange={handleChange} type="email" placeholder="Email" />
                    </div>

                    <div>
                        <div className="space-y-2">
                            <Label>Password</Label>
                            <InputPassword name="password" value={payload.password} onChange={handleChange} placeholder="Password" />
                        </div>
                        <div className="text-right mt-1">
                            <Link className="text-sm text-blue-500 font-bold text" to={'/auth/forgot-password'}>Forgot Password</Link>
                        </div>
                    </div>
                </div>


                <div >
                    <Button onClick={onLogin} size={'sm'} className="w-full">
                        Log In
                    </Button>
                </div>
            </div>
        </div>
    )
}
