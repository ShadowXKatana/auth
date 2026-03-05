type CardProps = {
  children: React.ReactNode
}

export const Card = ({ children }: CardProps) => {
  return (
    <section className="rounded-xl border border-black/10 bg-white p-5 shadow-sm dark:border-white/10 dark:bg-black/20">
      {children}
    </section>
  )
}
