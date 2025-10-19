import { motion } from "framer-motion";

export default function RotatingGradientRing() {
    return (
        <div >
            <motion.div
                className="w-16 h-16 rounded-full border-4 border-transparent border-t-primary border-r-primary/40"
                animate={{ rotate: 360 }}
                transition={{
                    repeat: Infinity,
                    duration: 1,
                    ease: "linear",
                }}
            />
            <motion.p
                className="mt-6 text-sm text-muted-foreground tracking-wide"
                animate={{ opacity: [0.3, 1, 0.3] }}
                transition={{ duration: 1.5, repeat: Infinity }}
            >
                Initializing...
            </motion.p>
        </div>
    );
}
