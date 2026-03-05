import type { InputHTMLAttributes } from 'react'

type TextInputProps = InputHTMLAttributes<HTMLInputElement> & {
  label: string
}

export const TextInput = ({ label, id, ...props }: TextInputProps) => {
  const inputId = id ?? label.toLowerCase().replace(/\s+/g, '-')

  return (
    <label htmlFor={inputId} className="block">
      <span className="mb-1.5 block text-sm font-medium">{label}</span>
      <input
        id={inputId}
        className="w-full rounded-lg border border-black/15 bg-transparent px-3 py-2 text-sm outline-none ring-0 placeholder:text-foreground/40 focus:border-black/30 dark:border-white/15 dark:focus:border-white/30"
        {...props}
      />
    </label>
  )
}
