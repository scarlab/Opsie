import { easeInOut, motion } from "framer-motion";

export default function BouncingDotsLoader() {
    const bounce = {
        y: [0, -10, 0],
        transition: {
            repeat: Infinity,
            duration: 0.6,
            ease: easeInOut
        },
    };

    return (
        <div className="w-svw h-svh grid place-items-center bg-background">
            <div className="flex items-center gap-2">
                {[0, 1, 2].map((i) => (
                    <motion.span
                        key={i}
                        className="w-3 h-3 bg-primary rounded-full"
                        animate={bounce}
                        transition={{ delay: i * 0.15, repeat: Infinity, duration: 0.6 }}
                    />
                ))}
            </div>
            <motion.p
                className="mt-6 text-sm text-muted-foreground"
                animate={{ opacity: [0.3, 1, 0.3] }}
                transition={{ duration: 1.5, repeat: Infinity }}
            >
                Checking session...
            </motion.p>
        </div>
    );
}
