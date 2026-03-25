<!-- src/views/PersonsView.vue -->
<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import PersonDialog from '@/components/PersonsComponents/NewPhysPersonModal.vue';
import { request } from '../../pkg/Request';
import { useToast } from 'primevue/usetoast'
import PersonPhonesDialog from '@/components/PhonesComponents/PersonsPhones.vue';

const roleId = ref(-1)
const persons = ref([])
const showDialog = ref(false)
const isEdit = ref(false)
const selectedPerson = ref({})
const sort = ref(false)

const selectedPersonId = ref(0)
const showPersonPhonesDialog = ref(false)
const selectedPersonName = ref('')

const groupByPhone = ref(false)
const groupByPhoneName = ref("Показать количество номеров")

const sortName = ref("Отсортировать по ФИО")

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

const fetchPersons = async () => {
  try {
    if (sort.value) {
        const data = await request('/persons/sorted', { method: 'GET' })
        persons.value = data
    } else if (groupByPhone.value) {
        const data = await request('/persons/phone-stats', { method: 'GET' })
        persons.value = data
        console.log(data)
    } else {
        const data = await request('/persons', { method: 'GET' })
        persons.value = data
    }
  } catch (error) {
    console.error('Failed to fetch persons:', error)
  }
}

async function sortFIO () {
    sort.value = !sort.value

    if (sort.value) {
        groupByPhone.value = false
        groupByPhoneName.value = "Показать количество номеров"
        sortName.value = "Убрать сортировку"
    } else {
        sortName.value = "Отсортировать по ФИО"
    }

    await fetchPersons()
}

async function groupPhones () {
    groupByPhone.value = !groupByPhone.value

    if (groupByPhone.value) {
        sort.value = false
        sortName.value = "Отсортировать по ФИО"
        groupByPhoneName.value = "Убрать количество номеров"
    } else {
        groupByPhoneName.value = "Показать количество номеров"
    }

    await fetchPersons()
}

const openCreateDialog = () => {
  isEdit.value = false
  selectedPerson.value = {}
  showDialog.value = true
}

const openEditDialog = (person: any) => {
  console.log('Opening edit dialog for:', person)  
  console.log('Person ID:', person.id)  
  isEdit.value = true
  selectedPerson.value = person
  console.log('selectedPerson:', selectedPerson.value)  
  showDialog.value = true
  console.log('showDialog:', showDialog.value)  
}

const deletePerson = async (person: any) => {
  if (confirm(`Удалить ${person.last_name} ${person.first_name}?`)) {
    try {
      await request('/persons', { 
        method: 'DELETE', 
        body: { id: person.id } 
      })
      showSuccess("Успешное удаление")
      await fetchPersons()
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
    await fetchPersons()
})
</script>

<template>
  <div class="persons-container">
    <div class="header">
      <h1>Список физических лиц</h1>
      <div>
        <Button @click="fetchPersons" severity="secondary" label="Обновить" />
        <Button @click="sortFIO" severity="warn" :label="sortName" />
        <Button @click="groupPhones" severity="warn" :label="groupByPhoneName" />
        <Button @click="openCreateDialog" v-if="roleId == 1 || roleId == 3" severity="primary" label="Новый" />
      </div>
    </div>

    <DataTable :value="persons">
      <Column field="id" header="ID"></Column>
      <Column field="last_name" header="Фамилия"></Column>
      <Column field="first_name" header="Имя"></Column>
      <Column field="second_name" header="Отчество"></Column>
      <Column field="city" header="Город"></Column>
      <Column field="address" header="Адрес"></Column>
      <Column field="born_year" header="Год рождения"></Column>
      <Column v-if="groupByPhone" field="phone_number_quantity" header="Номеров"></Column>
      <Column header="Действия">
        <template #body="{ data }">
          <Button 
            v-if="roleId == 1 || roleId == 3"
            icon="pi pi-pencil" 
            severity="secondary" 
            rounded 
            text 
            @click="openEditDialog(data)"
            style="margin-right: 8px"
          />
          <Button 
            v-if="roleId == 1 || roleId == 3"
            icon="pi pi-trash" 
            severity="danger" 
            rounded 
            text 
            @click="deletePerson(data)"
          />
          <Button 
            icon="pi pi-phone" 
            severity="warn" 
            rounded 
            text 
            @click="openPersonPhones(data.id, data.last_name + ' ' + data.first_name)"
          />
        </template>
      </Column>
    </DataTable>

    <PersonDialog 
      v-model:visible="showDialog"
      :person="selectedPerson"
      :is-edit="isEdit"
      @saved="fetchPersons"
    />

    <PersonPhonesDialog 
      v-model:visible="showPersonPhonesDialog"
      :person-id="selectedPersonId"
      :person-name="selectedPersonName"
    />
  </div>
</template>

<style scoped>
.persons-container {
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