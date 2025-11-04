import { useState } from "react";
import UserActivity from "./UserActivity";
import UserTeams from "./UserTeams";


export default function Tabs() {
    const [tab, setTab] = useState<"team" | "activity">("activity");
    return (
        <div>
            <div className="border-b mb-3">
                <button onClick={() => setTab("activity")} className={` py-1 px-3 border-b-2 cursor-pointer ${tab === "activity" ? "border-b-primary text-primary" : "border-b-transparent"}`}>Activity</button>
                <button onClick={() => setTab("team")} className={` py-1 px-3 border-b-2 cursor-pointer ${tab === "team" ? "border-b-primary text-primary" : "border-b-transparent"}`}>Teams</button>
            </div>
            {tab === "team" && <UserTeams />}
            {tab === "activity" && <UserActivity />}
        </div>
    )
}
