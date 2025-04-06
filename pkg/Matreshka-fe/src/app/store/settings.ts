import {defineStore} from "pinia";
import {setBackendAddress} from "@/processes/api/api.ts";

export const useSettingsStore = defineStore('settings', {
    state: () => {
        const s = {
            apiURL: getBackendUrl()
        } as Settings

        setBackendAddress(s.apiURL)

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
            setBackendAddress(this.apiURL)
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
    const beApiAddr = item !== null ? item : import.meta.env.VITE_MATRESHKA_BACKEND_URL

    localStorage.setItem(backendUrlLocalStorageKey(), beApiAddr)

    return beApiAddr
}
