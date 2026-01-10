"use client";

import { Edit3, MailWarning, Plus, Stethoscope } from "lucide-react";
import React, { useState } from "react";
import EditableField from "./InfoItem";
import { PatientResponse } from "@/types/patient.type";
import { TypographyP, TypographySmall } from "@/components/ui/Typography";
import { Button } from "@/components/ui/Button";
import { UserProfileResponse } from "@/types/user.type";
import { useSendEmailVerification } from "../hooks/useUser";
import EmailVerificationModal from "./EmailVerificationModal";
import PatientCreateModal from "./PatientCreateModal";
import { useUpdatePatient } from "../hooks/usePatient";

type DetailPatientProps = {
  data: UserProfileResponse;
  setData: React.Dispatch<
    React.SetStateAction<UserProfileResponse | undefined>
  >;
};

export default function DetailPatient({ setData, data }: DetailPatientProps) {
  const [isEditing, setIsEditing] = useState(false);
  const [isVerificationModalOpen, setIsVerificationModalOpen] = useState(false);
  const [temporaryData, setTemporaryData] = useState<UserProfileResponse>(data);
  const [isOnBoardPatientModalOpen, setIsOnBoardPatientModalOpen] =
    useState(false);
  const { mutate: sendEmailVerification } = useSendEmailVerification(
    setIsVerificationModalOpen
  );
  const { mutate: updatePatient } = useUpdatePatient(setIsEditing);

  const updateFieldPatient = <K extends keyof PatientResponse>(
    key: K,
    value: PatientResponse[K]
  ) => {
    if (!temporaryData.patient) return;

    setTemporaryData({
      ...temporaryData,
      patient: {
        ...temporaryData.patient,
        [key]: value,
      },
    });
  };

  const handleSendEmailVerification = () => {
    sendEmailVerification();
  };

  const handleUpdatePatient = () => {
    if (!data.patient) return;
    updatePatient(temporaryData.patient as PatientResponse);
    setData(temporaryData);
  };

  const handleCancel = () => {
    setTemporaryData(data);
    setIsEditing(false);
  };

  return (
    <div className="lg:col-span-2 space-y-6">
      {!data.user.is_verified && (
        <div className="bg-amber-50 border border-amber-200 p-4 rounded-xl flex gap-3 items-start">
          <MailWarning className="w-5 h-5 text-amber-600 mt-0.5" />
          <div>
            <p className="text-sm font-bold text-amber-900">
              Email Not Verified
            </p>
            <p className="text-xs text-amber-700 mt-1">
              Please verify your email to complete the patient's medical
              details.
            </p>
            <button
              onClick={handleSendEmailVerification}
              className="mt-2 text-xs font-bold text-amber-700 underline underline-offset-2 cursor-pointer"
            >
              Send Verification Email
            </button>
          </div>
        </div>
      )}

      <div
        className={`bg-white rounded-2xl border border-gray-200 shadow-sm overflow-hidden ${
          !data.user.is_verified ? "opacity-60 grayscale" : ""
        }`}
      >
        <div className="p-6 border-b border-gray-100 flex items-center justify-between">
          <h3 className="font-bold text-cyan-600 flex items-center gap-2">
            <Stethoscope className="w-5 h-5 text-cyan-500" /> Patient Medical
            Information
          </h3>
          {data.patient && data.user.is_verified && (
            <button
              onClick={() => setIsEditing(!isEditing)}
              className="text-cyan-600 hover:text-cyan-700 text-sm font-bold flex items-center gap-1"
            >
              <Edit3 className="w-4 h-4" /> Edit
            </button>
          )}
        </div>

        <div className="p-6">
          {!temporaryData.patient ? (
            <div className="text-center py-10 space-y-4">
              <div
                onClick={() => setIsOnBoardPatientModalOpen(true)}
                className="bg-slate-100 w-16 h-16 rounded-full flex items-center justify-center mx-auto text-slate-400 cursor-pointer"
              >
                <Plus className="w-8 h-8" />
              </div>
              <div className="max-w-xs mx-auto">
                <TypographyP className="font-bold text-gray-800">
                  There is no patient data yet
                </TypographyP>
                <TypographyP className="text-sm text-gray-500 mt-1">
                  Complete your medical data to facilitate the consultation
                  process with the doctor.
                </TypographyP>
              </div>
              <button
                onClick={() => setIsOnBoardPatientModalOpen(true)}
                disabled={!data.user.is_verified}
                className="inline-flex items-center gap-2 px-6 py-2.5 bg-cyan-500 text-white font-bold rounded-xl hover:bg-cyan-600 disabled:bg-gray-300 disabled:cursor-not-allowed transition-all cursor-pointer"
              >
                <Plus className="w-5 h-5" /> Complete the Patient Profile
              </button>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 gap-x-8 gap-y-2">
              <div>
                <label className="text-sm text-gray-500">Identity Number</label>
                <p className="text-gray-800 font-semibold">
                  {temporaryData.patient.identity_number}
                </p>
              </div>

              <div>
                <label className="text-sm text-gray-500">Birth Date</label>
                <EditableField
                  value={temporaryData.patient.birth_date}
                  isEditing={isEditing}
                  placeholder="Birth Date"
                  onChange={(value) => updateFieldPatient("birth_date", value)}
                />
              </div>

              <div>
                <label className="text-sm text-gray-500">Weight (kg)</label>
                <EditableField
                  value={temporaryData.patient.weight_kg}
                  isEditing={isEditing}
                  placeholder="Weight"
                  onChange={(value) =>
                    updateFieldPatient("weight_kg", Number(value))
                  }
                />
              </div>

              <div>
                <label className="text-sm text-gray-500">Height (cm)</label>
                <EditableField
                  value={temporaryData.patient.height_cm}
                  isEditing={isEditing}
                  placeholder="Height"
                  onChange={(value) =>
                    updateFieldPatient("height_cm", Number(value))
                  }
                />
              </div>

              <div>
                <label className="text-sm text-gray-500">Gender</label>
                <EditableField
                  value={temporaryData.patient.gender}
                  isEditing={isEditing}
                  placeholder="Gender"
                  onChange={(value) => updateFieldPatient("gender", value)}
                />
              </div>

              <div>
                <label className="text-sm text-gray-500">Blood Type</label>
                <EditableField
                  value={temporaryData.patient.blood_type}
                  isEditing={isEditing}
                  placeholder="Blood Type"
                  onChange={(value) => updateFieldPatient("blood_type", value)}
                />
              </div>

              <div>
                <label className="text-sm text-gray-500">Occupation</label>
                <EditableField
                  value={temporaryData.patient.occupation}
                  isEditing={isEditing}
                  placeholder="Occupation"
                  onChange={(value) => updateFieldPatient("occupation", value)}
                />
              </div>

              <div>
                <label className="text-sm text-gray-500">Allergies</label>
                <EditableField
                  value={temporaryData.patient.allergies}
                  isEditing={isEditing}
                  placeholder="Allergies"
                  onChange={(value) => updateFieldPatient("allergies", value)}
                />
              </div>

              <div>
                <label className="text-sm text-gray-500">
                  Chronic Conditions
                </label>
                <EditableField
                  value={temporaryData.patient.history_chronic_diseases}
                  isEditing={isEditing}
                  placeholder="Chronic Conditions"
                  onChange={(value) =>
                    updateFieldPatient("history_chronic_diseases", value)
                  }
                />
              </div>
              <div className="md:col-span-2 pt-4 border-t border-gray-100">
                <TypographySmall className="font-bold text-gray-800 mb-4">
                  Additional Information
                </TypographySmall>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label className="text-sm text-gray-500">
                      Emergency Contact
                    </label>
                    <EditableField
                      value={temporaryData.patient.emergency_contact}
                      isEditing={isEditing}
                      placeholder="Emergency Contact"
                      onChange={(value) =>
                        updateFieldPatient("emergency_contact", value)
                      }
                    />
                  </div>

                  <div>
                    <label className="text-sm text-gray-500">
                      Insurance Provider
                    </label>
                    <EditableField
                      value={temporaryData.patient.insurance_provider}
                      isEditing={isEditing}
                      placeholder="Insurance Provider"
                      onChange={(value) =>
                        updateFieldPatient("insurance_provider", value)
                      }
                    />
                  </div>

                  <div>
                    <label className="text-sm text-gray-500">
                      Insurance Number
                    </label>
                    <EditableField
                      value={temporaryData.patient.insurance_number}
                      isEditing={isEditing}
                      placeholder="Insurance Number"
                      onChange={(value) =>
                        updateFieldPatient("insurance_number", value)
                      }
                    />
                  </div>
                </div>
                {isEditing && (
                  <div className="flex justify-end gap-2 mt-4">
                    <Button onClick={handleCancel} className="w-32">
                      Cancel
                    </Button>
                    <Button onClick={handleUpdatePatient} className="w-32">
                      Save
                    </Button>
                  </div>
                )}
              </div>
            </div>
          )}
        </div>
      </div>
      {isVerificationModalOpen && (
        <EmailVerificationModal
          setIsOpen={setIsVerificationModalOpen}
          isOpen
          data={data}
          setData={setData}
        />
      )}

      {isOnBoardPatientModalOpen && (
        <PatientCreateModal
          isOpen
          setIsOpen={setIsOnBoardPatientModalOpen}
          userData={data}
          setUserData={setData}
        />
      )}
    </div>
  );
}
