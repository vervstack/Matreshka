import Home from "@/pages/HomePage.vue";
import NotFound from "@/pages/NotFoundPage.vue";
import DisplayConfigPage from "@/pages/DisplayConfigPage.vue";

import {createRouter, createWebHistory} from "vue-router";
import NewConfigPage from "@/pages/NewConfigPage.vue";
import IAMConfig from "@/components/config/minio/IAMConfig.vue";

export enum Pages {
    Unknown = "unknown",
    Home = "home_page",
    DisplayConfig = "display_config",
    NewAppConfig = "new_app_config",
    NewConfig = "new_config",
    // TODO
    NewMinioConfig = "new_minio_config",
}

export const routes = [
    {
        name: Pages.Unknown,
        path: "/:pathMatch(.*)*",
        component: NotFound,
    },
    {
        name: Pages.Home,
        path: '/',
        component: Home,
    },

    {
        name: Pages.NewAppConfig,
        path: '/new/app_config',
        component: NewConfigPage,
    },
    {
        name: Pages.DisplayConfig,
        path: "/config/:name/",
        component: DisplayConfigPage,
        props: true,
    },
    {
        name: Pages.NewMinioConfig,
        path: "/new/minio/",
        component: IAMConfig,
        props: true,
    }
]

export const router = createRouter({
    history: createWebHistory(),
    routes,
})


export function RouteToConfigDisplay(serviceName: string) {
    router
        .push({
            name: Pages.DisplayConfig,
            params: {name: serviceName}
        })
        .then()
}
