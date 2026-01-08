import * as React from "react";
import { cn } from "@/lib/utils";
import { Label } from "./Label";

export interface SwitchProps
  extends Omit<React.InputHTMLAttributes<HTMLInputElement>, "type"> {
  label?: string;
  description?: string;
  error?: string | boolean;
}

const Switch = React.forwardRef<HTMLInputElement, SwitchProps>(
  ({ className, label, description, error, id, disabled, ...props }, ref) => {
    // Generate unique ID for accessibility if not provided
    const uniqueId = id || React.useId();

    return (
      <div className={cn("flex items-center justify-between", className)}>
        <div className="grid gap-1.5">
          {label && <Label htmlFor={uniqueId}>{label}</Label>}

          {description && (
            <p className="text-sm text-slate-500">{description}</p>
          )}

          {/* Error Message */}
          {error && typeof error === "string" && (
            <p className="text-xs font-medium text-red-500 animate-pulse">
              {error}
            </p>
          )}
        </div>

        <div className="relative inline-flex items-center ml-4">
          <input
            type="checkbox"
            id={uniqueId}
            ref={ref}
            disabled={disabled}
            className="peer sr-only" // Hide native checkbox
            {...props}
          />

          {/* Switch Track (Background) */}
          <div
            className={cn(
              "h-6 w-11 cursor-pointer rounded-full border-2 border-transparent bg-slate-200 transition-colors duration-200 ease-in-out",
              // Focus state
              "peer-focus-visible:outline-none peer-focus-visible:ring-2 peer-focus-visible:ring-slate-400 peer-focus-visible:ring-offset-2",
              // Checked state (Background becomes dark/primary)
              "peer-checked:bg-slate-900",
              // Disabled state
              "peer-disabled:cursor-not-allowed peer-disabled:opacity-50",
              // Error state
              error && "bg-red-200 peer-checked:bg-red-500"
            )}
          ></div>

          {/* Switch Thumb (Moving Circle) */}
          <div
            className={cn(
              "pointer-events-none absolute left-0.5 h-5 w-5 rounded-full bg-white shadow-lg ring-0 transition-transform duration-200 ease-in-out",
              // Move to the right when checked
              "peer-checked:translate-x-5"
            )}
          ></div>
        </div>
      </div>
    );
  }
);
Switch.displayName = "Switch";

export { Switch };
