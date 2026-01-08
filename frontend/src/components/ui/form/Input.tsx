import * as React from "react";
import { cn } from "@/lib/utils";
import { ErrorText } from "./ErrorText";
import { Label } from "./Label";

export interface InputProps
  extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string | boolean;
  helperText?: string;
  startIcon?: React.ReactNode;
  endIcon?: React.ReactNode;
}

const Input = React.forwardRef<HTMLInputElement, InputProps>(
  (
    {
      label,
      className,
      type,
      error,
      helperText,
      startIcon,
      endIcon,
      disabled,
      ...props
    },
    ref
  ) => {
    return (
      <div className="w-full">
        {/* Render Label if provided */}
        <div>
          {label && (
            <Label htmlFor={props.id} required={props.required}>
              {label}
            </Label>
          )}
        </div>

        <div className="relative">
          {startIcon && (
            <div className="absolute left-3 top-1/2 -translate-y-1/2 text-slate-500 pointer-events-none">
              {startIcon}
            </div>
          )}

          <input
            type={type}
            className={cn(
              // Base Styles
              "flex h-10 w-full rounded-md border border-slate-300 bg-white px-3 py-2 text-sm placeholder:text-slate-400",
              // Focus Styles (Ring effect)
              "focus:outline-none focus:ring-2 focus:ring-slate-400 focus:ring-offset-2",
              // Disabled Styles
              "disabled:cursor-not-allowed disabled:opacity-50 disabled:bg-slate-100",
              // File Input Styles
              "file:border-0 file:bg-transparent file:text-sm file:font-medium",
              // Padding adjustment
              startIcon ? "pl-10" : "",
              endIcon ? "pr-10" : "",
              // Error State Style
              error && "border-red-500 focus:ring-red-200",
              className
            )}
            ref={ref}
            disabled={disabled}
            {...props}
          />

          {/* Render End Icon */}
          {endIcon && (
            <div className="absolute right-3 top-1/2 -translate-y-1/2 text-slate-500 pointer-events-none">
              {endIcon}
            </div>
          )}
        </div>

        {/* Render Error Message or Helper Text */}
        <ErrorText error={error} helperText={helperText} />
      </div>
    );
  }
);
Input.displayName = "Input";

export { Input };
