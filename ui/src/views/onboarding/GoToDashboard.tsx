import { Button } from "@/components/cn/button";
import CsImage from "@/constants/image";
import { motion } from "framer-motion";
import { Link } from "react-router-dom";

export default function GoToDashboard() {
    return (
        <div className="flex flex-col items-center justify-between h-full py-10">
            <div className="">
                <motion.h1
                    className="text-center font-black text-3xl pt-5 text-primary"
                    initial={{ opacity: 0, y: -20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ duration: 0.6 }}
                >
                    You&prime;re all done! 🎉
                </motion.h1>
                <motion.p
                    className="text-center text-muted-foreground mt-3"
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ delay: 0.3 }}
                >
                    Everything&prime;s ready to roll. Let&prime;s jump into your dashboard.
                </motion.p>
            </div>

            <motion.img
                className="max-w-sm aspect-auto"
                src={CsImage.onboarding.allDone}
                width={300}
                height={300}
                initial={{ scale: 0.8, opacity: 0 }}
                animate={{ scale: 1, opacity: 1 }}
                transition={{ duration: 0.3 }}
            />

            <motion.div
                initial={{ scale: 0.8, opacity: 0 }}
                animate={{ scale: 1, opacity: 1 }}
                transition={{ delay: 0.5, type: "spring" }}
            >
                <Link to={'/'} replace>
                    <Button size={'sm'}>
                        Let&prime;s Go →
                    </Button>
                </Link>

            </motion.div>
        </div>
    )
}
