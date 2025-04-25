import {Change} from "@/models/configs/verv/info/VervConfig.ts";
import {Component} from "vue";

export interface Config_content {
    getChanges(): Change[]
    rollback(): void

    isChanged(): boolean;

    getComponent(): Component
}
