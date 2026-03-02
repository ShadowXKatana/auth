import { NavLink, Outlet } from "react-router";
import { useTranslation } from "react-i18next";

export default function Layout() {
    const { t, i18n } = useTranslation();

    const toggleLanguage = () => {
        const next = i18n.language === "en" ? "th" : "en";
        i18n.changeLanguage(next);
    };

    return (
        <div className="flex min-h-screen flex-col bg-zinc-50 font-sans dark:bg-black">
            {/* ── Navigation ── */}
            <nav className="sticky top-0 z-10 flex items-center justify-between border-b border-zinc-200 bg-white/80 px-6 py-3 backdrop-blur dark:border-zinc-800 dark:bg-black/80">
                <div className="flex items-center gap-1">
                    <NavLink
                        to="/"
                        end
                        className={({ isActive }) =>
                            `rounded-md px-3 py-2 text-sm font-medium transition-colors ${isActive
                                ? "bg-zinc-900 text-white dark:bg-zinc-100 dark:text-black"
                                : "text-zinc-600 hover:bg-zinc-100 dark:text-zinc-400 dark:hover:bg-zinc-800"
                            }`
                        }
                    >
                        Home
                    </NavLink>
                    <NavLink
                        to="/page-1"
                        className={({ isActive }) =>
                            `rounded-md px-3 py-2 text-sm font-medium transition-colors ${isActive
                                ? "bg-zinc-900 text-white dark:bg-zinc-100 dark:text-black"
                                : "text-zinc-600 hover:bg-zinc-100 dark:text-zinc-400 dark:hover:bg-zinc-800"
                            }`
                        }
                    >
                        {t("nav.page1")}
                    </NavLink>
                    <NavLink
                        to="/page-2"
                        className={({ isActive }) =>
                            `rounded-md px-3 py-2 text-sm font-medium transition-colors ${isActive
                                ? "bg-zinc-900 text-white dark:bg-zinc-100 dark:text-black"
                                : "text-zinc-600 hover:bg-zinc-100 dark:text-zinc-400 dark:hover:bg-zinc-800"
                            }`
                        }
                    >
                        {t("nav.page2")}
                    </NavLink>
                </div>

                {/* Language switcher */}
                <button
                    onClick={toggleLanguage}
                    className="rounded-md border border-zinc-200 px-3 py-1.5 text-xs font-medium text-zinc-600 transition-colors hover:bg-zinc-100 dark:border-zinc-700 dark:text-zinc-400 dark:hover:bg-zinc-800"
                >
                    {i18n.language === "en" ? "🇹🇭 TH" : "🇺🇸 EN"}
                </button>
            </nav>

            {/* ── Page content ── */}
            <main className="mx-auto w-full max-w-3xl flex-1 px-6">
                <Outlet />
            </main>
        </div>
    );
}
