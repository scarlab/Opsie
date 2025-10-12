import Config from "@/config";
import { Plus, Server } from "lucide-react";

export default function NodesOverview() {
    return (
        <div className="px-3 mt-14">
            <h1 className="text-2xl font-bold mb-2 flex items-center gap-3"><Server size={20} /> Nodes</h1>
            <div className="h-0.5 bg-muted" />

            <div className="grid grid-cols-5 gap-x-3 gap-y-5 mt-3">
                <div className="hover:bg-accent/20 transition-all border h-24 px-3 py-2 rounded flex flex-col justify-between">
                    <div>
                        <p className="font-bold text-lg">Mother Node</p>
                        <p className="text-xs leading-tight text-muted-foreground">Where the {Config.projectName} lives</p>
                    </div>

                    <div className="flex items-center justify-between">
                        <div className="text-sm text-muted-foreground flex items-center gap-2">
                            <div className="rounded-full w-2 aspect-square bg-green-500" />
                            <span>Online</span>
                        </div>
                        <div>
                            <p className="text-xs text-muted-foreground"><strong>Uptime:</strong> 39 Days</p>
                        </div>
                    </div>
                </div>



                <div className="hover:bg-accent/20 transition-all cursor-pointer border border-dashed rounded flex items-center justify-center h-24">
                    <Plus />
                    <span>Add New Node</span>
                </div>
            </div>
        </div>
    )
}
