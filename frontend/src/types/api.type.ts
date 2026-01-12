export type ApiResponse<T> = {
  status: boolean;
  message: string;
  data: T;
  error?: unknown;
  meta?: unknown;
};
