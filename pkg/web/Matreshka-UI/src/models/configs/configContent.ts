import {Change} from "@/models/configs/verv/info/AppInfo.ts";
import {Component} from "vue";

export interface ConfigContent {
    getChanges(): Change[]
    rollback(): void

    isChanged(): boolean;

    getComponent(): Component
}
