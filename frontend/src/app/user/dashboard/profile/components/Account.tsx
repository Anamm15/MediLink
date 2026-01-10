"use client";
import { Edit3, UserIcon } from "lucide-react";
import React, { useState } from "react";
import { Input } from "@/components/ui/form/Input";
import { UserProfileResponse, UserResponse } from "@/types/user.type";
import { Button } from "@/components/ui/Button";
import EditableField from "./InfoItem";
import { useUpdateUser } from "../hooks/useUser";

type AccountInformationProps = {
  data: UserProfileResponse;
  setData: React.Dispatch<
    React.SetStateAction<UserProfileResponse | undefined>
  >;
};

export function AccountInformation({ data, setData }: AccountInformationProps) {
  const [isEditing, setIsEditing] = useState(false);
  const [temporaryData, setTemporaryData] = useState<UserProfileResponse>(data);
  const { mutate: updateUser } = useUpdateUser(setIsEditing);

  const updateFieldUser = <K extends keyof UserResponse>(
    key: K,
    value: UserResponse[K]
  ) => {
    if (!temporaryData.user) return;

    setTemporaryData({
      ...temporaryData,
      user: {
        ...temporaryData.user,
        [key]: value,
      },
    });
  };

  const handleupdateUser = () => {
    updateUser(temporaryData.user);
    setData(temporaryData);
  };

  const handleCancel = () => {
    setTemporaryData(data);
    setIsEditing(false);
  };

  return (
    <div className="lg:col-span-1 space-y-6">
      <div className="bg-white p-6 rounded-2xl border border-gray-200 shadow-sm">
        <div className="flex items-center justify-between mb-6">
          <h3 className="font-bold text-cyan-600 flex items-center gap-2">
            <UserIcon className="w-5 h-5 text-cyan-500" /> Account
          </h3>
          <button
            onClick={() => setIsEditing(!isEditing)}
            className="text-cyan-600 hover:text-cyan-700 text-sm font-bold flex items-center gap-1"
          >
            <Edit3 className="w-4 h-4" /> Edit
          </button>
        </div>

        <div className="space-y-2">
          <div>
            <label className="text-sm text-gray-500">Full Name</label>
            <EditableField
              value={temporaryData.user.name}
              isEditing={isEditing}
              placeholder="Full Name"
              onChange={(value) => updateFieldUser("name", value)}
            />
          </div>
          <div>
            <label className="text-sm text-gray-500">Email</label>
            <EditableField
              value={temporaryData.user.email}
              isEditing={isEditing}
              placeholder="yourmail@gmail.com"
              onChange={(value) => updateFieldUser("email", value)}
            />
          </div>
          <div>
            <label className="text-sm text-gray-500">Phone Number</label>
            <EditableField
              value={temporaryData.user.phone_number}
              isEditing={isEditing}
              placeholder="+6283111056728"
              onChange={(value) => updateFieldUser("phone_number", value)}
            />
          </div>
          {isEditing && (
            <div className="flex gap-2">
              <Button type="button" className="flex-1" onClick={handleCancel}>
                Cancel
              </Button>
              <Button
                type="button"
                className="flex-1"
                onClick={handleupdateUser}
              >
                Save
              </Button>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
