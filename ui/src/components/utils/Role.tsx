import type { SystemRolesType } from "@/constants";
import { cn } from "@/lib/utils";


export default function Role({ sysRole, className, ...props }: React.ComponentProps<"span"> & { sysRole: SystemRolesType }) {
    return (
        <span
            className={cn(
                "capitalize rounded-full font-semibold text-xs inline-block",
                sysRole === 'owner' && ' text-green-600 dark:text-green-500',
                sysRole === 'admin' && ' text-orange-500',
                sysRole === 'staff' && ' text-cyan-500',
                className,
            )}
            {...props}
        >
            {sysRole}
        </span>
    )
}
