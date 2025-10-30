import TeamGeneralInfo from "./TeamGeneralInfo";
import View from "@/components/utils/View";

export default function TeamSettingsView() {


    return (
        <View title="Team Settings" className="space-y-5 pb-20">
            <h1 className="text-2xl font-bold  px-3">Team Settings</h1>
            <TeamGeneralInfo />
        </View>
    )
}
