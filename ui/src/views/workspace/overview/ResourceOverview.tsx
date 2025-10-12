import { Button } from "@/components/cn/button";
import { Box, Ellipsis, Globe } from "lucide-react";
import { Link } from "react-router-dom";

export default function ResourceOverview() {
    return (
        <div className="px-3 mt-14">
            <h1 className="text-2xl font-bold mb-2 flex items-center gap-3"><Box size={20} />  Resources</h1>
            <div className="h-0.5 bg-muted" />

            <div className="mt-3">
                <div className="grid group grid-cols-12 text-xs px-3 mb-1 font-semibold text-accent-foreground/90">
                    <p className="col-span-6">Resource Name </p>
                    <p className="col-span-2">Resource Node</p>
                    <p className="col-span-1">Uptime</p>
                    <p className="col-span-1">Runtime</p>
                    <p className="col-span-1 text-center">Status</p>
                    <p className="col-span-1 text-right">Actions</p>
                </div>
                {
                    [1, 2, 3].map((r, i) => (
                        <div key={`kaj_syd_b87a7sd${i}`} className="grid group grid-cols-12 bg-accent/30 hover:bg-accent/50 transition-colors border-b px-3 py-2">
                            <Link to={`/resources/${r}`} className="col-span-6 hover:underline hover:text-blue-600 dark:hover:text-blue-400">
                                <p className="flex items-center gap-1.5 font-medium text-lg"><Globe size={17} className="text-muted-foreground " /> Raspid_r{(r / 4 * 345 - (734 * 8374)).toFixed().slice(3, -2)}</p>
                            </Link>

                            <div className="col-span-2 flex items-center">
                                <p className="text-sm text-muted-foreground">Node - nx{r * (7 * 77)}</p>
                            </div>

                            <div className="col-span-1 flex items-center">
                                <p className="text-sm text-muted-foreground">96 Days</p>
                            </div>

                            <div className="col-span-1 flex items-center">
                                <p className="text-sm font-medium">Docker</p>
                            </div>

                            <div className="col-span-1 flex justify-center items-center">
                                <span className="text-xs text-green-950 font-medium border px-2 rounded-2xl bg-green-500">Deployed</span>
                            </div>

                            <div className="col-span-1 flex justify-end items-center">
                                <Button size={'icon-sm'} variant={'ghost'}><Ellipsis /></Button>
                            </div>
                        </div>
                    ))
                }
            </div>
        </div>
    )
}
