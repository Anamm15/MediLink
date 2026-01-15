"use client";

import { useState, useRef, useEffect } from "react";
import { Plus, Trash2, Search, X, Pill } from "lucide-react";
import { Input } from "@/components/ui/form/Input";
import { Modal, ModalHeader, ModalTitle } from "@/components/ui/Modal";
import { useCreatePrescription } from "../hooks/useCreatePrescription";
import { useSearchMedicineQuery } from "@/hooks/useMedicine";
import { Button } from "@/components/ui/Button";
import { toast } from "sonner";
import { useDoctorIdQuery } from "@/hooks/useDoctor";
import { DEFAULT_LIMIT_QUERY, DEFAULT_PAGE_QUERY } from "@/helpers/constant";
import { MedicineResponse } from "@/types/medicine.type";
// import { useDebounce } from "@/hooks/useDebounce";

interface SelectedMedicine {
  id: string;
  name: string;
  price: number;
  quantity: number;
}

interface PrescriptionCreateModalProps {
  isOpen: boolean;
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>;
  patient_id: string;
  medical_record_id: string;
}

export default function PrescriptionCreateModal({
  isOpen,
  setIsOpen,
  patient_id,
  medical_record_id,
}: PrescriptionCreateModalProps) {
  const [notes, setNotes] = useState("");
  const [selectedMedicines, setSelectedMedicines] = useState<
    SelectedMedicine[]
  >([]);
  const [searchQuery, setSearchQuery] = useState("");
  const [isSearchOpen, setIsSearchOpen] = useState(false);
  const { data: doctorId } = useDoctorIdQuery();
  const { data: searchResultsWithMetadata, isLoading: isSearching } =
    useSearchMedicineQuery(
      searchQuery,
      DEFAULT_PAGE_QUERY,
      DEFAULT_LIMIT_QUERY
    );
  const [medicine, setMedicine] = useState<MedicineResponse[]>([]);
  const { mutateAsync: createPrescription, isPending: isSubmitting } =
    useCreatePrescription(doctorId!);

  useEffect(() => {
    if (!searchResultsWithMetadata) return;
    const { data } = searchResultsWithMetadata;
    setMedicine(data);
  }, [searchResultsWithMetadata]);

  useEffect(() => {
    if (isOpen) {
      setNotes("");
      setSelectedMedicines([]);
      setSearchQuery("");
    }
  }, [isOpen]);

  const handleAddMedicine = (medicine: any) => {
    const isExist = selectedMedicines.find((item) => item.id === medicine.id);

    if (isExist) {
      handleUpdateQuantity(medicine.id, isExist.quantity + 1);
    } else {
      setSelectedMedicines((prev) => [
        ...prev,
        {
          id: medicine.id,
          name: medicine.name,
          price: medicine.base_price || 0,
          quantity: 1,
        },
      ]);
    }
    setSearchQuery("");
    setIsSearchOpen(false);
  };

  const handleRemoveMedicine = (id: string) => {
    setSelectedMedicines((prev) => prev.filter((item) => item.id !== id));
  };

  const handleUpdateQuantity = (id: string, val: number) => {
    if (val < 1) return;
    setSelectedMedicines((prev) =>
      prev.map((item) => (item.id === id ? { ...item, quantity: val } : item))
    );
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (selectedMedicines.length === 0) {
      toast.message("Please add at least one medicine", { duration: 3000 });
      return;
    }
    const payload = {
      patient_id,
      medical_record_id,
      notes,
      medicines: selectedMedicines.map((med) => ({
        medicine_id: med.id,
        quantity: med.quantity,
      })),
    };

    await createPrescription(payload);
    setIsOpen(false);
  };

  return (
    <Modal open={isOpen} setIsOpen={setIsOpen}>
      <ModalHeader>
        <ModalTitle>Create Prescription</ModalTitle>
      </ModalHeader>

      <form onSubmit={handleSubmit} className="space-y-6 mt-4">
        {/* --- Notes --- */}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Notes
          </label>
          <Input
            type="text"
            placeholder="Notes for the use of the prescription"
            value={notes}
            onChange={(e) => setNotes(e.target.value)}
          />
        </div>

        {/* --- Search Medicine --- */}
        <div className="relative z-20">
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Add Medicine
          </label>
          <div className="relative">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-4 h-4" />
            <Input
              type="text"
              placeholder="Type medicine name (e.g. Paracetamol)..."
              value={searchQuery}
              onChange={(e) => {
                setSearchQuery(e.target.value);
                setIsSearchOpen(true);
              }}
              onFocus={() => setIsSearchOpen(true)}
              className="pl-9"
            />
            {searchQuery && (
              <button
                type="button"
                onClick={() => setSearchQuery("")}
                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
              >
                <X className="w-4 h-4" />
              </button>
            )}
          </div>

          {/* --- Search Results Dropdown --- */}
          {isSearchOpen && searchQuery && (
            <div className="absolute w-full mt-1 bg-white border border-gray-200 rounded-md shadow-lg max-h-60 overflow-y-auto">
              {isSearching ? (
                <div className="p-4 text-center text-sm text-gray-500">
                  Searching...
                </div>
              ) : medicine && medicine.length > 0 ? (
                <ul>
                  {medicine.map((med: any) => (
                    <li
                      key={med.id}
                      onClick={() => handleAddMedicine(med)}
                      className="px-4 py-3 hover:bg-slate-50 cursor-pointer flex justify-between items-center border-b last:border-0"
                    >
                      <div>
                        <p className="text-sm font-medium text-gray-800">
                          {med.name}
                        </p>
                        <p className="text-xs text-gray-500">
                          {med.category} • Stock: Available
                        </p>
                      </div>
                      <Plus className="w-4 h-4 text-blue-500" />
                    </li>
                  ))}
                </ul>
              ) : (
                <div className="p-4 text-center text-sm text-gray-500">
                  No medicine found.
                </div>
              )}
            </div>
          )}
        </div>

        {/* --- Selected Medicines --- */}
        <div className="bg-slate-50 rounded-lg p-4 border border-slate-200">
          <h3 className="text-sm font-semibold text-gray-700 mb-3 flex items-center gap-2">
            <Pill className="w-4 h-4" /> Selected Medicines
          </h3>

          {selectedMedicines.length === 0 ? (
            <p className="text-sm text-gray-400 text-center py-4 italic">
              No medicines added yet.
            </p>
          ) : (
            <div className="space-y-3">
              {selectedMedicines.map((med) => (
                <div
                  key={med.id}
                  className="flex items-center justify-between bg-white p-3 rounded-md shadow-sm border border-gray-100"
                >
                  <div className="flex-1">
                    <p className="text-sm font-medium text-gray-800">
                      {med.name}
                    </p>
                    <p className="text-xs text-gray-500">
                      Target ID: {med.id.slice(0, 8)}...
                    </p>
                  </div>

                  <div className="flex items-center gap-3">
                    <div className="flex items-center gap-2">
                      <label className="text-xs text-gray-500">Qty:</label>
                      <Input
                        type="number"
                        min={1}
                        value={med.quantity}
                        onChange={(e) =>
                          handleUpdateQuantity(
                            med.id,
                            parseInt(e.target.value) || 0
                          )
                        }
                        className="w-16 h-8 text-center"
                      />
                    </div>
                    <button
                      type="button"
                      onClick={() => handleRemoveMedicine(med.id)}
                      className="p-2 text-red-500 hover:bg-red-50 rounded-md transition-colors"
                    >
                      <Trash2 className="w-4 h-4" />
                    </button>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* --- Action Buttons --- */}
        <div className="flex justify-end gap-3 pt-4 border-t">
          <Button
            type="button"
            variant="secondary"
            onClick={() => setIsOpen(false)}
            className="w-32"
          >
            Cancel
          </Button>
          <Button
            type="submit"
            variant="primary"
            disabled={isSubmitting || selectedMedicines.length === 0}
          >
            {isSubmitting ? "Saving..." : "Create Prescription"}
          </Button>
        </div>
      </form>
    </Modal>
  );
}
