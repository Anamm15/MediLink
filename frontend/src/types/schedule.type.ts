import { DAYS } from "@/helpers/constant";

export type DoctorScheduleResponse = {
  id: string;
  day_of_week: string;
  start_time: string;
  end_time: string;
  consultation_fee: number;
  is_active: boolean;
  max_quota: number;
  type: string;
};

export type CreateScheduleRequest = {
  clinic_id?: string;
  day_of_week: string;
  start_time: string;
  end_time: string;
  consultation_fee: number;
  max_quota: number;
  is_active?: boolean;
  type: string;
};

export type UpdateScheduleRequest = {
  day_of_week?: string;
  start_time?: string;
  end_time?: string;
  consultation_fee?: number;
  max_quota?: number;
  type?: string;
};

export type UpdateStatusScheduleRequest = {
  is_active: boolean;
};

export type Slot = {
  id: string;
  isActive: boolean;
  startTime: string;
  endTime: string;
  consultation_fee: number;
  max_quota: number;
  type: string;
};

export type formattedSchedule = {
  day: string;
  slots: Slot[];
};

export function formatSchedule(schedules: DoctorScheduleResponse[]) {
  const formattedSchedules: formattedSchedule[] = [];

  for (const day of DAYS) {
    const scheduleDay = schedules.filter(
      (schedule) => schedule.day_of_week === day
    );
    const slots = scheduleDay.map((schedule) => ({
      id: schedule.id,
      startTime: schedule.start_time,
      endTime: schedule.end_time,
      consultation_fee: schedule.consultation_fee,
      max_quota: schedule.max_quota,
      isActive: schedule.is_active,
      type: schedule.type,
    }));
    formattedSchedules.push({ day, slots });
  }
  return formattedSchedules;
}
