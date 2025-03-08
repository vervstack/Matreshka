import {defineStore} from "pinia";
import {setBackendUrl} from "@/processes/Api/Api.ts";


export const useSettingsStore = defineStore('settings', {
    state: () => {
        const s = {
            apiURL: getBackendUrl()
        } as Settings

        setBackendUrl(s.apiURL)

        return s
    },
    getters: {
        getApiURL: (state) => {
            return state.apiURL
        },
    },
    actions: {
        setApiURL(url: string) {
            this.apiURL = url
            setBackendUrl(this.apiURL)
            localStorage.setItem(backendUrlLocalStorageKey(), url)
        },
    },
})

export type Settings = {
    apiURL: string
}

function backendUrlLocalStorageKey(): string {
    return "backend_api_url"
}

export function getBackendUrl(): string {
    const item = localStorage.getItem(backendUrlLocalStorageKey())
    const beApiAddr =  item || import.meta.env.VITE_MATRESHKA_BACKEND_URL
    localStorage.setItem(backendUrlLocalStorageKey(), beApiAddr)
    return beApiAddr
}
