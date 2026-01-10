"use client";

import { Button } from "@/components/ui/Button";
import { Input } from "@/components/ui/form/Input";
import { Select } from "@/components/ui/form/Select";
import {
  Modal,
  ModalFooter,
  ModalHeader,
  ModalTitle,
} from "@/components/ui/Modal";
import { OnBoardPatientRequest } from "@/types/patient.type";
import { useState } from "react";
import { useOnBoardPatient } from "../hooks/useUser";
import { UserProfileResponse } from "@/types/user.type";

type PatientCreateModalProps = {
  isOpen: boolean;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  userData: UserProfileResponse;
  setUserData: React.Dispatch<
    React.SetStateAction<UserProfileResponse | undefined>
  >;
};

export default function PatientCreateModal({
  isOpen,
  setIsOpen,
  setUserData,
  userData,
}: PatientCreateModalProps) {
  const [data, setData] = useState<OnBoardPatientRequest>({
    birth_date: "",
    gender: "male",
    identity_number: "",
    blood_type: "",
    weight_kg: 0,
    height_cm: 0,
    allergies: null,
    history_chronic_diseases: null,
    emergency_contact: null,
    insurance_provider: null,
    insurance_number: null,
    occupation: null,
  });
  const { mutate: onBoardPatient } = useOnBoardPatient(
    setIsOpen,
    userData,
    setUserData
  );

  const handleSubmit = () => {
    onBoardPatient(data);
  };

  return (
    <Modal open={isOpen} setIsOpen={setIsOpen}>
      <ModalHeader>
        <ModalTitle>On Board Patient Form</ModalTitle>
      </ModalHeader>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <Input
          label="Identity Number"
          type="text"
          required
          value={data.identity_number}
          placeholder="Identity Number"
          onChange={(e) =>
            setData({ ...data, identity_number: e.target.value })
          }
        />

        <Input
          label="Birth Date"
          type="date"
          required
          value={data.birth_date}
          onChange={(e) => setData({ ...data, birth_date: e.target.value })}
        />

        <Select
          label="Gender"
          required
          value={data.gender}
          onChange={(e) => setData({ ...data, gender: e.target.value })}
        >
          <option value="male">Male</option>
          <option value="female">Female</option>
        </Select>

        <Input
          label="Blood Type"
          type="text"
          value={data.blood_type}
          placeholder="A / B / AB / O"
          onChange={(e) => setData({ ...data, blood_type: e.target.value })}
        />

        <Input
          label="Weight (kg)"
          type="number"
          value={data.weight_kg}
          onChange={(e) =>
            setData({ ...data, weight_kg: Number(e.target.value) })
          }
        />

        <Input
          label="Height (cm)"
          type="number"
          value={data.height_cm}
          onChange={(e) =>
            setData({ ...data, height_cm: Number(e.target.value) })
          }
        />

        <Input
          label="Occupation"
          type="text"
          value={data.occupation ?? ""}
          placeholder="Occupation"
          onChange={(e) => setData({ ...data, occupation: e.target.value })}
        />

        <Input
          label="Emergency Contact"
          type="text"
          value={data.emergency_contact ?? ""}
          placeholder="+6283176759821"
          onChange={(e) =>
            setData({ ...data, emergency_contact: e.target.value })
          }
        />

        <Input
          label="Allergies"
          type="text"
          value={data.allergies ?? ""}
          placeholder="dust mites, pet dander, etc"
          onChange={(e) => setData({ ...data, allergies: e.target.value })}
        />

        <Input
          label="Chronic Diseases"
          type="text"
          value={data.history_chronic_diseases ?? ""}
          placeholder="Cardiovascular, Cancer, etc"
          onChange={(e) =>
            setData({
              ...data,
              history_chronic_diseases: e.target.value,
            })
          }
        />
      </div>

      {/* Insurance Section */}
      <div className="mt-6 border-t pt-6 grid grid-cols-1 md:grid-cols-2 gap-6">
        <Input
          label="Insurance Provider"
          type="text"
          value={data.insurance_provider ?? ""}
          placeholder="Insurance Provider"
          onChange={(e) =>
            setData({ ...data, insurance_provider: e.target.value })
          }
        />

        <Input
          label="Insurance Number"
          type="text"
          value={data.insurance_number ?? ""}
          placeholder="Insurance Number"
          onChange={(e) =>
            setData({ ...data, insurance_number: e.target.value })
          }
        />
      </div>
      <ModalFooter>
        <Button onClick={handleSubmit} className="w-full">
          Submit
        </Button>
      </ModalFooter>
    </Modal>
  );
}
