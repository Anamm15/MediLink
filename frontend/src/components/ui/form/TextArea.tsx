import * as React from "react";
import { cn } from "@/lib/utils";
import { ErrorText } from "./ErrorText";
import { Label } from "./Label";

export interface TextareaProps
  extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {
  label?: string;
  helperText?: string;
  error?: string | boolean;
}

const Textarea = React.forwardRef<HTMLTextAreaElement, TextareaProps>(
  ({ className, label, helperText, error, id, ...props }, ref) => {
    // Generate unique ID for accessibility if not provided
    const uniqueId = id || React.useId();

    return (
      <div className="w-full">
        {/* Render Label if provided */}
        {label && (
          <Label htmlFor={uniqueId} required={props.required}>
            {label}
          </Label>
        )}

        <textarea
          id={uniqueId}
          className={cn(
            // Base styles
            "flex min-h-[80px] w-full rounded-md border border-slate-300 bg-white px-3 py-2 text-sm placeholder:text-slate-400",
            // Focus styles (Ring effect)
            "focus:outline-none focus:ring-2 focus:ring-slate-400 focus:ring-offset-2",
            // Disabled styles
            "disabled:cursor-not-allowed disabled:opacity-50",
            // Error state styling
            error && "border-red-500 focus:ring-red-200",
            className
          )}
          ref={ref}
          {...props}
        />

        {/* Helper Text or Error Message logic */}
        <ErrorText error={error} helperText={helperText} />
      </div>
    );
  }
);
Textarea.displayName = "Textarea";

export { Textarea };
