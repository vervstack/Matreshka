import {Component} from "vue";

import ConfigContent from "@/models/configs/ConfigContent.ts";
import {Change} from "@/models/configs/Change.ts";

import KeyValueConfigView from "@/components/config/keyvalue/KeyValueConfigView.vue";

// TODO Implement
export default class KeyValueConfigContent implements ConfigContent{
    getChanges(): Change[] {
        return []
    }
    rollback(): void {

    }
    isChanged(): boolean {
        return false;
    }

    getComponent() : Component {
        return KeyValueConfigView
    }
}
