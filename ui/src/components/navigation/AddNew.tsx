import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/cn/dropdown-menu"
import { Button } from "../cn/button"
import { Database, Globe, PanelTop, Plus, SendToBack, Shapes, Timer } from "lucide-react"



export default function AddNew() {
    return (
        <div>
            <DropdownMenu >
                <DropdownMenuTrigger asChild >
                    <Button variant={'outline'} >Add New <Plus /></Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-44" onCloseAutoFocus={e => e.preventDefault()}>
                    <DropdownMenuItem><PanelTop /> Static Site</DropdownMenuItem>
                    <DropdownMenuItem><Globe /> Web Service</DropdownMenuItem>
                    <DropdownMenuItem><SendToBack /> Worker</DropdownMenuItem>
                    <DropdownMenuItem><Timer /> Corn Job</DropdownMenuItem>
                    <DropdownMenuSeparator />
                    <DropdownMenuItem><Database /> Database</DropdownMenuItem>
                    <DropdownMenuItem><Shapes /> Peoject</DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    )
}
