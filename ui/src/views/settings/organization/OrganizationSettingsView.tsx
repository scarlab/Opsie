import OrgGeneralInfo from "./OrgGeneralInfo";

export default function OrganizationSettingsView() {
    return (
        <div className="space-y-5 pb-20">
            <h1 className="text-2xl font-bold pt-3 ">Organization Settings</h1>

            <OrgGeneralInfo />
        </div>
    )
}
