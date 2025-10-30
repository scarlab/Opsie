import { useContext } from "react";
import { ViewContext } from "@/contexts/ViewContext";

export function useViewContext() {
    return useContext(ViewContext);
}

/**
 * Optional helper hook to auto-set title on mount.
 * Example: useSetView("Projects", "All active projects")
 */
export const useSetView = () => {
    const { setView } = useContext(ViewContext);
    return setView;
};
