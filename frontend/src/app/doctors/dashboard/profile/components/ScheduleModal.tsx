"use client";

import { Button } from "@/components/ui/Button";
import { Input } from "@/components/ui/form/Input";
import { Select } from "@/components/ui/form/Select";
import { Modal, ModalHeader, ModalTitle } from "@/components/ui/Modal";
import { DAYS } from "@/helpers/constant";
import {
  CreateScheduleRequest,
  Slot,
  UpdateScheduleRequest,
} from "@/types/schedule.type";
import { useEffect, useState } from "react";
import { useCreateSchedule, useUpdateSchedule } from "../hooks/useSchedule";

type ScheduleModalProps = {
  day: string;
  isEdit: boolean;
  setIsEdit: React.Dispatch<React.SetStateAction<boolean>>;
  slot?: Slot | null;
  setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>;
};

export default function ScheduleModal({
  day,
  isEdit = false,
  slot,
  setIsEdit,
  setIsModalOpen,
}: ScheduleModalProps) {
  const [temporarySchedule, setTemporarySchedule] = useState<
    CreateScheduleRequest | UpdateScheduleRequest
  >({
    day_of_week: day,
    start_time: "",
    end_time: "",
    consultation_fee: 0,
    max_quota: 1,
    is_active: true,
  });

  useEffect(() => {
    if (isEdit && slot) {
      setTemporarySchedule({
        day_of_week: day,
        start_time: slot.startTime,
        end_time: slot.endTime,
        consultation_fee: slot.consultation_fee,
        max_quota: slot.max_quota,
        is_active: slot.isActive,
      });
    }
  }, []);

  const { mutateAsync: createSchedule } = useCreateSchedule();
  const { mutateAsync: updateSchedule } = useUpdateSchedule();

  const handleSubmit = async () => {
    if (isEdit) {
      await updateSchedule({
        id: slot?.id as string,
        payload: temporarySchedule as UpdateScheduleRequest,
      }).then(() => {
        setIsModalOpen(false);
        setIsEdit(false);
      });
    } else {
      await createSchedule(temporarySchedule as CreateScheduleRequest).then(
        () => setIsModalOpen(false)
      );
    }
  };

  return (
    <Modal open setIsOpen={setIsModalOpen} className="max-w-sm">
      <ModalHeader>
        <ModalTitle>{isEdit ? "Edit Schedule" : "Add Schedule"}</ModalTitle>
      </ModalHeader>

      <form className="space-y-4 ">
        <Select
          name="day"
          label="Day"
          required
          disabled
          value={temporarySchedule.day_of_week}
          onChange={(e) =>
            setTemporarySchedule({
              ...temporarySchedule,
              day_of_week: e.target.value,
            })
          }
        >
          {DAYS.map((day) => (
            <option key={day} value={day}>
              {day}
            </option>
          ))}
        </Select>

        <div className="flex gap-2">
          <Input
            type="time"
            name="start time"
            label="Start Time"
            required
            value={temporarySchedule.start_time}
            onChange={(e) => {
              setTemporarySchedule({
                ...temporarySchedule,
                start_time: e.target.value,
              });
            }}
          />

          <Input
            type="time"
            name="End Time"
            label="End Time"
            required
            value={temporarySchedule.end_time}
            onChange={(e) => {
              setTemporarySchedule({
                ...temporarySchedule,
                end_time: e.target.value,
              });
            }}
          />
        </div>

        <Input
          type="text"
          name="Consultation Fee"
          label="Consultation Fee"
          required
          value={temporarySchedule.consultation_fee}
          onChange={(e) => {
            setTemporarySchedule({
              ...temporarySchedule,
              consultation_fee: Number(e.target.value),
            });
          }}
        />

        <Input
          type="text"
          name="Max Quota"
          label="Max Quota"
          required
          value={temporarySchedule.max_quota}
          onChange={(e) => {
            setTemporarySchedule({
              ...temporarySchedule,
              max_quota: Number(e.target.value),
            });
          }}
        />

        <div className="flex justify-end gap-2">
          <Button
            type="button"
            variant="outline"
            className="w-28"
            onClick={() => setIsModalOpen(false)}
          >
            Cancel
          </Button>
          <Button
            type="submit"
            variant="primary"
            className="w-28"
            onClick={handleSubmit}
          >
            Submit
          </Button>
        </div>
      </form>
    </Modal>
  );
}
