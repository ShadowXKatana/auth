import { Route, Routes } from "react-router";
import Layout from "../components/Layout";
import { WithMiddleware, loggerMiddleware, authMiddleware } from "../middleware";
import Home from "../pages/Home";
import Page1 from "../pages/Page1";
import Page2 from "../pages/Page2";

export default function AppRoutes() {
    return (
        <Routes>
            <Route
                element={
                    <WithMiddleware middlewares={[loggerMiddleware, authMiddleware]}>
                        <Layout />
                    </WithMiddleware>
                }
            >
                <Route index element={<Home />} />
                <Route path="page-1" element={<Page1 />} />
                <Route path="page-2" element={<Page2 />} />
            </Route>
        </Routes>
    );
}