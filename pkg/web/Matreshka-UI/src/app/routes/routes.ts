import Home from "@/pages/HomePage.vue";
import NotFound from "@/pages/NotFoundPage.vue";
import DisplayConfigPage from "@/pages/DisplayConfigPage.vue";

import {createRouter, createWebHistory} from "vue-router";
import NewConfigPage from "@/pages/NewConfigPage.vue";

export enum Pages {
    Unknown = "unknown",
    Home = "home_page",
    DisplayConfig = "display_config",
    NewConfig = "new",
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
        name: Pages.NewConfig,
        path: '/new',
        component: NewConfigPage,
    },
    {
        name: Pages.DisplayConfig,
        path: "/config/:configName/",
        component: DisplayConfigPage,
        props: true,
    },
]

export const router = createRouter({
    history: createWebHistory(),
    routes,
})

export function RouteToConfigDisplay(serviceName: string) {
    router
        .push({
            name: Pages.DisplayConfig,
            params: {configName: serviceName}
        })
        .then()
}
