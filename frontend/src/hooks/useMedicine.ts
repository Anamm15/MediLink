import { searchMedicine } from "@/services/medicine.service";
import { useQuery } from "@tanstack/react-query";

export function useSearchMedicineQuery(
  name: string,
  page: number,
  limit: number
) {
  return useQuery({
    queryKey: ["medicine", name],
    queryFn: () => searchMedicine(name, page, limit),
    enabled: !!name,
    staleTime: 30 * 1000,
  });
}
