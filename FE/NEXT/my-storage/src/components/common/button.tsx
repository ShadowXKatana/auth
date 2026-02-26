import type { ButtonHTMLAttributes } from 'react'

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: 'primary' | 'secondary'
}

export const Button = ({ variant = 'primary', className = '', ...props }: ButtonProps) => {
  const variantClass =
    variant === 'primary'
      ? 'bg-foreground text-background hover:opacity-90'
      : 'border border-black/15 bg-transparent hover:bg-black/5 dark:border-white/20 dark:hover:bg-white/10'

  return (
    <button
      className={`inline-flex items-center justify-center rounded-lg px-4 py-2 text-sm font-medium transition ${variantClass} ${className}`}
      {...props}
    />
  )
}
