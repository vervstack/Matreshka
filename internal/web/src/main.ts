import {createApp} from 'vue'
import '@/style.css'

import {router} from "@/app/routes/routes.ts";

import {createPinia} from "pinia";

import 'primeicons/primeicons.css';
import ToastService from 'primevue/toastservice';
import PrimeVue from 'primevue/config';
import Tooltip from 'primevue/tooltip';
import Aura from '@primevue/themes/aura';

import App from "@/app/App.vue";

createApp(App)
    .directive('tooltip', Tooltip)
    .use(ToastService)
    .use(createPinia())
    .use(router)
    .use(PrimeVue, {
        theme: {
            preset: Aura,
            options: {
                prefix: 'p',
                darkModeSelector: 'light',
                cssLayer: false
            }
        }
    })
    .mount('#app');
