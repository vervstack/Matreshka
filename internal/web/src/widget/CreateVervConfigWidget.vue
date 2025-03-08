<script setup lang="ts">

import {ref} from "vue";

import {useToast} from "primevue/usetoast";
import Button from 'primevue/button';
import InputText from "primevue/inputtext";
import FloatLabel from "primevue/floatlabel";
import InputGroup from "primevue/inputgroup";

import {RouteToConfigDisplay} from "@/app/routes/routes.ts";

import {CreateConfig} from "@/processes/Api/Api.ts";
import {handleGrpcError} from "@/processes/Api/ErrorCodes.ts";

const serviceName = ref<string>('');

const toastApi = useToast()

function isNameValid(): boolean {
  return true
}

function inputName() {
  serviceName.value = serviceName.value.replace(" ", "_")
}

function createConfig() {
  CreateConfig(serviceName.value)
      .then(() => {
        toastApi.add({
          closable: true,
          life: 2_000,
          severity: 'success',
          summary: `Service created. Check it out`,
        })

        RouteToConfigDisplay(serviceName.value)
      })
      .catch(handleGrpcError(toastApi))
}

</script>

<template>
  <div class="container">
    <InputGroup>
      <FloatLabel variant="on">
        <InputText
            v-model="serviceName"
            @invalid="isNameValid"
            @input="inputName"
        />
        <label>Service name</label>
      </FloatLabel>
    </InputGroup>
    <div :style="{display: 'flex', justifyContent: 'end'}">
      <Button
          severity="contrast"
          raised
          outlined
          :disabled="serviceName === '' && !isNameValid()"
          icon="pi pi-hammer"
          label="Create"
          :onclick="createConfig"
      />
    </div>
  </div>
</template>

<style scoped>

.container {
  display: flex;
  gap: 1em;

  padding: 1em 0 0 0;

  flex-direction: column;

  width: 100%;
  height: 100%;
}
</style>
