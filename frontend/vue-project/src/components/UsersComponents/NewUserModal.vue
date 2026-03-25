<!-- src/components/UserDialog.vue -->
<script setup lang="ts">
import { ref, watch } from 'vue';
import Dialog from 'primevue/dialog';
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Button from 'primevue/button';
import { request } from '../../../pkg/Request';

const props = defineProps<{
  visible: boolean
  user: any
  isEdit: boolean
}>()

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void
  (e: 'saved'): void
}>()

const roles = ref([])
const formData = ref({
  id: null,
  login: '',
  password: '',
  role_id: null
})

const fetchRoles = async () => {
  try {
    const data = await request('/roles', { method: 'GET' })
    roles.value = data
  } catch (error) {
    console.error('Failed to fetch roles:', error)
  }
}

const save = async () => {
  try {
    if (props.isEdit) {
      await request('/users', {
        method: 'PUT',
        body: formData.value
      })
    } else {
      await request('/users', {
        method: 'POST',
        body: {
          login: formData.value.login,
          password: formData.value.password,
          role_id: formData.value.role_id
        }
      })
    }
    
    emit('saved')
    emit('update:visible', false)
  } catch (error) {
    console.error('Failed to save user:', error)
  }
}

watch(() => props.user, (newUser) => {
  if (newUser) {
    formData.value = {
      id: newUser.id || null,
      login: newUser.login || '',
      password: '',
      role_id: newUser.role_id || null
    }
  }
}, { immediate: true })

fetchRoles()
</script>

<template>
  <Dialog 
    :visible="visible" 
    :header="isEdit ? 'Редактирование пользователя' : 'Создание пользователя'"
    :modal="true" 
    :style="{ width: '450px' }"
    @update:visible="$emit('update:visible', $event)"
  >
    <div class="dialog-content">
      <div class="field">
        <label>Логин</label>
        <InputText v-model="formData.login" placeholder="Введите логин" />
      </div>
      <div class="field">
        <label>Пароль</label>
        <InputText v-model="formData.password" type="password" placeholder="Введите пароль" />
      </div>
      <div class="field">
        <label>Роль</label>
        <Dropdown 
          v-model="formData.role_id" 
          :options="roles" 
          optionLabel="role_name" 
          optionValue="id"
          placeholder="Выберите роль" 
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
.field .p-dropdown {
  width: 100%;
}
</style>