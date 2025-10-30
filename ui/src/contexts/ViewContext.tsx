import type { LucideIcon } from "lucide-react";
import { createContext, useState, useCallback, type ReactNode } from "react";

type ViewData = {
    title: string;
    subtitle?: string;
    icon?: LucideIcon | null;
};

type ViewContextType = ViewData & {
    setView: (data: Partial<ViewData>) => void;
};

export const ViewContext = createContext<ViewContextType>({
    title: "",
    subtitle: "",
    icon: null,
    setView: () => { },
});

export function ViewProvider({ children }: { children: ReactNode }) {
    const [view, setViewState] = useState<ViewData>({
        title: "",
        subtitle: "",
        icon: null,
    });

    // useCallback ensures stable reference for dependency arrays (like in useEffect)
    const setView = useCallback((data: Partial<ViewData>) => {
        setViewState((prev) => ({ ...prev, ...data }));
    }, []);

    return (
        <ViewContext.Provider value={{ ...view, setView }}>
            {children}
        </ViewContext.Provider>
    );
}
