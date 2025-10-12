import * as React from "react"
import { Eye, EyeOff } from "lucide-react"
import { Button } from "./button"
import { Input } from "./input"

interface InputPasswordProps extends React.InputHTMLAttributes<HTMLInputElement> {
    label?: string
    placeholder?: string
}

export const InputPassword = React.forwardRef<HTMLInputElement, InputPasswordProps>(
    ({ label, placeholder = "Enter password", className, ...props }, ref) => {
        const [show, setShow] = React.useState(false)

        return (
            <div className="space-y-2">
                {label && <label className="text-sm font-medium">{label}</label>}

                <div className="relative">
                    <Input
                        type={show ? "text" : "password"}
                        placeholder={placeholder}
                        ref={ref}
                        className={`pr-10 ${className || ""}`}
                        {...props}
                    />

                    <Button
                        type="button"
                        variant="ghost"
                        size="icon-lg"
                        onClick={() => setShow(!show)}
                        className="absolute right-1 top-1/2 -translate-y-1/2 h-7 w-7 text-muted-foreground hover:text-foreground"
                    >
                        {show ? <EyeOff className="h-4 w-4" /> : <Eye className="h-4 w-4" />}
                    </Button>
                </div>
            </div>
        )
    }
)

InputPassword.displayName = "InputPassword"
