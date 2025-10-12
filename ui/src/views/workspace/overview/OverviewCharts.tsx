import { cn } from "@/lib/utils"
import { Cpu, Database, HardDrive, Server, type LucideIcon } from "lucide-react"
import { useEffect, useRef, useState, type ReactNode } from "react"


// type MetricType = {
//     type: "node" | "cpu" | "memory" | "storage";
//     values: number[];
//     unit: string;
//     title: string;
//     icon: LucideIcon;
//     cardBgColor: string;
//     progressFillColor: string;
// };


type CardDataType = {
    children: ReactNode;
    title: string;
    icon: LucideIcon;
};

// Simulated data
const data = {
    node: {
        online: 16,
        total: 23,
        unit: "nodes"
    },
    cpu: {
        used: 69,
        total: 100,
        unit: "%",
        details: {
            cores: 128,
            threads: 256,
            speed_ghz: 3.5,
        }
    },
    memory: {
        used: 121,
        total: 164,
        unit: "GB"
    },
    storage: {
        used: 161,
        total: 256,
        unit: "TB"
    },
};


export default function OverviewCharts() {

    const cardClassName = "border rounded dark:bg-accent/20 bg-accent/30 aspect-video grid grid-cols-5"

    return (
        <div className="grid grid-cols-4 gap-5 px-3 py-2 ">
            <div className={cn(cardClassName,)}>
                <ProgressBar
                    type="node"
                    values={[data.node.online, data.node.total]}
                    unit={data.node.unit}
                    className="bg-gradient-to-tr from-pink-500 via-red-500 to-orange-500"
                />
                <CardData title="Nodes" icon={Server}>
                    <table className="w-full text-xs text-muted-foreground">
                        <tbody className="divide-y divide-border">
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Total</td>
                                <td className="py-1 text-right">{data.node.total} nodes</td>
                            </tr>
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Online</td>
                                <td className="py-1 text-right text-green-500">{data.node.online} nodes</td>
                            </tr>
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Offline</td>
                                <td className="py-1 text-right text-red-400">
                                    {data.node.total - data.node.online} nodes
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </CardData>
            </div>


            <div className={cn(cardClassName,)}>
                <ProgressBar
                    type="cpu"
                    values={[data.cpu.used, data.cpu.total]}
                    unit={data.cpu.unit}
                    className="bg-gradient-to-tr from-cyan-500 via-blue-500 to-indigo-500"
                />
                <CardData title="CPU" icon={Cpu}>
                    <table className="w-full text-xs text-muted-foreground">
                        <tbody className="divide-y divide-border">
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Utilized</td>
                                <td className="py-1 text-right">{data.cpu.used} {data.cpu.unit}</td>
                            </tr>
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">CPU Cores</td>
                                <td className="py-1 text-right">{data.cpu.details.cores}</td>
                            </tr>
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Threads</td>
                                <td className="py-1 text-right">{data.cpu.details.threads}</td>
                            </tr>
                        </tbody>
                    </table>
                </CardData>
            </div>


            <div className={cn(cardClassName,)}>
                <ProgressBar
                    type="memory"
                    values={[data.memory.used, data.memory.total]}
                    unit={data.memory.unit}
                    className="bg-gradient-to-tr from-emerald-400 via-teal-500 to-sky-500"
                />
                <CardData title="Memory" icon={Database}>
                    <table className="w-full text-xs text-muted-foreground">
                        <tbody className="divide-y divide-border">
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Utilized</td>
                                <td className="py-1 text-right">{data.memory.used} {data.memory.unit}</td>
                            </tr>
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Total</td>
                                <td className="py-1 text-right">{data.memory.total} {data.memory.unit}</td>
                            </tr>
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Free</td>
                                <td className="py-1 text-right  text-green-500">{data.memory.total - data.memory.used} {data.memory.unit}</td>
                            </tr>
                        </tbody>
                    </table>
                </CardData>
            </div>


            <div className={cn(cardClassName,)}>
                <ProgressBar
                    type="storage"
                    values={[data.storage.used, data.storage.total]}
                    unit={data.storage.unit}
                    className="bg-gradient-to-br from-purple-500 via-fuchsia-500 to-pink-500"
                />
                <CardData title="Storage" icon={HardDrive}>
                    <table className="w-full text-xs text-muted-foreground">
                        <tbody className="divide-y divide-border">
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Utilized</td>
                                <td className="py-1 text-right">{data.storage.used} {data.storage.unit}</td>
                            </tr>
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Total</td>
                                <td className="py-1 text-right">{data.storage.total} {data.storage.unit}</td>
                            </tr>
                            <tr className="h-6">
                                <td className="py-1 pr-4 font-medium text-foreground/80">Free</td>
                                <td className="py-1 text-right  text-green-500">{data.storage.total - data.storage.used} {data.storage.unit}</td>
                            </tr>
                        </tbody>
                    </table>
                </CardData>
            </div>
        </div>
    )
}


function CardData({ children, icon: Icon, title }: CardDataType) {
    return (
        <div className="col-span-4 px-3 py-2 flex flex-col justify-between">
            <div className="flex justify-between">
                <h3 className="text-lg font-semibold">{title}</h3>

                <div className="w-20 aspect-square grid place-items-center">
                    <Icon size={60} className="text-foreground/70" />
                </div>
            </div>
            {children}
        </div>
    )
}


function ProgressBar({ values, type, className, unit }: { values: number[], type: "node" | "cpu" | "memory" | "storage", className?: string, unit?: string }) {

    const percentage = Math.round((values[0] / values[1]) * 100)

    const [fill, setFill] = useState(0);
    const prevFill = useRef(0);


    let label;
    switch (type) {
        case 'node':
            label = `${values[0]} Online`
            break;

        case 'cpu':
            label = `${values[0]} ${unit}`
            break;

        case 'memory':
            label = `${values[0]}/${values[1]} ${unit}`
            break;

        case 'storage':
            label = `${values[0]}/${values[1]} ${unit}`
            break;

        default:
            break;
    }



    function lerp(start: number, end: number, t: number): number {
        return start + (end - start) * t;
    }

    useEffect(() => {
        const start = prevFill.current;
        const end = percentage;
        let progress = 0;

        const step = () => {
            progress += 0.05; // animation speed
            const value = lerp(start, end, progress);
            setFill(value);

            if (progress < 1) requestAnimationFrame(step);
            else prevFill.current = end; // remember the last value
        };

        step();
    }, [percentage]);



    return (
        <div className="w-full h-full pt-4 dark:bg-accent/20 bg-accent/50 border-r rounded-r">
            <div className="h-[calc(100%-2rem)]">
                <div className="w-[2rem] h-full mx-auto bg-accent flex flex-col-reverse border rounded-t overflow-hidden">
                    <div className={cn(className, "w-full transition-colors duration-300 ease-in-out",)} style={{ height: `${fill}%` }} />
                </div>
            </div>

            <hr />

            <div className="h-8 flex items-center justify-center" >
                {label ?
                    <p className="text-xs text-muted-foreground">{label}</p>
                    :
                    <div className="w-8 h-3 bg-accent rounded animate-pulse" role="status" />
                }
            </div>
        </div>
    )
}
