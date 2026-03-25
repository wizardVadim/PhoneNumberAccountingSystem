<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import UserDialog from '@/components/UsersComponents/NewUserModal.vue';
import { request } from '../../pkg/Request';
import { useToast } from 'primevue/usetoast'

const roleId = ref(-1)
const users = ref([])
const showDialog = ref(false)
const isEdit = ref(false)
const selectedUser = ref({})

const toast = useToast()

const showSuccess = (message: string) => {
  toast.add({
    severity: 'success',
    summary: 'Успех',
    detail: message,
    life: 3000
  })
}

const showError = (message: string) => {
  toast.add({
    severity: 'error',
    summary: 'Ошибка',
    detail: message,
    life: 3000
  })
}


const fetchUsers = async () => {
  try {
    const data = await request('/users', { method: 'GET' })
    users.value = data
  } catch (error) {
    console.error('Failed to fetch users:', error)
  }
}

const openCreateDialog = () => {
  isEdit.value = false
  selectedUser.value = {}
  showDialog.value = true
}

const openEditDialog = (user: any) => {
  isEdit.value = true
  selectedUser.value = user
  showDialog.value = true
}

const deleteUser = async (user: any) => {
  if (confirm(`Удалить пользователя ${user.login}?`)) {
    try {
      await request('/users', { 
        method: 'DELETE', 
        body: { id: user.id } 
      })
      showSuccess("Успешное удаление")
      await fetchUsers()
    } catch (error) {
      showError(error)
    }
  }
}

onBeforeMount(async () => {
  roleId.value = Number(localStorage.getItem("roleId"))
  if (roleId.value === 1) {
    await fetchUsers()
  }
})
</script>

<template>
  <div class="users-container">
    <div class="header">
      <h1>Список пользователей</h1>
      <div>
        <Button @click="fetchUsers" v-if="roleId == 1" severity="secondary" label="Обновить" />
        <Button @click="openCreateDialog" v-if="roleId == 1" severity="primary" label="Новый" />
      </div>
    </div>

    <DataTable :value="users">
      <Column field="id" header="ID"></Column>
      <Column field="login" header="Логин"></Column>
      <Column field="role_id" header="ID роли"></Column>
      <Column field="role_name" header="Роль"></Column>
      <Column header="Действия">
        <template #body="{ data }">
          <Button 
            icon="pi pi-pencil" 
            severity="secondary" 
            rounded 
            text 
            @click="openEditDialog(data)"
            style="margin-right: 8px"
          />
          <Button 
            icon="pi pi-trash" 
            severity="danger" 
            rounded 
            text 
            @click="deleteUser(data)"
          />
        </template>
      </Column>
    </DataTable>

    <UserDialog 
      v-model:visible="showDialog"
      :user="selectedUser"
      :is-edit="isEdit"
      @saved="fetchUsers"
    />
  </div>
</template>

<style scoped>
.users-container {
  width: 80%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.header button {
  margin-left: 10px;
}
</style>