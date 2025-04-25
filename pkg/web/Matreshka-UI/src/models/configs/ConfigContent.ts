import {Change} from "@/models/configs/Change.ts";
import {Component} from "vue";

export default interface ConfigContent {
    getChanges(): Change[]
    rollback(): void

    isChanged(): boolean;

    getComponent(): Component
}
