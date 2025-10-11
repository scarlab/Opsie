import { Button } from "@/components/cn/button"
import Config from "@/config"
import CsImage from "@/constants/image"
import { motion } from "framer-motion"

export default function Welcome({ next }: { next?: () => void }) {
    return (
        <div className="flex flex-col items-center justify-center h-full ">
            <motion.img
                className=" "
                src={CsImage.onboarding.hello}
                width={300}
                height={300}
                initial={{ scale: 0.8, opacity: 0 }}
                animate={{ scale: 1, opacity: 1 }}
                transition={{ duration: 0.3 }}
            />


            <motion.h1
                className="text-center font-black text-4xl text-primary mt-10"
                initial={{ opacity: 0, y: -20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.6 }}
            >
                Welcome to your workspace
            </motion.h1>

            <motion.p
                className="text-center text-muted-foreground mt-3"
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                transition={{ delay: 0.3 }}
            >
                Everything you need to manage, analyze, and grow -<br />right inside your {Config.projectName} dashboard.
            </motion.p>

            <motion.div
                className="mt-10"
                initial={{ scale: 0.8, opacity: 0 }}
                animate={{ scale: 1, opacity: 1 }}
                transition={{ delay: 0.5, type: "spring" }}
            >
                <Button onClick={next} className="text-lg px-6 py-3 mt-7 shadow-md hover:shadow-xl transition-all">
                    Let's Started
                </Button>
            </motion.div>
        </div>
    )
}
