import { Card, CardAction, CardDescription, CardHeader, CardTitle } from "@/components/cn/card";
import { useCsSelector } from "@/cs-redux";

export default function AccountAvatar() {
    const { authUser } = useCsSelector(state => state.auth);

    return (
        <Card>
            <CardHeader>
                <CardTitle>Avatar</CardTitle>
                <CardDescription>This is your avatar. <br />
                    Click on the avatar to upload a custom one from your files.
                </CardDescription>
                <CardAction className="aspect-square w-22 rounded-full overflow-hidden">
                    <img width={200} height={200} src={authUser?.avatar} alt={authUser?.display_name} />
                </CardAction>
            </CardHeader>
        </Card>
    )
}
