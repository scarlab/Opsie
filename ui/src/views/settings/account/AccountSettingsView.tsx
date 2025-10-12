import AccountAvatar from "./AccountAvatar";
import AccountDisplayName from "./AccountDisplayName";
import AccountEmail from "./AccountEmail";
import AccountPassword from "./AccountPassword";

export default function AccountSettingsView() {
    return (
        <div className="space-y-5 pb-20">
            <h1 className="text-2xl font-bold pt-3 ">Account Settings</h1>

            <AccountAvatar />
            <AccountDisplayName />
            <AccountEmail />
            <AccountPassword />
        </div>
    )
}
