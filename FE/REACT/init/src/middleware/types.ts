import type { ReactNode } from "react";

/**
 * A middleware function receives route context and returns either:
 * - `null` to allow the route to render normally
 * - A `ReactNode` to replace the route content (e.g. redirect, error page)
 */
export type MiddlewareFn = (context: MiddlewareContext) => ReactNode | null;

export interface MiddlewareContext {
    /** Current pathname */
    pathname: string;
}
