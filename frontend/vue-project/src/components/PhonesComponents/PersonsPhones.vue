<script setup lang="ts">
import { ref, watch } from 'vue';
import Dialog from 'primevue/dialog';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import { request } from '../../../pkg/Request';
import { useToast } from 'primevue/usetoast';

const props = defineProps<{
  visible: boolean
  personId: number
  personName: string
}>()

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void
}>()

const phones = ref([])
const toast = useToast()

const fetchPersonPhones = async () => {
  if (!props.personId) return
  
  try {
    const data = await request(`/persons/${props.personId}/phones`, { method: 'GET' })
    phones.value = data
  } catch (error) {
    console.error('Failed to fetch person phones:', error)
    toast.add({
      severity: 'error',
      summary: 'Ошибка',
      detail: 'Не удалось загрузить телефоны',
      life: 3000
    })
  }
}

watch(() => props.visible, (newVal) => {
  if (newVal && props.personId) {
    fetchPersonPhones()
  }
})
</script>

<template>
  <Dialog 
    :visible="visible" 
    :header="`Телефоны: ${personName}`"
    :modal="true" 
    :style="{ width: '700px' }"
    @update:visible="$emit('update:visible', $event)"
  >
    <DataTable :value="phones" class="p-datatable-sm">
      <Column field="id" header="ID"></Column>
      <Column field="phone_number_value" header="Номер"></Column>
      <Column field="phone_number_type" header="Тип"></Column>
      <Column field="comment" header="Комментарий"></Column>
    </DataTable>
    
    <template #footer>
      <Button label="Закрыть" severity="secondary" @click="$emit('update:visible', false)" />
    </template>
  </Dialog>
</template>