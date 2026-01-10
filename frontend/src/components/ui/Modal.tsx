"use client";

import * as React from "react";
import { createPortal } from "react-dom";
import { X } from "lucide-react";
import { cn } from "@/lib/utils";
import { Button } from "./Button";

interface ModalProps {
  open: boolean;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  children: React.ReactNode;
  className?: string;
  preventCloseOutside?: boolean;
}

export function Modal({
  open,
  setIsOpen,
  children,
  className,
  preventCloseOutside = false,
}: ModalProps) {
  const [mounted, setMounted] = React.useState(false);

  React.useEffect(() => {
    setMounted(true);
  }, []);

  React.useEffect(() => {
    if (open) {
      document.body.style.overflow = "hidden";

      const handleEsc = (e: KeyboardEvent) => {
        if (e.key === "Escape") setIsOpen(false);
      };
      window.addEventListener("keydown", handleEsc);

      return () => {
        document.body.style.overflow = "unset";
        window.removeEventListener("keydown", handleEsc);
      };
    }
  }, [open, setIsOpen]);

  if (!mounted || !open) return null;

  return createPortal(
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-6">
      <div
        className="fixed inset-0 bg-slate-950/50 backdrop-blur-sm transition-opacity animate-in fade-in duration-200"
        onClick={() => !preventCloseOutside && setIsOpen(false)}
      />

      {/* Modal Content Wrapper */}
      <div
        role="dialog"
        aria-modal="true"
        className={cn(
          "relative w-full max-w-lg transform overflow-hidden rounded-xl bg-white p-6 text-left shadow-xl transition-all animate-in zoom-in-95 duration-200 sm:rounded-2xl",
          className
        )}
      >
        <button
          onClick={() => setIsOpen(false)}
          className="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-white transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-slate-400 focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-slate-100 data-[state=open]:text-slate-500"
        >
          <X className="h-4 w-4" />
          <span className="sr-only">Close</span>
        </button>

        {children}
      </div>
    </div>,
    document.body
  );
}

export function ModalHeader({
  className,
  children,
}: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      className={cn(
        "flex flex-col space-y-1.5 text-center sm:text-left mb-4",
        className
      )}
    >
      {children}
    </div>
  );
}

export function ModalTitle({
  className,
  children,
}: React.HTMLAttributes<HTMLHeadingElement>) {
  return (
    <h3
      className={cn(
        "text-lg font-semibold leading-none tracking-tight text-slate-900",
        className
      )}
    >
      {children}
    </h3>
  );
}

export function ModalDescription({
  className,
  children,
}: React.HTMLAttributes<HTMLParagraphElement>) {
  return <p className={cn("text-sm text-slate-500", className)}>{children}</p>;
}

export function ModalFooter({
  className,
  children,
}: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      className={cn(
        "flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2 mt-6",
        className
      )}
    >
      {children}
    </div>
  );
}
