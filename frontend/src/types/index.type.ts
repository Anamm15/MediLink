export type Doctor = {
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
  dosage: string;
  imageUrl: string;
  price: number;
  quantity: number;
};

export type Prescription = {
  id: string;
  doctorName: string;
  doctorSpecialty: string;
  date: string;
  medicines: Medicine[];
  isRedeemed: boolean;
};
