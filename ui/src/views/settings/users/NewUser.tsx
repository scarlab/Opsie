import {
    Dialog,
    DialogBody,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
} from "@/components/cn/dialog"
import { SystemRoles, type SystemRolesType } from '@/constants';
import { Actions, useCsDispatch, useCsSelector } from '@/cs-redux';
import { useState } from 'react';
import type { NewUserPayload } from '@/types/user.type';
import { Button } from '@/components/cn/button';
import { Input } from '@/components/cn/input';
import { Label } from '@/components/cn/label';
import { toast } from 'sonner';
import { Key, Plus } from 'lucide-react';
import SelectAvatar from "@/components/utils/SelectAvatar";
import { RadioGroup, RadioGroupItem } from "@/components/cn/radio-group";
import { cn } from "@/lib/utils";
import { InputPassword } from "@/components/cn/input-password";
import { generatePassword, getRandomValue } from "@/lib/random";
import CsImage from "@/constants/image";

const default_payload: NewUserPayload = {
    display_name: '',
    email: '',
    password: '',
    system_role: 'staff',
    avatar: getRandomValue(CsImage.avatar),
}

export default function NewUser() {
    const dispatch = useCsDispatch();

    const [isNewUser, setIsNewUser] = useState<boolean>(false)

    const { authUser } = useCsSelector(state => state.auth);
    const [payload, setPayload] = useState<NewUserPayload>(default_payload);

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const { name, value } = e.target;
        setPayload(prev => ({ ...prev, [name]: value }));
    }

    async function onSave() {

        if (!payload.display_name) return toast.error("Display name is required");
        else if (!payload.email) return toast.error("Enter valid email");
        else if (!payload.password) return toast.error("Enter or generate password");

        const res = await dispatch(Actions.user.CreateUser(payload));
        if (res.meta.requestStatus === "fulfilled") {
            toast.success(res.payload.message);
            setPayload(default_payload);

            setIsNewUser(false)
        }
        else if (res.meta.requestStatus === "rejected") {
            toast.error(res.payload.error);
        }
    }

    function setSystemRole(role: SystemRolesType) {
        setPayload(prev => ({ ...prev, system_role: role }));
    }

    if (!authUser || authUser?.system_role === SystemRoles.Staff) return null;
    return (
        <div>
            <Button size={'sm'} variant={"outline"} onClick={() => setIsNewUser(true)}><Plus />New User</Button>

            <Dialog
                open={isNewUser}
                onOpenChange={() => {
                    setIsNewUser(false);
                    setPayload({ ...default_payload, avatar: payload.avatar });
                }}
            >
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Create New User...</DialogTitle>
                        <DialogDescription>
                            Enter user details to create a new user.
                        </DialogDescription>
                    </DialogHeader>
                    <DialogBody className='flex flex-col gap-5 '>
                        <div className='flex justify-center items-center w-22 aspect-square'>
                            <SelectAvatar avatar={payload.avatar} setAvatar={a => setPayload(prev => ({ ...prev, avatar: a }))} />
                        </div>


                        <div className='flex flex-col gap-1 grow'>
                            <Label htmlFor='display_name'>Display name</Label>
                            <Input value={payload.display_name} name='display_name' id='display_name' type="text" onChange={handleChange} placeholder='Display name' />
                        </div>


                        <div className='flex flex-col gap-1'>
                            <Label htmlFor='email'>Email</Label>
                            <Input value={payload.email} name='email' id='email' type="email" onChange={handleChange} placeholder='email@address.in' />
                        </div>


                        <div className='flex flex-col gap-1'>
                            <Label htmlFor='password'>Password</Label>
                            <div className="flex items-center justify-between gap-2">
                                <div className="grow">
                                    <InputPassword value={payload.password} name='password' id='password' onChange={handleChange} placeholder='* * * * * * * * *' />
                                </div>
                                <Button onClick={() => setPayload({ ...payload, password: generatePassword() })} title="Generate & fill secure password" size={"icon"} variant={"outline"}><Key /></Button>
                            </div>
                        </div>
                        <div className='flex flex-col gap-1 '>
                            <Label className="mb-2">System Role</Label>
                            <RadioGroup onValueChange={r => { setSystemRole(r as SystemRolesType) }} className="flex items-center gap-3" defaultValue={SystemRoles.Staff}>
                                {
                                    Object.values(SystemRoles).map((role, i) => (
                                        <div key={`sys_r-${i}`} className="flex items-center">
                                            <RadioGroupItem className="cursor-pointer" value={role} id={role} />
                                            <Label htmlFor={role} className={cn("capitalize cursor-pointer ps-2", role === payload.system_role && "text-primary")}>{role}</Label>
                                        </div>
                                    ))
                                }
                            </RadioGroup>
                        </div>

                    </DialogBody>
                    <DialogFooter>
                        <Button onClick={onSave}>Add User</Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>
        </div>
    )
}
