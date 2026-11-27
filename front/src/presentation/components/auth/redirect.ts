export const LOGIN_URL = "/auth/login";
export const HOME_URL = "/";

/** Login URL that brings the user back to `pathname` once logged in. */
export const loginUrlFor = (pathname: string | null): string =>
  !pathname || pathname === HOME_URL
    ? LOGIN_URL
    : `${LOGIN_URL}?next=${encodeURIComponent(pathname)}`;

/** Only same-site paths: "/x" is allowed, "//evil.com" or "https://..." are not (open redirect). */
export const safeNextPath = (next: string | null | undefined): string =>
  next && next.startsWith("/") && !next.startsWith("//") && !next.startsWith("/\\")
    ? next
    : HOME_URL;

export const nextPathFromLocation = (): string =>
  typeof window === "undefined"
    ? HOME_URL
    : safeNextPath(new URLSearchParams(window.location.search).get("next"));
