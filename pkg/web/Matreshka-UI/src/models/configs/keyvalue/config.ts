import {Config_content} from "@/models/configs/config_content.ts";
import {Change} from "@/models/configs/verv/info/VervConfig.ts";
import {Component} from "vue";
import KeyValueConfigView from "@/components/config/keyvalue/KeyValueConfigView.vue";

// TODO Implement
export class KeyValueConfigContent implements Config_content{
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
