// src/hooks/useMetrics.ts
import { useEffect, useState } from "react";

export function useMetrics() {
    const [metrics, setMetrics] = useState<any[]>([]);

    useEffect(() => {
        const ws = new WebSocket("ws://192.168.0.202:3905/api/v1/ws/ui");

        ws.onmessage = (event) => {
            try {
                const data = JSON.parse(event.data);
                setMetrics((prev) => [data, ...prev,]); // append new metrics
            } catch (err) {
                console.error("Bad message:", event.data);
            }
        };

        ws.onopen = () => console.log("✅ UI connected to server");
        ws.onclose = () => console.log("❌ UI disconnected");

        return () => ws.close();
    }, []);

    return metrics;
}
