import { defineConfig } from 'astro/config';
import vercel from '@astrojs/vercel';

export default defineConfig({
    output: 'server',

    adapter: vercel({
        isr: {
            expiration: 3600,
            bypassToken: process.env.VERCEL_ISR_BYPASS_TOKEN,
            exclude: ["/search"],
        },
    }),

    compressHTML: true,
});
