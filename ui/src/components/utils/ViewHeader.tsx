import { useEffect } from "react";
import { useLocation } from "react-router-dom";
import type { LucideIcon } from "lucide-react";
import { useViewContext } from "@/hooks/useViewContext";


type ViewProps = React.HTMLAttributes<HTMLDivElement> & {
    title: string;
    subtitle?: string;
    icon?: LucideIcon;
};

/**
 * View component
 * Handles page metadata (title, subtitle, icon) and resets on route change.
 * Can be used as the root layout for a page.
 *
 * Example:
 * <View title="Dashboard" subtitle="Overview of system metrics">
 *   <DashboardContent />
 * </View>
 */
export default function ViewHeader({ title, subtitle, icon, }: ViewProps) {
    const { setView } = useViewContext();
    const location = useLocation();

    useEffect(() => {
        // Set current page context
        setView({ title, subtitle, icon });

        // Reset on unmount or route change
        return () => setView({ title: "", subtitle: "", icon: null });
    }, [location.pathname, title, subtitle, icon, setView]);

    return null
}
