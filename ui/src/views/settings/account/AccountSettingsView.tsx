import AccountAvatar from "./AccountAvatar";
import AccountDisplayName from "./AccountDisplayName";
import AccountEmail from "./AccountEmail";
import AccountPassword from "./AccountPassword";
import View from "@/components/utils/View";

export default function AccountSettingsView() {

    return (
        <View title="Account Settings" className="space-y-5 pb-20">
            <h1 className="text-2xl font-bold  px-3">Account Settings</h1>
            <AccountAvatar />
            <AccountDisplayName />
            <AccountEmail />
            <AccountPassword />
        </View>
    )
}
