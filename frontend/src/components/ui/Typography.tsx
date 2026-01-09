import React from "react";
import { cn } from "@/lib/utils";

export function TypographyH1({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLHeadingElement>) {
  return (
    <h1
      className={cn(
        "scroll-m-20 font-extrabold tracking-tight text-slate-900 dark:text-slate-50",
        "text-3xl",
        "md:text-4xl",
        "lg:text-5xl",
        "2xl:text-6xl",
        className
      )}
      {...props}
    >
      {children}
    </h1>
  );
}

export function TypographyH2({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLHeadingElement>) {
  return (
    <h2
      className={cn(
        "scroll-m-20 font-bold tracking-tight text-slate-900 dark:text-slate-50 first:mt-0 mt-10 pb-2 border-b border-slate-200",
        "text-2xl",
        "md:text-3xl",
        "lg:text-4xl",
        className
      )}
      {...props}
    >
      {children}
    </h2>
  );
}

export function TypographyH3({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLHeadingElement>) {
  return (
    <h3
      className={cn(
        "scroll-m-20 font-semibold tracking-tight text-slate-900 dark:text-slate-50 mt-8",
        "text-xl",
        "md:text-2xl",
        "2xl:text-3xl",
        className
      )}
      {...props}
    >
      {children}
    </h3>
  );
}

export function TypographyH4({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLHeadingElement>) {
  return (
    <h4
      className={cn(
        "scroll-m-20 font-semibold tracking-tight text-slate-900 dark:text-slate-50 mt-2",
        "text-lg",
        "md:text-xl",
        "2xl:text-2xl",
        className
      )}
      {...props}
    >
      {children}
    </h4>
  );
}

export function TypographyP({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLParagraphElement>) {
  return (
    <p
      className={cn(
        "leading-7 text-slate-700 dark:text-slate-300 [&:not(:first-child)]:mt-6",
        "text-base",
        "2xl:text-lg",
        className
      )}
      {...props}
    >
      {children}
    </p>
  );
}

export function TypographyLead({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLParagraphElement>) {
  return (
    <p
      className={cn(
        "text-slate-500 dark:text-slate-400",
        "text-lg",
        "md:text-xl",
        "2xl:text-2xl",
        className
      )}
      {...props}
    >
      {children}
    </p>
  );
}

export function TypographyLarge({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <p
      className={cn(
        "font-semibold text-slate-900 dark:text-slate-50",
        "text-lg md:text-xl",
        className
      )}
      {...props}
    >
      {children}
    </p>
  );
}

export function TypographySmall({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  return (
    <p className={cn("text-sm font-medium leading-none", className)} {...props}>
      {children}
    </p>
  );
}

export function TypographyMuted({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLParagraphElement>) {
  return (
    <p
      className={cn("text-sm text-slate-500 dark:text-slate-400", className)}
      {...props}
    >
      {children}
    </p>
  );
}

export function TypographyBlockquote({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLQuoteElement>) {
  return (
    <blockquote
      className={cn(
        "mt-6 border-l-2 border-slate-300 pl-6 italic text-slate-800 dark:text-slate-200",
        "md:pl-8",
        className
      )}
      {...props}
    >
      {children}
    </blockquote>
  );
}

export function TypographyInlineCode({
  className,
  children,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  return (
    <code
      className={cn(
        "relative rounded bg-slate-100 px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold text-slate-900 dark:bg-slate-800 dark:text-slate-400",
        className
      )}
      {...props}
    >
      {children}
    </code>
  );
}
