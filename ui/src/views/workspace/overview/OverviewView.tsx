
export default function OverviewView() {
    return (
        <div className="">
            <div className="flex  flex-col gap-4 p-4">
                <div className="flex gap-4 flex-wrap">
                    {Array.from({ length: 20 }).map((_, i) => (
                        <div key={i} className="bg-slate-600 aspect-square w-52 rounded-lg" >
                            {i + 1}
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}
