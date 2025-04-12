<script setup lang="ts">

import {CreateConfig} from "@/processes/api/api.ts";
import {RouteToConfigDisplay} from "@/app/routes/routes.ts";
import {handleGrpcError} from "@/processes/api/error_codes.ts";
import {useToast} from "primevue/usetoast";
import Button from "primevue/button";

const toastApi = useToast()

const props = defineProps({
  serviceName: {
    type: String,
    required: true
  },
})

function createConfig() {
  CreateConfig(props.serviceName)
      .then(() => {
        toastApi.add({
          closable: true,
          life: 2_000,
          severity: 'success',
          summary: `Service created. Check it out`,
        })

        RouteToConfigDisplay(props.serviceName)
      })
      .catch(handleGrpcError(toastApi))
}

function isNameValid(): boolean {
  return true
}

</script>

<template>
  <Button
      severity="contrast"
      raised
      outlined
      :disabled="props.serviceName === '' && !isNameValid()"
      icon="pi pi-hammer"
      label="Create"
      :onclick="createConfig"
  />
</template>

<style scoped>

</style>
