import { defineConfig } from 'astro/config';

import tailwind from "@astrojs/tailwind";

// https://astro.build/config
export default defineConfig({
  site: 'https://nayanthespaceguy.github.io',
  integrations: [tailwind()]
});