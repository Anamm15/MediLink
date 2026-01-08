import * as React from "react";
import { cn } from "@/lib/utils";
import { Check } from "lucide-react";
import { ErrorText } from "./ErrorText";
import { Label } from "./Label";

export interface CheckboxProps
  extends Omit<React.InputHTMLAttributes<HTMLInputElement>, "type"> {
  label?: React.ReactNode;
  description?: string;
  error?: string | boolean;
}

const Checkbox = React.forwardRef<HTMLInputElement, CheckboxProps>(
  ({ className, label, description, error, disabled, id, ...props }, ref) => {
    // Generate random ID if not provided, for linking accessbility
    const uniqueId = id || React.useId();

    return (
      <div className={cn("flex items-start space-x-3", className)}>
        <div className="relative flex items-center">
          <input
            type="checkbox"
            id={uniqueId}
            ref={ref}
            disabled={disabled}
            className="peer sr-only" // Hide checkbox from screen readers
            {...props}
          />

          {/* Custom Checkbox Box */}
          <div
            className={cn(
              "h-5 w-5 shrink-0 rounded-md border border-slate-300 bg-white ring-offset-white transition-all",
              // State: Focus (Keyboard navigation)
              "peer-focus-visible:outline-none peer-focus-visible:ring-2 peer-focus-visible:ring-slate-400 peer-focus-visible:ring-offset-2",
              // State: Checked
              "peer-checked:bg-slate-900 peer-checked:text-white peer-checked:border-slate-900",
              // State: Disabled
              "peer-disabled:cursor-not-allowed peer-disabled:opacity-50",
              // State: Error
              error &&
                "border-red-500 peer-checked:bg-red-500 peer-checked:border-red-500",
              "flex items-center justify-center"
            )}
          >
            {/* Icon Checkmark */}
            <Check
              size={14}
              strokeWidth={3}
              className={cn(
                "text-current opacity-0 transition-opacity duration-200"
              )}
            />
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="3"
              strokeLinecap="round"
              strokeLinejoin="round"
              className="absolute h-3.5 w-3.5 text-white opacity-0 peer-checked:opacity-100 pointer-events-none transition-opacity"
            >
              <polyline points="20 6 9 17 4 12" />
            </svg>
          </div>
        </div>

        <div className="grid gap-1.5 leading-none">
          {label && (
            <Label htmlFor={uniqueId} required={props.required}>
              {label}
            </Label>
          )}

          {description && (
            <p className="text-sm text-slate-500">{description}</p>
          )}

          <ErrorText error={error} helperText={error as string} />
        </div>
      </div>
    );
  }
);
Checkbox.displayName = "Checkbox";

export { Checkbox };
