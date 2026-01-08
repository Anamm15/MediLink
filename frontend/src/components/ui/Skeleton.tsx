import { cn } from "@/lib/utils";

function Skeleton({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      className={cn(
        // Base styles:
        "animate-pulse rounded-md bg-slate-200",
        className
      )}
      {...props}
    />
  );
}

export { Skeleton };
