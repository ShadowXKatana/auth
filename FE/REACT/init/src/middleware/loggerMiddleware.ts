import type { MiddlewareFn } from "./types";

/**
 * Example middleware: logs every route navigation to the console.
 */
export const loggerMiddleware: MiddlewareFn = ({ pathname }) => {
    console.log(`[Logger] Navigated to: ${pathname}`);
    return null; // allow route to render
};
