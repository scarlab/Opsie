import { useSearchParams } from 'react-router-dom';
import {
    Dialog,
    DialogBody,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
} from "@/components/cn/dialog"
import { SystemRoles } from '@/constants';
import { Actions, useCsDispatch, useCsSelector } from '@/cs-redux';
import { useState } from 'react';
import type { NewTeamPayload } from '@/types/team.type';
import { Button } from '@/components/cn/button';
import { Input } from '@/components/cn/input';
import { Label } from '@/components/cn/label';
import { toast } from 'sonner';
import { Plus } from 'lucide-react';


export default function NewTeam() {
    const dispatch = useCsDispatch();
    const [searchParams, setSearchParams] = useSearchParams();
    const isNewTeam = searchParams.get("new-team") === "true";

    const { authUser } = useCsSelector(state => state.auth);
    const [payload, setPayload] = useState<NewTeamPayload>({ name: '', description: '' });

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const { name, value } = e.target;
        setPayload(prev => ({ ...prev, [name]: value }));
    }

    async function onSave() {
        if (!payload.name) return toast.error("Team name is required");

        const res = await dispatch(Actions.team.CreateTeam(payload));
        if (res.meta.requestStatus === "fulfilled") {
            toast.success(res.payload.message);
            setPayload({ name: '', description: '' })

            searchParams.delete("new-team");
            setSearchParams(searchParams);
        }
        else if (res.meta.requestStatus === "rejected") {
            toast.error(res.payload.error);
        }
    }

    if (!authUser || authUser?.system_role === SystemRoles.Staff) return null;
    return (
        <div>
            <Button size={'sm'} variant={"outline"} onClick={() => setSearchParams({ "new-team": "true" })}><Plus />New Team</Button>

            <Dialog
                open={isNewTeam}
                onOpenChange={() => {
                    searchParams.delete("new-team");
                    setSearchParams(searchParams);
                    setPayload({ name: '', description: '' })
                }}
            >
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Create New Team...</DialogTitle>
                        <DialogDescription>
                            Enter name & description to create a new team.
                        </DialogDescription>
                    </DialogHeader>
                    <DialogBody className='flex flex-col gap-4 p-6'>
                        <div className='flex flex-col gap-1'>
                            <Label htmlFor='name'>Name</Label>
                            <Input value={payload.name} name='name' id='name' onChange={handleChange} placeholder='Team name' />
                        </div>

                        <div className='flex flex-col gap-1'>
                            <Label htmlFor='description'>Description</Label>
                            <Input value={payload.description} name='description' id='description' onChange={handleChange} placeholder='Description' />
                        </div>
                    </DialogBody>
                    <DialogFooter>
                        <Button onClick={onSave}>Create</Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>
        </div>
    )
}
