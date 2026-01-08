import * as React from "react";
import { cn } from "@/lib/utils";
import { ChevronDown } from "lucide-react";
import { ErrorText } from "./ErrorText";
import { Label } from "./Label";

export interface SelectProps
  extends React.SelectHTMLAttributes<HTMLSelectElement> {
  label?: string;
  helperText?: string;
  error?: string | boolean;
}

const Select = React.forwardRef<HTMLSelectElement, SelectProps>(
  (
    { className, children, label, helperText, error, id, disabled, ...props },
    ref
  ) => {
    // Generate unique ID if not provided
    const uniqueId = id || React.useId();

    return (
      <div className="w-full space-y-2">
        {/* Render Label if provided */}
        {label && (
          <Label htmlFor={uniqueId} required={props.required}>
            {label}
          </Label>
        )}

        <div className="relative">
          <select
            id={uniqueId}
            ref={ref}
            disabled={disabled}
            className={cn(
              // Base styles
              "flex h-10 w-full items-center justify-between rounded-md border border-slate-300 bg-white px-3 py-2 text-sm placeholder:text-slate-400",
              // Remove default browser appearance (arrow)
              "appearance-none",
              // Focus styles (Ring effect to match Input component)
              "focus:outline-none focus:ring-2 focus:ring-slate-400 focus:ring-offset-2",
              // Disabled state
              "disabled:cursor-not-allowed disabled:opacity-50",
              // Error state styling
              error && "border-red-500 focus:ring-red-200",
              "pr-10",
              className
            )}
            {...props}
          >
            {children}
          </select>

          <div className="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-slate-500">
            <ChevronDown size={16} />
          </div>
        </div>

        {/* Helper Text or Error Message logic */}
        <ErrorText error={error} helperText={helperText} />
      </div>
    );
  }
);

Select.displayName = "Select";

export { Select };
