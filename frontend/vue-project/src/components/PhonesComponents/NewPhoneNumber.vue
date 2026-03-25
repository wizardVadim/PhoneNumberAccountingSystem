<script setup lang="ts">
import { ref, watch, onMounted } from 'vue';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Button from 'primevue/button';
import { request } from '../../../pkg/Request';

const props = defineProps<{
  visible: boolean
  phone: any
  isEdit: boolean
  persons?: any[]
}>()

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void
  (e: 'saved'): void
}>()

const phoneTypes = ref([])

const formData = ref({
  id: null,
  phone_number_value: '',
  person_id: null,
  phone_number_type_id: null,
  comment: ''
})

const fetchPhoneTypes = async () => {
  try {
    const data = await request('/phone-types', { method: 'GET' })
    phoneTypes.value = data
  } catch (error) {
    console.error('Failed to fetch phone types:', error)
  }
}

const save = async () => {
  try {
    if (props.isEdit) {
      await request('/phones', {
        method: 'PUT',
        body: formData.value
      })
    } else {
      await request('/phones', {
        method: 'POST',
        body: {
          phone_number_value: formData.value.phone_number_value,
          person_id: formData.value.person_id,
          phone_number_type_id: formData.value.phone_number_type_id,
          comment: formData.value.comment || null
        }
      })
    }
    
    emit('saved')
    emit('update:visible', false)
  } catch (error) {
    console.error('Failed to save phone:', error)
  }
}

watch(() => props.phone, (newPhone) => {
  if (newPhone) {
    formData.value = {
      id: newPhone.id || null,
      phone_number_value: newPhone.phone_number_value || '',
      person_id: newPhone.person_id || null,
      phone_number_type_id: newPhone.phone_number_type_id || null,
      comment: newPhone.comment || ''
    }
  }
}, { immediate: true })

onMounted(() => {
  fetchPhoneTypes()
})
</script>

<template>
  <Dialog 
    :visible="visible" 
    :header="isEdit ? 'Редактирование телефона' : 'Добавление телефона'"
    :modal="true" 
    :style="{ width: '500px' }"
    @update:visible="$emit('update:visible', $event)"
  >
    <div class="dialog-content">
      <div class="field">
        <label>Номер телефона *</label>
        <InputText v-model="formData.phone_number_value" placeholder="+7XXXXXXXXXX" />
      </div>
      <div class="field">
        <label>Владелец *</label>
        <Dropdown 
          v-model="formData.person_id" 
          :options="persons" 
          optionLabel="last_name" 
          optionValue="id"
          placeholder="Выберите владельца"
        >
          <template #option="slotProps">
            <span>{{ slotProps.option.last_name }} {{ slotProps.option.first_name }}</span>
          </template>
        </Dropdown>
      </div>
      <div class="field">
        <label>Тип телефона *</label>
        <Dropdown 
          v-model="formData.phone_number_type_id" 
          :options="phoneTypes" 
          optionLabel="type_name" 
          optionValue="id"
          placeholder="Выберите тип"
        />
      </div>
      <div class="field">
        <label>Комментарий</label>
        <InputText v-model="formData.comment" placeholder="Комментарий" />
      </div>
    </div>
    <template #footer>
      <Button label="Отмена" severity="secondary" @click="$emit('update:visible', false)" />
      <Button label="Сохранить" severity="primary" @click="save" />
    </template>
  </Dialog>
</template>

<style scoped>
.dialog-content {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.field label {
  font-weight: 500;
}

.field input,
.field .p-dropdown {
  width: 100%;
}
</style>