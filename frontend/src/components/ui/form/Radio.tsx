import * as React from "react";
import { cn } from "@/lib/utils";
import { Label } from "./Label";

export interface RadioProps
  extends React.InputHTMLAttributes<HTMLInputElement> {
  label: string;
  description?: string;
  error?: boolean;
}

const Radio = React.forwardRef<HTMLInputElement, RadioProps>(
  ({ className, label, description, error, id, ...props }, ref) => {
    // Generate unique ID for accessibility if not provided
    const uniqueId = id || React.useId();

    return (
      <div className={cn("flex items-start space-x-3", className)}>
        <div className="relative flex items-center pt-0.5">
          <input
            type="radio"
            id={uniqueId}
            ref={ref}
            className="peer sr-only" // Hide native radio button
            {...props}
          />

          {/* Custom Radio Circle (Outer Ring) */}
          <div
            className={cn(
              "h-4 w-4 rounded-full border border-slate-300 bg-white ring-offset-white transition-all",
              // Focus state (Keyboard navigation)
              "peer-focus-visible:ring-2 peer-focus-visible:ring-slate-400 peer-focus-visible:ring-offset-2",
              // Checked state (Border color changes)
              "peer-checked:border-slate-900 peer-checked:text-slate-900",
              // Disabled state
              "peer-disabled:cursor-not-allowed peer-disabled:opacity-50",
              // Error state
              error && "border-red-500",
              // Flex to center the inner dot
              "flex items-center justify-center"
            )}
          >
            {/* Inner Dot (Only visible when checked) */}
            <div className="h-2 w-2 rounded-full bg-current opacity-0 scale-0 transition-all duration-200 peer-checked:opacity-100 peer-checked:scale-100" />
          </div>
        </div>

        <div className="grid gap-1.5 leading-none">
          <Label htmlFor={uniqueId} required={props.required}>
            {label}
          </Label>

          {description && (
            <p className="text-sm text-slate-500">{description}</p>
          )}
        </div>
      </div>
    );
  }
);
Radio.displayName = "Radio";

export { Radio };
