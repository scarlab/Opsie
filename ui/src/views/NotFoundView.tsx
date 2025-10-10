import Image from "@/constants/image";

export default function NotFoundView() {
    return (
        <div className="grid place-items-center h-[calc(100vh-var(--header-height))]">
            <div className="text-center">
                <img
                    className="mb-20"
                    width={400}
                    height={400}
                    src={Image.pageNotFound}
                    alt=""
                />
                <h1 className="text-primary text-3xl font-black">404 - Not Found!</h1>
                <p className="text-muted-foreground">The page you are trying to reach is not found!</p>
            </div>
        </div>
    )
}
