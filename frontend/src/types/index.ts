export type Doctor = {
  // ... properti dari halaman daftar sebelumnya
  id: string;
  name: string;
  specialty: string;
  avatarUrl: string;
  rating: number;
  reviews: number;
  isOnline: boolean;
  nextAvailable: string;
};

export type Education = {
  degree: string;
  university: string;
  year: string;
};

export type Experience = {
  position: string;
  hospital: string;
  period: string;
};

export type PatientReview = {
  id: string;
  author: string;
  avatarUrl: string;
  rating: number;
  comment: string;
  date: string;
};

// Tipe data lengkap untuk halaman detail
export type DoctorDetail = Doctor & {
  yearsOfExperience: number;
  patientCount: number;
  bio: string;
  education: Education[];
  experience: Experience[];
  reviewsList: PatientReview[];
  consultationFee: number;
  clinic: {
    name: string;
    address: string;
  };
};

export type Medicine = {
  id: string;
  name: string;
  dosage: string; // e.g., "500 mg"
  imageUrl: string;
  price: number;
  quantity: number;
};

export type Prescription = {
  id: string;
  doctorName: string;
  doctorSpecialty: string;
  date: string; // e.g., "19 Sep 2025"
  medicines: Medicine[];
  isRedeemed: boolean;
};
