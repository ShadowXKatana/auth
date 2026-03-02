import { useTranslation } from "react-i18next";

export default function Home() {
    const { t } = useTranslation();

    return (
        <div className="flex min-h-[calc(100vh-57px)] flex-col items-center justify-between py-32 sm:items-start">
            <img src="/vite.svg" alt="Vite logo" className="h-8 dark:invert" />
            <div className="flex flex-col items-center gap-6 text-center sm:items-start sm:text-left">
                <h1 className="max-w-xs text-3xl font-semibold leading-10 tracking-tight text-black dark:text-zinc-50">
                    {t("home.title")}
                </h1>
                <p className="max-w-md text-lg leading-8 text-zinc-600 dark:text-zinc-400">
                    {t("home.description")}
                </p>
            </div>
            <div className="flex flex-col gap-4 text-base font-medium sm:flex-row">
                <a
                    className="flex h-12 w-full items-center justify-center gap-2 rounded-full bg-foreground px-5 text-background transition-colors hover:bg-[#383838] dark:hover:bg-[#ccc] md:w-[158px]"
                    href="https://vite.dev"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    <img src="/vite.svg" alt="Vite logomark" className="h-4 w-4 dark:invert" />
                    Vite Docs
                </a>
                <a
                    className="flex h-12 w-full items-center justify-center rounded-full border border-solid border-black/[.08] px-5 transition-colors hover:border-transparent hover:bg-black/[.04] dark:border-white/[.145] dark:hover:bg-[#1a1a1a] md:w-[158px]"
                    href="https://react.dev"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    React Docs
                </a>
            </div>
        </div>
    );
}
