import TeamGeneralInfo from "./TeamGeneralInfo";

export default function TeamSettingsView() {
    return (
        <div className="space-y-5 pb-20">
            <h1 className="text-2xl font-bold pt-3 ">Team Settings</h1>

            <TeamGeneralInfo />
        </div>
    )
}
