import * as React from "react";
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationPrevious,
  PaginationNext,
  PaginationEllipsis,
} from "./Pagination";
import { getPaginationRange } from "@/helpers/utils";

interface DefaultPaginationProps {
  page: number;
  totalPages: number;
  onPageChange: React.Dispatch<React.SetStateAction<number>>;
  siblingCount?: number;
}

export const DefaultPagination = ({
  page,
  totalPages,
  onPageChange,
  siblingCount = 1,
}: DefaultPaginationProps) => {
  if (totalPages <= 1) return null;

  const pages = getPaginationRange(page, totalPages, siblingCount);

  return (
    <Pagination>
      <PaginationContent>
        <PaginationItem>
          <PaginationPrevious
            href="#"
            onClick={(e) => {
              e.preventDefault();
              if (page > 1) onPageChange(page - 1);
            }}
            className={page === 1 ? "pointer-events-none opacity-50" : ""}
          />
        </PaginationItem>

        {pages.map((p, idx) => (
          <PaginationItem key={idx}>
            {p === "ellipsis" ? (
              <PaginationEllipsis />
            ) : (
              <PaginationLink
                href="#"
                isActive={p === page}
                onClick={(e) => {
                  e.preventDefault();
                  onPageChange(p);
                }}
              >
                {p}
              </PaginationLink>
            )}
          </PaginationItem>
        ))}

        <PaginationItem>
          <PaginationNext
            href="#"
            onClick={(e) => {
              e.preventDefault();
              if (page < totalPages) onPageChange(page + 1);
            }}
            className={
              page === totalPages ? "pointer-events-none opacity-50" : ""
            }
          />
        </PaginationItem>
      </PaginationContent>
    </Pagination>
  );
};
