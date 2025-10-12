import { Card, CardAction, CardDescription, CardHeader, CardTitle } from "@/components/cn/card";
import CsImage from "@/constants/image";

export default function AccountAvatar() {
    return (
        <Card>
            <CardHeader>
                <CardTitle>Avatar</CardTitle>
                <CardDescription>This is your avatar. <br />
                    Click on the avatar to upload a custom one from your files.
                </CardDescription>
                <CardAction className="aspect-square w-22 rounded-full overflow-hidden">
                    <img width={200} height={200} src={CsImage.user} alt="" />
                </CardAction>
            </CardHeader>
        </Card>
    )
}
