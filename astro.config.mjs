import { defineConfig } from 'astro/config';

import svelte from "@astrojs/svelte";
import tailwind from "@astrojs/tailwind";

export default defineConfig({
  site: 'https://nayanthespaceguy.github.io',
  integrations: [
    svelte(),
    tailwind({
      nesting: true,
    }),
  ]
});