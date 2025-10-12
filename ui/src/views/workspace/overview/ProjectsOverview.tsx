import { Plus, Shapes } from "lucide-react";

export default function ProjectsOverview() {
    return (
        <div className="px-3 mt-14">
            <h1 className="text-2xl font-bold mb-2 flex items-center gap-3"><Shapes size={20} />  Projects</h1>
            <div className="h-0.5 bg-muted" />

            <div className="grid grid-cols-5 gap-x-3 gap-y-5 mt-3">
                <div className="hover:bg-accent/20 transition-all border h-24 px-3 py-2 rounded flex flex-col justify-between">
                    <div>
                        <p className="font-bold text-lg">Bajra</p>
                        <p className="text-xs leading-tight text-muted-foreground">Intelligent drone swarm for Border Security...</p>
                    </div>

                    <div>
                        <p className="text-sm text-muted-foreground">3/3 Active Resources</p>
                    </div>
                </div>

                <div className="hover:bg-accent/20 transition-all border h-24 px-3 py-2 rounded flex flex-col justify-between">
                    <div>
                        <p className="font-bold text-lg">NodX</p>
                        <p className="text-xs leading-tight text-muted-foreground">Drone Swarm Communication</p>
                    </div>

                    <div>
                        <p className="text-sm text-muted-foreground">3/4 Active Resources</p>
                    </div>
                </div>


                <div className="hover:bg-accent/20 transition-all cursor-pointer border border-dashed rounded flex items-center justify-center h-24">
                    <Plus />
                    <span>Create New Project</span>
                </div>
            </div>
        </div>
    )
}
