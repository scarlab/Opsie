

import { LayoutDashboardIcon, LogOut, UserCog } from 'lucide-react'
import { Button } from "@/components/cn/button"
import {
    Command,
    CommandItem,
    CommandList,
    CommandSeparator,
    CommandShortcut,
} from "@/components/cn/command"

import { useState } from 'react';
import { Popover, PopoverContent, PopoverTrigger } from '@/components/cn/popover';
import { Link, useNavigate } from 'react-router-dom';
import { Theme } from '../theme';
import { Actions, useCsDispatch, useCsSelector } from '@/cs-redux';
import { toast } from 'sonner';


export default function NavUserMenu() {
    const navigate = useNavigate()
    const dispatch = useCsDispatch();
    const { authUser } = useCsSelector(state => state.auth);

    const [open, setOpen] = useState(false);



    async function onLogout() {
        const res = await dispatch(Actions.auth.Logout())

        if (res.payload.message) {
            toast.success(res.payload.message)
            navigate("/auth/login", { replace: true })
        }
        else if (res.payload.error) {
            toast.error(res.payload.message)
        }

        setOpen(false)
    }
    return (
        <>
            <Button size={'icon'} variant={'ghost'} className='overflow-hidden !rounded-full' onClick={() => { setOpen(true) }}>
                <img src={authUser?.avatar} alt={authUser?.display_name} />
            </Button>

            <Popover open={open} onOpenChange={() => { setOpen(false) }}>
                <PopoverTrigger>

                </PopoverTrigger>
                <PopoverContent className="w-60">
                    <Command>

                        <div className='mb-3'>
                            <div className="flex items-center gap-2 px-1 py-1.5 text-left text-sm">
                                <Button size={'icon'} variant={'ghost'} className='overflow-hidden !rounded-full' onClick={() => { setOpen(true) }}>
                                    <img src={authUser?.avatar} alt={authUser?.display_name} />
                                </Button>
                                <div className="grid flex-1 text-left text-sm leading-tight">
                                    <span className="truncate font-semibold">{authUser?.display_name}</span>
                                    <span className="truncate text-xs">{authUser?.email}</span>
                                </div>
                            </div>
                        </div>
                        <CommandList>
                            <Link to={`/`} onClick={() => setOpen(false)}>
                                <CommandItem className='cursor-pointer'>
                                    Overview
                                    <CommandShortcut><LayoutDashboardIcon /></CommandShortcut>
                                </CommandItem>
                            </Link>
                            <Link to={`/settings/account`} onClick={() => setOpen(false)}>
                                <CommandItem className='cursor-pointer'>
                                    Account Settings
                                    <CommandShortcut><UserCog /></CommandShortcut>
                                </CommandItem>
                            </Link>
                            {/* <Link to={`/settings`} onClick={() => setOpen(false)}>
                                <CommandItem className='cursor-pointer'>
                                    Settings
                                    <CommandShortcut><Settings /></CommandShortcut>
                                </CommandItem>
                            </Link> */}

                            <CommandSeparator className='my-2' />

                            <div className="text-sm px-2 py-1.5 flex items-center justify-between">
                                <span className="grow">Theme</span>
                                <Theme />
                            </div>


                            <div onClick={onLogout} >
                                <CommandItem className='!text-destructive hover:!text-destructive cursor-pointer'>
                                    Logout <CommandShortcut><LogOut className='!text-destructive hover:!text-destructive' /></CommandShortcut>
                                </CommandItem>
                            </div>

                            <CommandSeparator className='my-2' />
                            <Button onClick={() => { setOpen(false) }} className='mt-3 w-full' size={'sm'}>Upgrade to Pro</Button>
                        </CommandList>
                    </Command>
                </PopoverContent>
            </Popover>
        </>
    )
}