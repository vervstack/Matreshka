import { defineStore } from "pinia";

import { setBackendAddress, setPass } from "@/processes/api/ApiService.ts";

export const useSettingsStore = defineStore("settings", {
  state: () => {
    const s = {
      apiURL: getBackendUrl(),
      pass: getPass(),
    } as Settings;

    setBackendAddress(s.apiURL);
    setPass(s.pass);

    return s;
  },
  getters: {
    getApiURL: (state) => {
      return state.apiURL;
    },
    getPass: (state) => {
      return state.pass;
    }
  },
  actions: {
    setApiURL(url: string) {
      this.apiURL = url;
      setBackendAddress(this.apiURL);
      localStorage.setItem(backendUrlLocalStorageKey(), url);
    },
    setPass(pass: string) {
      this.pass = pass;
      setPass(this.pass);
      localStorage.setItem(backendPassLocalStorageKey(), pass);
    },
  },
});

export type Settings = {
  apiURL: string;
  pass: string;
};

function backendUrlLocalStorageKey(): string {
  return "backend_api_url";
}

function backendPassLocalStorageKey(): string {
  return "backend_api_pass";
}

export function getBackendUrl(): string {
  const item = localStorage.getItem(backendUrlLocalStorageKey());

  if (!item) {
    console.debug("No backend url in local storage. Extracting from env");
  }

  const beApiAddr = item !== null ? item : import.meta.env.VITE_MATRESHKA_BACKEND_URL;

  localStorage.setItem(backendUrlLocalStorageKey(), beApiAddr);

  console.debug(
    beApiAddr
      ? `Backend url is: ${beApiAddr}`
      : "No backend url is specified. Requests will be routed to root"
  );
  return beApiAddr;
}

export function getPass() : string {
  const item = localStorage.getItem(backendPassLocalStorageKey());

  if (!item) {
    return ""
  }

  return item;
}
