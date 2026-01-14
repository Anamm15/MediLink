import { Mic, MicOff, Video, VideoOff, PhoneOff } from "lucide-react";
import Image from "next/image";

export const VideoInterface = () => {
  return (
    <div className="flex-grow flex flex-col bg-slate-900 h-full relative text-white items-center justify-center rounded-lg">
      <div className="absolute inset-0 flex items-center justify-center">
        <div className="text-center">
          <Image
            src="https://i.pravatar.cc/150?u=adinda"
            alt="Dr. Adinda"
            width={120}
            height={120}
            className="rounded-full mx-auto"
          />
          <p className="mt-4 font-semibold">Dr. Adinda Melati, Sp.A</p>
          <p className="text-sm text-gray-400">Connecting...</p>
        </div>
      </div>

      {/* Patient's Video (Picture-in-Picture) */}
      <div className="absolute bottom-6 right-6 w-48 h-36 bg-slate-800 rounded-lg border-2 border-slate-700 overflow-hidden">
        <div className="w-full h-full flex items-center justify-center text-xs text-gray-400">
          Your Video
        </div>
      </div>

      {/* Controls */}
      <div className="absolute bottom-6 left-1/2 -translate-x-1/2 flex items-center gap-4 bg-black/50 p-3 rounded-full">
        <button className="p-3 bg-white/10 rounded-full hover:bg-white/20">
          <Mic className="w-6 h-6" />
        </button>
        <button className="p-3 bg-white/10 rounded-full hover:bg-white/20">
          <Video className="w-6 h-6" />
        </button>
        <button className="p-3 bg-red-500 rounded-full hover:bg-red-600">
          <PhoneOff className="w-6 h-6" />
        </button>
      </div>
    </div>
  );
};
