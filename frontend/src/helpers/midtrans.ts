export const getSnapTokenFromUrl = (url: string): string => {
  const parsedUrl = new URL(url);
  const parts = parsedUrl.pathname.split("/");
  return parts[parts.length - 1];
};
