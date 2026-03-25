<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import PhoneDialog from '@/components/PhonesComponents/NewPhoneNumber.vue';
import PersonPhonesDialog from '@/components/PhonesComponents/PersonsPhones.vue';
import { request } from '../../pkg/Request';
import { useToast } from 'primevue/usetoast'

const roleId = ref(-1)
const phones = ref([])
const persons = ref([])
const showDialog = ref(false)
const showPersonPhonesDialog = ref(false)
const isEdit = ref(false)
const selectedPhone = ref({})
const selectedPersonId = ref(0)
const selectedPersonName = ref('')

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

const fetchPhones = async () => {
  try {
    const data = await request('/phones', { method: 'GET' })
    phones.value = data
  } catch (error) {
    console.error('Failed to fetch phones:', error)
  }
}

const fetchPersons = async () => {
  try {
    const data = await request('/persons', { method: 'GET' })
    persons.value = data
  } catch (error) {
    console.error('Failed to fetch persons:', error)
  }
}

const openCreateDialog = () => {
  isEdit.value = false
  selectedPhone.value = {}
  showDialog.value = true
}

const openEditDialog = (phone: any) => {
  isEdit.value = true
  selectedPhone.value = phone
  showDialog.value = true
}

const deletePhone = async (phone: any) => {
  if (confirm(`Удалить телефон ${phone.phone_number_value}?`)) {
    try {
      await request('/phones', { 
        method: 'DELETE', 
        body: { id: phone.id } 
      })
      showSuccess("Успешное удаление")
      await fetchPhones()
    } catch (error) {
      showError('Ошибка удаления')
    }
  }
}

const openPersonPhones = (personId: number, personName: string) => {
  selectedPersonId.value = personId
  selectedPersonName.value = personName
  showPersonPhonesDialog.value = true
}

onBeforeMount(async () => {
  roleId.value = Number(localStorage.getItem("roleId"))
  if (roleId.value === 1 || roleId.value === 2) {
    await fetchPhones()
    await fetchPersons()
  }
})
</script>

<template>
  <div class="phones-container">
    <div class="header">
      <h1>Список телефонов</h1>
      <div>
        <Button @click="fetchPhones" v-if="roleId == 1 || roleId == 2" severity="secondary" label="Обновить" />
        <Button @click="openCreateDialog" v-if="roleId == 1 || roleId == 2" severity="primary" label="Новый" />
      </div>
    </div>

    <DataTable :value="phones">
      <Column field="id" header="ID"></Column>
      <Column field="phone_number_value" header="Номер"></Column>
      <Column field="person_last_name" header="Фамилия">
        <template #body="{ data }">
          <Button 
            link 
            @click="openPersonPhones(data.person_id, data.person_last_name + ' ' + data.person_first_name)"
          >
            {{ data.person_last_name }}
          </Button>
        </template>
      </Column>
      <Column field="person_first_name" header="Имя"></Column>
      <Column field="phone_number_type" header="Тип"></Column>
      <Column field="comment" header="Комментарий"></Column>
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
            @click="deletePhone(data)"
          />
        </template>
      </Column>
    </DataTable>

    <PhoneDialog 
      v-model:visible="showDialog"
      :phone="selectedPhone"
      :is-edit="isEdit"
      :persons="persons"
      @saved="fetchPhones"
    />

    <PersonPhonesDialog 
      v-model:visible="showPersonPhonesDialog"
      :person-id="selectedPersonId"
      :person-name="selectedPersonName"
    />
  </div>
</template>

<style scoped>
.phones-container {
  width: 90%;
  max-width: 1400px;
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