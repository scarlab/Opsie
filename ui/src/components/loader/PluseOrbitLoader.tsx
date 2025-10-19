import { motion } from "framer-motion";

export default function PluseOrbitLoader() {
    return (
        <div className="w-svw h-svh grid place-items-center bg-background">
            <div className="relative w-16 h-16">
                <motion.span
                    className="absolute inset-0 rounded-full border-4 border-primary border-t-transparent"
                    animate={{ rotate: 360 }}
                    transition={{
                        repeat: Infinity,
                        duration: 1,
                        ease: "linear",
                    }}
                />
                <motion.span
                    className="absolute inset-3 rounded-full bg-primary/10"
                    animate={{ scale: [1, 1.1, 1] }}
                    transition={{
                        repeat: Infinity,
                        duration: 1.2,
                        ease: "easeInOut",
                    }}
                />
            </div>
            <motion.p
                className="mt-6 text-sm text-muted-foreground"
                animate={{ opacity: [0.3, 1, 0.3] }}
                transition={{ duration: 1.5, repeat: Infinity }}
            >
                Loading your session...
            </motion.p>
        </div>
    );
}
