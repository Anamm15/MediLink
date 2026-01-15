export type DoctorMinimumResponse = {
  id: string;
  name: string;
  specialization: string;
  avatar_url?: string;
};

export type DoctorClinicRepsonse = {
  id: string;
  name: string;
  address: string;
  city: string;
  is_active: string;
};

export type DoctorProfileResponse = {
  id: string;
  name: string;
  email: string;
  phone_number: string;
  specialization: string;
  license_number: string;
  bio: string;
  avatar_url: string;
  experience_years: number;
  education: {
    year: string;
    degree: string;
    institution: string;
  }[];
  review_count: number;
  rating_total: number;
  rating_count: number;
  clinic: DoctorClinicRepsonse[];
};

export type DoctorUpdateRequest = {
  specialization: string;
  license_number: string;
  bio: string;
  experience_years: number;
  education: string;
};
