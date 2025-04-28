import "@/style.css";
import "primeicons/primeicons.css";

import Aura from "@primevue/themes/aura";

import { createPinia } from "pinia";
import PrimeVue from "primevue/config";
import ToastService from "primevue/toastservice";
import Tooltip from "primevue/tooltip";
import { createApp } from "vue";

import App from "@/app/App.vue";
import { router } from "@/app/routes/Routes.ts";

createApp(App)
  .directive("tooltip", Tooltip)
  .use(ToastService)
  .use(createPinia())
  .use(router)
  .use(PrimeVue, {
    theme: {
      preset: Aura,
      options: {
        prefix: "p",
        darkModeSelector: "light",
        cssLayer: false,
      },
    },
  })
  .mount("#app");
