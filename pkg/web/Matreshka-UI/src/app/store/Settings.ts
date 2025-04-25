import {defineStore} from "pinia";
import {setBackendAddress} from "@/processes/api/ApiService.ts";

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

    if (!item) {
        console.debug('No backend url in local storage. Extracting from env')
    }

    const beApiAddr = item !== null ? item : import.meta.env.VITE_MATRESHKA_BACKEND_URL

    localStorage.setItem(backendUrlLocalStorageKey(), beApiAddr)

    console.debug(beApiAddr ?
        `Backend url is: ${beApiAddr}` :
        `No backend url is specified. Requests will be routed to root`)
    return beApiAddr
}
