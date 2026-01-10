import Image from "next/image";
import { KeyRound, Camera, ShieldCheck, ShieldAlert } from "lucide-react";

export const ProfileHeader = ({ user }: { user: any }) => {
  return (
    <div className="bg-white rounded-2xl shadow-sm border border-gray-200 overflow-hidden">
      <div className="h-24 bg-gradient-to-r from-blue-500 to-cyan-600"></div>
      <div className="px-8 pb-8">
        <div className="relative -mt-12 flex flex-col md:flex-row items-end gap-6">
          <div className="relative">
            <Image
              src={user.avatar_url}
              alt={user.name}
              width={120}
              height={120}
              className="rounded-full border-4 border-white shadow-md bg-white"
            />
            <button className="absolute bottom-1 right-1 p-2 bg-white rounded-full shadow-border hover:bg-gray-50 border border-gray-200 transition-colors">
              <Camera className="w-4 h-4 text-gray-600" />
            </button>
          </div>

          <div className="flex-grow pb-2">
            <div className="flex items-center gap-2">
              <h2 className="text-2xl font-bold text-gray-900">{user.name}</h2>
              {user.is_verified ? (
                <span className="flex items-center gap-1 text-xs font-bold bg-green-100 text-green-700 px-2 py-1 rounded-full">
                  <ShieldCheck className="w-3 h-3" /> Verified
                </span>
              ) : (
                <span className="flex items-center gap-1 text-xs font-bold bg-amber-100 text-amber-700 px-2 py-1 rounded-full">
                  <ShieldAlert className="w-3 h-3" />
                  Not Verified
                </span>
              )}
            </div>
            <p className="text-gray-500">{user.email}</p>
          </div>

          <div className="flex gap-3 pb-2">
            <button className="flex items-center gap-2 px-4 py-2 text-sm font-semibold text-gray-700 bg-white border border-gray-200 rounded-lg hover:bg-gray-50 transition-all">
              <KeyRound className="w-4 h-4" /> Password
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};
