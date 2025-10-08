import { Theme } from "@/components/theme";
import Config from "@/config";
import Welcome from "./Welcome";
import { useState } from "react";
import Owner from "./Owner";
import GoToDashboard from "./GoToDashboard";


export default function OnboardingView() {
    const [state, setState] = useState(0);
    function next() {
        setState(state + 1)
    }

    return (
        <div className="grid place-items-center bg-secondary h-svh w-vw">
            <div className="px-5 pb-3 fixed bottom-0 left-0 w-full flex items-end justify-between">
                <div>
                    <small className="text-muted-foreground">&copy; scarlab.in / opsie - {Config.version}</small>
                </div>
                <div>
                    <Theme />
                </div>
            </div>
            <main className="shadow rounded-lg bg-background px-3 py-2 max-w-5xl w-full aspect-video ">
                {state === 0 && <Welcome next={next} />}
                {state === 1 && <Owner next={next} />}
                {state === 2 && <GoToDashboard />}
            </main>
        </div>
    )
}


// project-scarlab: 22.717724992702212, 88.48923629425165