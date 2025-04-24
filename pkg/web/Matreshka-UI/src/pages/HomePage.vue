<script setup lang="ts">
import {onMounted, ref, watch} from "vue";

import {ServiceListClass} from "@/models/configs/verv/info/AppInfo.ts";

import {ListServices} from "@/processes/api/api.ts";
import {handleGrpcError} from "@/processes/api/error_codes.ts";

import Dialog from "primevue/dialog";
import DisplayConfigWidget from "@/widget/DisplayConfigWidget.vue";

import {Pages, router} from "@/app/routes/routes.ts";

import {useToast} from "primevue/usetoast";
import ProgressSpinner from 'primevue/progressspinner';
import Paginator from 'primevue/paginator';

import ServicesListWidget from "@/widget/config_list/ConfigListWidget.vue";
import TopControls from "@/widget/config_list/TopControls.vue";

import {ListServicesReq, Paging, Sort} from "@/models/search/search.ts";
import {SortType} from "@vervstack/matreshka";

const toastApi = useToast();

const isDialogOpen = ref<boolean>(false);
const openedConfigName = ref<string>('');
const isLoading = ref<boolean>(true);

// Service list
const listRequest = ref<ListServicesReq>({
  searchPattern: '',
  sort: {
    type: SortType.default,
    desc: false,
  } as Sort,
  paging: {
    limit: 6,
    offset: 0,
  } as Paging
} as ListServicesReq)

const servicesList = ref<ServiceListClass>()

const pagingTotalRecords = ref<number>(0)

function updateList() {
  if (servicesList.value?.servicesInfo.length == 0) {
    isLoading.value = true;
  }

  ListServices(listRequest.value)
      .then((resp) => {
        servicesList.value = resp
        pagingTotalRecords.value = resp.total
      })
      .catch(handleGrpcError(toastApi))
      .then(() => isLoading.value = false)
}

function openPage(page: number) {
  listRequest.value.paging.offset = (listRequest.value.paging.limit || 10) * page;
  updateList()
}

function updateSearchReq(pattern: string, sort: Sort) {
  listRequest.value.searchPattern = pattern
  listRequest.value.sort = sort
  updateList()
}

onMounted(updateList)

watch(isDialogOpen, () => {
  if (!isDialogOpen.value) {
    updateList()
  }
})

//  Service info
function openDisplayConfigDialog(serviceName: string) {
  openedConfigName.value = serviceName
  isDialogOpen.value = true
}

function openServiceInfo(event: MouseEvent, serviceName: string) {
  if (!(event.ctrlKey || event.metaKey)) {
    openDisplayConfigDialog(serviceName ?? '')
    return
  }

  const routeTo = {
    name: Pages.DisplayConfig,
    params: {name: serviceName},
  }

  window.open(
      router.resolve(routeTo).href, '_blank')
}


</script>

<template>
  <!--  List of services -->
  <div class="Home">
    <div class="ListWrapper">

      <TopControls
          @updateSearchRequest="updateSearchReq"
      />

      <Transition name="load-fader" mode="out-in">
        <div v-if="!isLoading">
          <ServicesListWidget
              v-if="servicesList && servicesList.servicesInfo.length > 0"
              :servicesList="servicesList.servicesInfo"
              @click-service="openServiceInfo"
          />
          <p v-else class="EmptyNodeMessage">No configs on this node</p>
        </div>
        <ProgressSpinner v-else/>
      </Transition>

      <Paginator
          :rows="listRequest.paging.limit"
          :totalRecords="pagingTotalRecords"
          @page="event => openPage(event.page)"
      />
    </div>
  </div>

  <Dialog
      v-model:visible="isDialogOpen"
      modal
      :closable="false"
      :draggable="false"
      :dismissableMask="true"
      :pt="{
          root: 'border-none',
          mask: {
            style: 'backdrop-filter: blur(2px)'
          }
        }"
      :style="{
          width: '80vw',
          height: '95vh',
        }"
      position="center"
  >
    <DisplayConfigWidget
        :config-name="openedConfigName"
    />
  </Dialog>
</template>

<style scoped>
.Home {
  padding: 2em;

  display: flex;
  flex-direction: column;
  align-items: center;

  gap: 1em;
  height: 100%;
}

.ListWrapper {
  display: flex;
  flex-direction: column;
  gap: 0.5em;
  width: 100%;
  align-items: center;
}

.ListWrapper > * {
  width: 100%;
}

.EmptyNodeMessage {
  display: flex;
  justify-content: center;
}

.load-fader-enter-active,
.load-fader-leave-active {
  transition: 0.25s;
}

.load-fader-enter-to,
.load-fader-leave-from {
  opacity: 1;
}

.load-fader-enter-from,
.load-fader-leave-to {
  opacity: 0;
}
</style>
