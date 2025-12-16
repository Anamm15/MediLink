"use client";

import React from "react";

interface HelperInputTextProps {
    children: React.ReactNode;
}

const HelperInputText: React.FC<HelperInputTextProps> = ({ children }) => {
    return <p className="text-sm text-red-600">{children}</p>;
};

export default HelperInputText;
