import { searchMedicine } from "@/services/medicine.service";
import { useQuery } from "@tanstack/react-query";

export function useSearchMedicineQuery(name: string) {
  return useQuery({
    queryKey: ["medicine", name],
    queryFn: () => searchMedicine(name),
    enabled: !!name,
    staleTime: 30 * 1000,
  });
}
