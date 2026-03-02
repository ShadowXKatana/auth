import { z } from "zod";

const envSchema = z.object({
    VITE_APP_NAME: z.string().min(1).default("my-app"),
});

export const env = envSchema.parse({
    VITE_APP_NAME: import.meta.env.VITE_APP_NAME,
});
