"use client";
import { useState } from "react";
import { MessageSquare, Video } from "lucide-react";
import { ChatInterface } from "../components/ChatInterface";
import { VideoInterface } from "../components/VideoInterface";
import { SessionInfoPanel } from "../components/SessionInfoPanel";
import { Doctor } from "@/types/index.type";
import { Navbar } from "@/components/layout/Navbar";

// Data Mock
const mockDoctor: Doctor = {
  id: "1",
  name: "Dr. Adinda Melati, Sp.A",
  specialty: "Dokter Anak",
  avatarUrl: "https://i.pravatar.cc/150?u=adinda",
  rating: 4.9,
  reviews: 128,
  isOnline: true,
  nextAvailable: "",
};

export default function ConsultationPage({
  params,
}: {
  params: { id: string };
}) {
  const [activeView, setActiveView] = useState<"chat" | "video">("chat");

  return (
    <>
      <Navbar />
      <main className="flex h-screen bg-slate-50 font-sans">
        {/* Kolom Kiri - Riwayat Medis (Bisa dibuat collapsible nanti) */}
        {/* <MedicalRecordSidebar /> */}

        <div className="flex-grow flex flex-col">
          {/* Header dengan Toggle */}
          <header className="flex-shrink-0 bg-white border-b border-gray-200 p-4 flex justify-between items-center">
            <div>
              <h1 className="font-bold text-gray-800">Ruang Konsultasi</h1>
              <p className="text-xs text-gray-500">ID Sesi: {params.id}</p>
            </div>
            <div className="flex items-center p-1 bg-slate-100 rounded-lg">
              <button
                onClick={() => setActiveView("chat")}
                className={`flex items-center gap-2 px-4 py-1.5 text-sm rounded-md transition-colors ${
                  activeView === "chat"
                    ? "bg-white shadow-sm text-cyan-600 font-semibold"
                    : "text-gray-500"
                }`}
              >
                <MessageSquare className="w-4 h-4" /> Chat
              </button>
              <button
                onClick={() => setActiveView("video")}
                className={`flex items-center gap-2 px-4 py-1.5 text-sm rounded-md transition-colors ${
                  activeView === "video"
                    ? "bg-white shadow-sm text-cyan-600 font-semibold"
                    : "text-gray-500"
                }`}
              >
                <Video className="w-4 h-4" /> Video Call
              </button>
            </div>
          </header>

          {/* Area Interaksi Utama */}
          <div className="flex-grow p-4 overflow-hidden">
            {activeView === "chat" ? <ChatInterface /> : <VideoInterface />}
          </div>
        </div>

        {/* Kolom Kanan - Info Sesi */}
        <SessionInfoPanel doctor={mockDoctor} durationInMinutes={30} />
      </main>
    </>
  );
}
