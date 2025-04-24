import {ConfigContent} from "@/models/configs/configContent.ts";
import {Change} from "@/models/configs/verv/info/AppInfo.ts";
import {Component} from "vue";
import KeyValueConfigView from "@/components/config/keyvalue/KeyValueConfigView.vue";

// TODO Implement
export class KeyValueConfigContent implements ConfigContent{
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
