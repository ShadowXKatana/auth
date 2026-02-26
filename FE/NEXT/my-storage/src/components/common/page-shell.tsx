type PageShellProps = {
  title: string
  subtitle?: string
  children: React.ReactNode
}

export const PageShell = ({ title, subtitle, children }: PageShellProps) => {
  return (
    <main className="mx-auto w-full max-w-3xl p-6 md:p-10">
      <header className="mb-6">
        <h1 className="text-2xl font-semibold tracking-tight">{title}</h1>
        {subtitle ? <p className="mt-1 text-sm text-foreground/80">{subtitle}</p> : null}
      </header>
      {children}
    </main>
  )
}
