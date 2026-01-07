package entity

import (
	"time"

	"MediLink/internal/helpers/enum"

	"github.com/google/uuid"
)

type Appointment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index:idx_appointments_patient"`
	DoctorID  uuid.UUID `gorm:"type:uuid;not null;index:idx_appointments_doctor_date"`
	ClinicID  uuid.UUID `gorm:"type:uuid;not null;index:idx_appointments_on_clinic_id"`

	AppointmentDate time.Time `gorm:"type:date;not null"`
	StartTime       string    `gorm:"type:time;not null;index:idx_appointments_doctor_date"`
	EndTime         string    `gorm:"type:time;not null"`

	Status enum.AppointmentStatus `gorm:"type:varchar(20);default:'pending';not null;index:idx_appointments_status"`
	Type   enum.AppointmentType   `gorm:"type:varchar(20);not null"`

	ConsultationFeeSnapshot float64 `gorm:"type:numeric(12,2);not null"`

	QueueNumber *int    `gorm:"type:int;"`
	MeetingLink *string `gorm:"type:text;"`

	SymptomComplaint *string `gorm:"type:text;"`
	DoctorNotes      *string `gorm:"type:text;"`

	Patient        Patient         `gorm:"foreignKey:PatientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Doctor         Doctor          `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Clinic         Clinic          `gorm:"foreignKey:ClinicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MedicalRecords []MedicalRecord `gorm:"foreignKey:AppointmentID"`

	CreatedAt time.Time `gorm:"type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now()"`
}
