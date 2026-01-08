import * as React from "react";
import { cn } from "@/lib/utils";
import { Loader2 } from "lucide-react";

// Define available variants
export type ButtonVariant =
  | "primary"
  | "secondary"
  | "outline"
  | "destructive"
  | "ghost"
  | "link";

// Define available sizes
type ButtonSize = "sm" | "default" | "lg" | "icon";

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant;
  size?: ButtonSize;
  isLoading?: boolean;
  startIcon?: React.ReactNode;
  endIcon?: React.ReactNode;
}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  (
    {
      className,
      variant = "primary",
      size = "default",
      isLoading = false,
      startIcon,
      endIcon,
      children,
      disabled,
      type = "button",
      ...props
    },
    ref
  ) => {
    // Base Styles: Layout, typography, focus states, and disabled states
    const baseStyles =
      "inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-slate-400 disabled:pointer-events-none disabled:opacity-50";

    // Variant Styles: Colors and backgrounds
    const variants: Record<ButtonVariant, string> = {
      primary: "bg-slate-900 text-slate-50 hover:bg-slate-900/90 shadow-sm",
      secondary: "bg-slate-100 text-slate-900 hover:bg-slate-100/80",
      outline:
        "border border-slate-200 bg-transparent hover:bg-slate-100 text-slate-900",
      destructive: "bg-red-500 text-slate-50 hover:bg-red-500/90 shadow-sm",
      ghost: "hover:bg-slate-100 hover:text-slate-900 text-slate-600",
      link: "text-slate-900 underline-offset-4 hover:underline",
    };

    // Size Styles: Padding and height
    const sizes: Record<ButtonSize, string> = {
      default: "h-10 px-4 py-2",
      sm: "h-9 rounded-md px-3",
      lg: "h-11 rounded-md px-8",
      icon: "h-10 w-10",
    };

    return (
      <button
        ref={ref}
        type={type}
        disabled={disabled || isLoading}
        className={cn(baseStyles, variants[variant], sizes[size], className)}
        {...props}
      >
        {isLoading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}

        {!isLoading && startIcon && (
          <span className="mr-2 inline-flex">{startIcon}</span>
        )}

        {children}

        {!isLoading && endIcon && (
          <span className="ml-2 inline-flex">{endIcon}</span>
        )}
      </button>
    );
  }
);
Button.displayName = "Button";

export { Button };
