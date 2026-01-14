import {
  createSchedule,
  deleteSchedule,
  getSchedules,
  updateSchedule,
  updateStatusSchedule,
} from "@/services/doctor.service";
import {
  DoctorScheduleResponse,
  formattedSchedule,
  UpdateScheduleRequest,
  UpdateStatusScheduleRequest,
} from "@/types/schedule.type";
import { useQuery, useMutation } from "@tanstack/react-query";
import { toast } from "sonner";

export function useSchedulesQuery(doctor_id: string) {
  return useQuery<DoctorScheduleResponse[]>({
    queryKey: ["schedule"],
    queryFn: () => getSchedules(doctor_id),
  });
}

export function useCreateSchedule() {
  return useMutation({
    mutationFn: createSchedule,

    onMutate: () => {
      const toastId = toast.loading("Creating schedule...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      toast.success("Schedule created successfully", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to create schedule", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

export function useUpdateSchedule() {
  return useMutation({
    mutationFn: ({
      id,
      payload,
    }: {
      id: string;
      payload: UpdateScheduleRequest;
    }) => updateSchedule(id, payload),

    onMutate: () => {
      const toastId = toast.loading("Updating schedule...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      toast.success("Schedule updated successfully", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to update schedule", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

export function useUpdateStatusSchedule(
  setSchedules: React.Dispatch<React.SetStateAction<formattedSchedule[]>>,
  slotId: string
) {
  return useMutation({
    mutationFn: async ({
      id,
      is_active,
    }: {
      id: string;
      is_active: boolean;
    }) => {
      const data = await updateStatusSchedule(id, is_active);
      return data;
    },

    onMutate: () => {
      const toastId = toast.loading("Updating schedule status...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      setSchedules((prev) =>
        prev.map((day) => ({
          ...day,
          slots: day.slots.map((slot) =>
            slot.id === slotId ? { ...slot, isActive: !slot.isActive } : slot
          ),
        }))
      );

      toast.success("Schedule status updated successfully", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to update schedule status", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}

export function useDeleteSchedule(
  schedules: formattedSchedule[],
  setData: React.Dispatch<React.SetStateAction<formattedSchedule[]>>,
  slotId: string
) {
  return useMutation({
    mutationFn: (id: string) => deleteSchedule(id),

    onMutate: () => {
      const toastId = toast.loading("Deleting schedule...");
      return { toastId };
    },

    onSuccess: (data, variables, context) => {
      console.log("slot id: " + slotId);

      setData(
        schedules.map((day) => ({
          ...day,
          slots: day.slots.filter((slot) => slot.id !== slotId),
        }))
      );

      console.log(schedules);
      toast.success("Schedule deleted successfully", {
        id: context?.toastId,
        duration: 3000,
      });
    },

    onError: (error, variables, context) => {
      toast.error("Failed to delete schedule", {
        id: context?.toastId,
        duration: 3000,
      });
      console.log(error);
    },
  });
}
