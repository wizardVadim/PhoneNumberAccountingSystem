<script setup lang="ts">
import { ref, watch } from 'vue';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Calendar from 'primevue/calendar';
import { request } from '../../../pkg/Request';

const props = defineProps<{
  visible: boolean
  person: any
  isEdit: boolean
}>()

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void
  (e: 'saved'): void
}>()

const formData = ref({
  id: null,
  city: '',
  address: '',
  first_name: '',
  last_name: '',
  second_name: '',
  born_year: null as Date | null
})

const save = async () => {
  try {
    const bornYearValue = formData.value.born_year ? formData.value.born_year.getFullYear() : null
    
    if (props.isEdit) {
      await request('/persons', {
        method: 'PUT',
        body: {
          ...formData.value,
          born_year: bornYearValue
        }
      })
    } else {
      await request('/persons', {
        method: 'POST',
        body: {
          city: formData.value.city,
          address: formData.value.address || null,
          first_name: formData.value.first_name,
          last_name: formData.value.last_name,
          second_name: formData.value.second_name || null,
          born_year: bornYearValue
        }
      })
    }
    
    emit('saved')
    emit('update:visible', false)
  } catch (error) {
    console.error('Failed to save person:', error)
  }
}

watch(() => props.person, (newPerson) => {
  if (newPerson && Object.keys(newPerson).length > 0) {
    formData.value = {
      id: newPerson.id || null,
      city: newPerson.city || '',
      address: newPerson.address || '',
      first_name: newPerson.first_name || '',
      last_name: newPerson.last_name || '',
      second_name: newPerson.second_name || '',
      born_year: newPerson.born_year ? new Date(newPerson.born_year, 0, 1) : null
    }
  } else {
    formData.value = {
      id: null,
      city: '',
      address: '',
      first_name: '',
      last_name: '',
      second_name: '',
      born_year: null
    }
  }
}, { immediate: true, deep: true })
</script>

<template>
  <Dialog 
    :visible="visible" 
    :header="isEdit ? 'Редактирование физ. лица' : 'Создание физ. лица'"
    :modal="true" 
    :style="{ width: '500px' }"
    @update:visible="$emit('update:visible', $event)"
  >
    <div class="dialog-content">
      <div class="field">
        <label>Фамилия *</label>
        <InputText v-model="formData.last_name" placeholder="Введите фамилию" />
      </div>
      <div class="field">
        <label>Имя *</label>
        <InputText v-model="formData.first_name" placeholder="Введите имя" />
      </div>
      <div class="field">
        <label>Отчество</label>
        <InputText v-model="formData.second_name" placeholder="Введите отчество" />
      </div>
      <div class="field">
        <label>Город *</label>
        <InputText v-model="formData.city" placeholder="Введите город" />
      </div>
      <div class="field">
        <label>Адрес</label>
        <InputText v-model="formData.address" placeholder="Введите адрес" />
      </div>
      <div class="field">
        <label>Год рождения</label>
        <Calendar 
          v-model="formData.born_year" 
          view="year" 
          dateFormat="yy" 
          placeholder="Выберите год"
        />
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
.field .p-calendar {
  width: 100%;
}
</style>