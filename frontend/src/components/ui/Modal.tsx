import React from "react";

type ModalProps = {
  title: string;
  setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>;
  children: React.ReactNode;
};

export function Modal({ title, setIsModalOpen, children }: ModalProps) {
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4 sm:p-6">
      <div
        className="absolute inset-0 bg-stone-900/20 backdrop-blur-sm"
        onClick={() => setIsModalOpen(false)}
      ></div>
      <div className="relative w-full max-w-lg bg-white rounded-3xl p-6 animate-in zoom-in-95 duration-300 transition-all">
        <h3 className="font-bold text-xl mb-2">{title}</h3>
        {children}
        <button
          onClick={() => setIsModalOpen(false)}
          className="w-full bg-stone-900 text-white py-2.5 rounded-xl mt-4 cursor-pointer"
        >
          Close
        </button>
      </div>
    </div>
  );
}
