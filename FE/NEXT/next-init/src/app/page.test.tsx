import { render, screen } from "@testing-library/react";
import Home from "./page";

describe("Home page", () => {
  it("renders starter heading", () => {
    render(<Home />);
    expect(screen.getByText("To get started, edit the page.tsx file.")).toBeInTheDocument();
  });
});
