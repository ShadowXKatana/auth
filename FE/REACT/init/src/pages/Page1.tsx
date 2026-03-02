import { useTranslation } from "react-i18next";

export default function Page1() {
    const { t } = useTranslation();

    return (
        <div className="flex flex-col items-center gap-6 py-16 px-8 text-center sm:items-start sm:text-left">
            <h1 className="text-3xl font-semibold tracking-tight text-black dark:text-zinc-50">
                {t("page1.title")}
            </h1>
            <p className="max-w-md text-lg leading-8 text-zinc-600 dark:text-zinc-400">
                {t("page1.description")}
            </p>
            <div className="rounded-lg border border-zinc-200 bg-zinc-50 p-6 dark:border-zinc-800 dark:bg-zinc-900">
                <code className="text-sm text-zinc-700 dark:text-zinc-300">src/pages/Page1.tsx</code>
            </div>
        </div>
    );
}
