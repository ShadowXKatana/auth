import { Navigate } from "react-router";
import type { MiddlewareFn } from "./types";

/**
 * Example middleware: checks if the user is authenticated.
 * Replace `isAuthenticated()` with your actual auth check.
 *
 * If not authenticated, redirects to "/" (or a login page).
 */

// TODO: Replace with your actual auth check
function isAuthenticated(): boolean {
    return true;
}

export const authMiddleware: MiddlewareFn = ({ pathname }) => {
    if (!isAuthenticated()) {
        console.warn(`[Auth] Unauthorized access to ${pathname}, redirecting…`);
        return <Navigate to="/" replace />;
    }
    return null; // allow route to render
};
