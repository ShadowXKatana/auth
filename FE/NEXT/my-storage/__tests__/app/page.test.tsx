import Home from '@/app/page'
import { render, screen } from '@testing-library/react'

describe('Home page', () => {
  it('renders starter heading', () => {
    render(<Home />)
    expect(screen.getByText('To get started, edit the page.tsx file.')).toBeInTheDocument()
  })
})
