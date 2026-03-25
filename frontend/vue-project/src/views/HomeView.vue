<script setup lang="ts">
  import { onBeforeMount, ref } from 'vue';
  import Button from 'primevue/button';
  import router from '@/router';

  const roleId = ref(-1)

  function openUserList() {
    router.push("/users")
  }

  function openPersonList() {
    router.push("/persons")
  }

  function openPhoneList() {
    router.push("/phones")
  }

  onBeforeMount(async () => {
    roleId.value = Number(localStorage.getItem("roleId"))
    console.log(roleId.value)
  })

</script>

<template>
  <div class="main-menu-container">
    <div class="header">
      <h1>Главное меню</h1>
    </div>
    <div class="control-panel">
      <div class="btn-container">
        <Button @click="openUserList()" v-if="roleId == 1" severity="primary" label="f"> <i class="pi pi-cog" style="font-size: 100px;"></i> </Button>
        <Button v-else disabled severity="primary" label="f"> <i class="pi pi-cog" style="font-size: 100px;"></i> </Button>
        <p>Настройки пользователей</p>
      </div>
      <div class="btn-container">
        <Button @click="openPersonList()" v-if="roleId == 1 || roleId == 2 || roleId == 3" severity="primary" label="f"> <i class="pi pi-user" style="font-size: 100px;"></i> </Button>
        <Button v-else disabled severity="primary" label="f"> <i class="pi pi-user" style="font-size: 100px;"></i> </Button>
        <p>Справочник физических лиц</p>
      </div>
      <div class="btn-container"> 
        <Button @click="openPhoneList()" v-if="roleId == 1 || roleId == 2" severity="primary" label="f"> <i class="pi pi-phone" style="font-size: 100px;"></i> </Button>
        <Button v-else disabled severity="primary" label="f"> <i class="pi pi-phone" style="font-size: 100px;"></i> </Button>
        <p>Справочник телефонных номеров</p>
      </div>
    </div>
  </div>
</template>

<style scoped>

  .btn-container p {
    text-align: center;
  }

  .btn-container button {
    margin-bottom: 10px;
    padding: 40px;
  }

  .control-panel {
    display: flex;
  }

  .main-menu-container {
    min-height: 80vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .btn-container {
    margin: 0 60px;
    width: 33%;
  }

  .header {
    text-align: center;
    margin-bottom: 2rem;
  }

  .input-wrapper {
    margin-bottom: 1rem;
    padding: 4px;
  }

  .input-wrapper input {
    width: 100%;
  }

  button {
    width: 100%;
  }
</style>