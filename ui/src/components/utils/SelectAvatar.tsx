import CsImage from "@/constants/image";
import { Popover, PopoverContent, PopoverTrigger } from "../cn/popover";
import { useState } from "react";

export default function SelectAvatar({ avatar, setAvatar }: { avatar: string, setAvatar: (a: string) => void }) {
    const [open, setOpen] = useState<boolean>(false);

    function onOpenChange(isOpen: boolean) {
        setOpen(isOpen);


    }
    return (
        <>
            <button onClick={() => setOpen(true)} className="w-full rounded-full aspect-square bg-accent cursor-pointer">
                <img src={avatar} />
            </button>

            <Popover open={open} onOpenChange={onOpenChange} >
                <PopoverTrigger></PopoverTrigger>
                <PopoverContent>
                    <div className="grid grid-cols-5 gap-3 p-3 ">
                        {Object.values(CsImage.avatar).map((avatar, i) => (
                            <button onClick={() => { setAvatar(avatar); setOpen(false) }} key={i} className="rounded-full aspect-square bg-accent cursor-pointer">
                                <img width={100} height={100} src={avatar} />
                            </button>
                        ))}
                    </div>
                </PopoverContent>
            </Popover>

        </>
    )
}
