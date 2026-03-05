import { render, screen } from "@testing-library/react";
import { MemoryRouter } from "react-router";
import App from "./App";
import "./i18n";

describe("App", () => {
    it("renders home page with navigation", () => {
        render(
            <MemoryRouter>
                <App />
            </MemoryRouter>,
        );
        expect(screen.getByText("Home")).toBeInTheDocument();
        expect(screen.getByText("Page 1")).toBeInTheDocument();
        expect(screen.getByText("Page 2")).toBeInTheDocument();
    });
});
