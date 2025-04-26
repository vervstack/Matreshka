<script setup lang="ts">
import Button from "primevue/button";
import { ref } from "vue";

import ConfigField from "@/components/base/config/fields/ConfigInput.vue";
import IamStatement from "@/components/config/minio/IamStatement.vue";
import {
  exportMinioStatement,
  IamConfig,
  IamMinioExportConfig,
  S3Action,
  Statement,
} from "@/models/resource_configs/s3/minio/minio.ts";
import { ConfigValue } from "@/models/shared/common.ts";

// const model = defineModel<IamConfig>({
//   required: true
// })

const model = ref<IamConfig>({
  version: {
    envName: "Version",
    value: "2012-10-17",
  },
  statements: [],
} as IamConfig);

function addStatement() {
  model.value.statements.push({
    allow: {
      envName: "Effect",
      value: true,
    } as ConfigValue<boolean>,
    action: {
      envName: "Action(s)",
      value: [],
    } as ConfigValue<S3Action[]>,
    resources: {
      envName: "Resources",
      value: [""],
    } as ConfigValue<string[]>,
  } as Statement);
}

function exportToJson() {
  const exportObject: IamMinioExportConfig = {} as IamMinioExportConfig;
  exportObject.Version = model.value.version.value;

  exportObject.Statement = model.value.statements.map(exportMinioStatement);

  const blob = new Blob([JSON.stringify(exportObject, null, 2)], {
    type: "application/json",
  });

  // Step 3: Create a download link
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");
  link.href = url;
  link.download = "data.json"; // Filename for the downloaded file

  // Step 4: Append link to body and trigger download
  document.body.appendChild(link);
  link.click();

  // Step 5: Cleanup
  document.body.removeChild(link);
  URL.revokeObjectURL(url);
}
</script>

<template>
  <div class="Node IAMConfig">
    <div class="TopControls Controls">
      <Button
        severity="secondary"
        icon="pi pi-file-export"
        v-tooltip.bottom="'Export to JSON file'"
        @click="exportToJson"
      />
    </div>
    <div class="NodeField">
      <ConfigField v-model="model.version" />
    </div>
    <div class="Node">
      <div class="Controls">
        <p>Statements</p>
        <Button rounded icon="pi pi-plus" @click="addStatement" v-tooltip.right="'Add statement'" />
      </div>
      <div class="Node">
        <IamStatement v-for="(_, idx) in model.statements" v-model="model.statements[idx]" />
      </div>
    </div>
  </div>
</template>

<style scoped>
@import "@/assets/styles/config_display.css";

.TopControls {
  justify-content: end;
}

.Controls {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 2em;
}

.IAMConfig {
  padding: 5vh 0 5vh 0;
}
</style>
