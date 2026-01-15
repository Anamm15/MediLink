"use client";

import { Button } from "@/components/ui/Button";
import { Input } from "@/components/ui/form/Input";
import { Textarea } from "@/components/ui/form/TextArea";
import { Modal, ModalHeader, ModalTitle } from "@/components/ui/Modal";
import { getCurrentDate } from "@/helpers/datetime";
import { MedicalRecordCreateRequest } from "@/types/medical_record.type";
import React, { useState } from "react";
import { useCreateMedicalRecord } from "../hooks/useMedicalRecord";
import { useDoctorIdQuery } from "@/hooks/useDoctor";

type MedicalRecordModalProps = {
  isOpen: boolean;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  patient_id: string;
  appointment_id: string;
};

export default function MedicalRecordModal({
  isOpen,
  setIsOpen,
  patient_id,
  appointment_id,
}: MedicalRecordModalProps) {
  const [tempData, setTempData] = useState<MedicalRecordCreateRequest>({
    patient_id,
    appointment_id,
    date: getCurrentDate(),
    title: "",
    subjective: "",
    objective: "",
    assessment: "",
    plan: "",
  });
  const { data: doctor_id } = useDoctorIdQuery();
  const { mutateAsync: createMedicalRecord } = useCreateMedicalRecord(
    doctor_id!
  );

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log(tempData);
    await createMedicalRecord(tempData).then(() => setIsOpen(false));
  };

  return (
    <Modal open={isOpen} setIsOpen={setIsOpen}>
      <ModalHeader>
        <ModalTitle>Create Medical Record</ModalTitle>
      </ModalHeader>

      <form onSubmit={handleSubmit} className="space-y-4">
        <Input
          type="text"
          label="Title"
          placeholder="Result of the consultation"
          required
          value={tempData.title}
          onChange={(e) => setTempData({ ...tempData, title: e.target.value })}
        />

        <Textarea
          label="Subjective"
          placeholder="Patient complaints, symptoms, history"
          value={tempData.subjective}
          onChange={(e) =>
            setTempData({ ...tempData, subjective: e.target.value })
          }
        />

        <Textarea
          label="Objective"
          placeholder="Physical examination, vital signs, lab results"
          value={tempData.objective}
          onChange={(e) =>
            setTempData({ ...tempData, objective: e.target.value })
          }
        />

        <Textarea
          label="Assessment"
          placeholder="Diagnosis or clinical assessment"
          value={tempData.assessment}
          onChange={(e) =>
            setTempData({ ...tempData, assessment: e.target.value })
          }
        />

        <Textarea
          label="Plan"
          placeholder="Treatment plan, medication, follow-up"
          value={tempData.plan}
          onChange={(e) => setTempData({ ...tempData, plan: e.target.value })}
        />
        <div className="flex justify-end gap-3">
          <Button
            type="button"
            variant="secondary"
            className="w-32"
            onClick={() => setIsOpen(false)}
          >
            Cancel
          </Button>
          <Button type="submit" className="w-32">
            Submit
          </Button>
        </div>
      </form>
    </Modal>
  );
}
