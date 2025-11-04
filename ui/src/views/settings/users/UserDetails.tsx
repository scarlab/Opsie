import { Card, CardContent } from "@/components/cn/card";
import Role from "@/components/utils/Role";
import { useCsSelector } from "@/cs-redux";
import { Check, EllipsisVertical, X } from "lucide-react";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "@/components/cn/dropdown-menu"

export default function UserDetails() {
    const { user } = useCsSelector(state => state.user);
    const { user_teams } = useCsSelector(state => state.team);

    return (
        <div>
            <Card className="relative">
                <CardContent >
                    <div className="grid grid-cols-2  items-start">
                        <div>
                            <div className="w-24 rounded-full aspect-square bg-accent cursor-pointer">
                                <img src={user?.avatar} />
                            </div>
                            <div className="flex items-center gap-2">
                                <h3 className="text-xl font-semibold mb-1"> {user?.display_name}</h3>
                            </div>

                            <p className="text-foreground/80 leading-4">{user?.email}</p>
                        </div>

                        <div className="border-l h-full px-3">
                            <div className="text-sm text-left grid grid-cols-3 border-b px-2 py-1 items-center">
                                <span>System Role</span>
                                <span className="col-span-2 border-l px-2"><Role sysRole={user?.system_role!} /></span>
                            </div>

                            <div className="text-sm text-left grid grid-cols-3 border-b px-2  py-1 items-center">
                                <span>Status</span>
                                <span className="col-span-2 border-l px-2 ">{user?.is_active ? <Check size={19} className="text-green-500" /> : <X size={19} color='red' />}</span>
                            </div>

                            <div className="text-sm text-left grid grid-cols-3 border-b px-2  py-1 items-center">
                                <span>Teams</span>
                                <span className="col-span-2 border-l px-2 ">{user_teams?.length}</span>
                            </div>

                        </div>
                    </div>


                </CardContent>
                <DropdownMenu >
                    <DropdownMenuTrigger className="cursor-pointer absolute top-3 right-2"><EllipsisVertical size={18} /></DropdownMenuTrigger>
                    <DropdownMenuContent onCloseAutoFocus={(e) => e.preventDefault()}>
                        <DropdownMenuItem>Update</DropdownMenuItem>
                        <DropdownMenuItem variant="destructive">Delete</DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </Card>
        </div>
    )
}
