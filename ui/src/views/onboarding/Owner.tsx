import { Button } from "@/components/cn/button";
import { Input } from "@/components/cn/input";
import { Label } from "@/components/cn/label";
import { PasswordInput } from "@/components/cn/password-input";
import Image from "@/constants/image";
import type { NewOwnerPayload } from "@/types/user";
import { useState } from "react";
import { motion } from "framer-motion"

export default function Owner({ next }: { next?: () => void }) {

    const [payload, setPayload] = useState<NewOwnerPayload>({
        name: '',
        email: '',
        password: '',
        confirmPassword: ''
    });

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const { name, value } = e.target;
        setPayload(prev => ({
            ...prev,
            [name]: value
        }));
    }

    async function onSave() {
        console.log(payload);
        next && next();
    }

    return (
        <div className="flex flex-col h-full justify-between">
            <div className="">
                <motion.h1
                    className="text-center font-black text-3xl pt-5"
                    initial={{ opacity: 0, y: -20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.6 }}
                >
                    Create Owner Account
                </motion.h1>
                <motion.p
                    className="text-center text-muted-foreground mt-3"
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ delay: 0.3 }}
                >
                    Let&prime;s start by creating your owner account.<br />
                    It gives you access to everything your team will manage.
                </motion.p>
            </div>


            <div className="grid grid-cols-2 gap-5 mt-10 grow ">
                <div className=" rounded transform scale-x-[-1]">
                    <motion.img
                        className=""
                        width={400}
                        height={400}
                        src={Image.workspace}
                        initial={{ scale: 0.8, opacity: 0 }}
                        animate={{ scale: 1, opacity: 1 }}
                        transition={{ duration: 0.3 }}
                        alt=""
                    />
                </div>
                <div className=" w-full ">
                    <motion.div
                        className="max-w-sm mx-auto flex flex-col gap-y-5"
                        initial={{ opacity: 0, x: 20 }}
                        animate={{ opacity: 1, x: 0 }}
                        transition={{ duration: 0.6 }}
                    >
                        <div className="space-y-2">
                            <Label>Name</Label>
                            <Input name="name" value={payload.name} onChange={handleChange} type="text" placeholder="Full Name" />
                        </div>

                        <div className="space-y-2">
                            <Label>Email</Label>
                            <Input name="email" value={payload.email} onChange={handleChange} type="email" placeholder="Email" />
                        </div>

                        <div className="space-y-2">
                            <Label>Password</Label>
                            <PasswordInput name="password" value={payload.password} onChange={handleChange} placeholder="Password" />
                        </div>
                        <div className="space-y-2">
                            <Label>Confirm Password</Label>
                            <PasswordInput name="confirmPassword" value={payload.confirmPassword} onChange={handleChange} placeholder="Re-enter Password" />
                        </div>
                    </motion.div>
                </div>
            </div>

            <div className="flex items-center justify-end pb-3 px-3 gap-x-3">
                {/* <Button onClick={next} variant={'outline'} size={'sm'}>
                    Cancle
                </Button> */}
                <motion.div
                    initial={{ scale: 0.8, opacity: 0 }}
                    animate={{ scale: 1, opacity: 1 }}
                    transition={{ delay: 0.5, type: "spring" }}
                >
                    <Button onClick={onSave} size={'sm'}>
                        Save & Continue
                    </Button>
                </motion.div>
            </div>
        </div>
    )
}
