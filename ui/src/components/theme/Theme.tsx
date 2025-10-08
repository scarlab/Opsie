"use client";

import { useEffect, useState } from "react";
import { MonitorCog, Moon, Sun } from "lucide-react";
import { useTheme } from "next-themes";
import { cn } from "@/lib/utils";

export default function Theme({ className }: { className?: string }) {
    const { setTheme, theme } = useTheme();
    const [mounted, setMounted] = useState(false);

    useEffect(() => {
        setMounted(true);
    }, []);

    if (!mounted) return null; // avoid hydration mismatch
    const commonClass = 'rounded-full p-1 cursor-pointer';
    return (
        <span className={cn("flex items-center gap-0.5 border rounded-full", className)}>
            <button className={`${commonClass} ${theme === 'system' && "bg-primary text-primary-foreground"}`} onClick={() => setTheme('system')}><MonitorCog size={15} /></button>
            <button className={`${commonClass} ${theme === 'light' && "bg-primary text-primary-foreground"}`} onClick={() => setTheme('light')}><Sun size={15} /></button>
            <button className={`${commonClass} ${theme === 'dark' && "bg-primary text-primary-foreground"}`} onClick={() => setTheme('dark')}><Moon size={15} /></button>
        </span>
    );
}