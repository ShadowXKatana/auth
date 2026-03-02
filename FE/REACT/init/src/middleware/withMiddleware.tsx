import { useLocation } from "react-router";
import type { ReactNode } from "react";
import type { MiddlewareFn } from "./types";

interface WithMiddlewareProps {
    children: ReactNode;
    middlewares: MiddlewareFn[];
}

/**
 * Runs a chain of middleware functions before rendering the child route.
 * If any middleware returns a ReactNode, that node is rendered instead.
 *
 * Usage:
 * ```tsx
 * <WithMiddleware middlewares={[loggerMiddleware, authMiddleware]}>
 *   <Outlet />
 * </WithMiddleware>
 * ```
 */
export function WithMiddleware({ children, middlewares }: WithMiddlewareProps) {
    const location = useLocation();

    for (const mw of middlewares) {
        const result = mw({ pathname: location.pathname });
        if (result !== null) {
            return <>{result}</>;
        }
    }

    return <>{children}</>;
}
