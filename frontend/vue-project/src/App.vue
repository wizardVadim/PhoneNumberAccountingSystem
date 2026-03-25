<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Toast from 'primevue/toast';

import { onBeforeMount } from 'vue'
import { verify } from '../pkg/VerifyMe'
import { useRouter } from 'vue-router'
 import Button from 'primevue/button';

const router = useRouter()

function logout () {
  localStorage.removeItem("token")
  localStorage.removeItem("roleId")
  router.push('/login')
} 

async function home () {
  const isValid = await verify()
  if (isValid) {
    router.push('/')
  }
} 

onBeforeMount(async () => {
  const isValid = await verify()
  if (!isValid) {
    router.push('/login')
  }
})
</script>

<template>
  <div class="nav">
    <div class="home">
      <Button v-on:click="home()" severity="secondary" icon="pi pi-home" label=""></Button>
    </div>
    <div class="logout">
      <Button v-on:click="logout()" severity="danger" label="Выйти"></Button>
    </div>
  </div>
  <div class="container">
    <RouterView />
  </div>
  <Toast />
</template>

<style scoped>

  .nav button {
    margin-left: 10px;
  }

  .nav {
    display: flex;
    justify-content: right;
    margin-top: 20px;
    margin-right: 20px;
  }

</style>
