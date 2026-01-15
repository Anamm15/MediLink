export function getPaginationRange(
  currentPage: number,
  totalPages: number,
  siblingCount: number
): (number | "ellipsis")[] {
  const range: (number | "ellipsis")[] = [];
  const left = Math.max(2, currentPage - siblingCount);
  const right = Math.min(totalPages - 1, currentPage + siblingCount);

  range.push(1);

  if (left > 2) range.push("ellipsis");

  for (let i = left; i <= right; i++) {
    range.push(i);
  }

  if (right < totalPages - 1) range.push("ellipsis");

  if (totalPages > 1) range.push(totalPages);

  return range;
}
