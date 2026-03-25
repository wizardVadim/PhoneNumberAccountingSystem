<script setup lang="ts">
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { ref } from 'vue';
import { request } from '../../pkg/Request'
import { useToast } from 'primevue/usetoast'
import router from '@/router';

const login = ref("")
const password = ref("")

const toast = useToast()

const showSuccess = () => {
  toast.add({
    severity: 'success',
    summary: 'Успех',
    detail: 'Вы успешно авторизовались',
    life: 3000
  })
}

const showError = () => {
  toast.add({
    severity: 'error',
    summary: 'Ошибка',
    detail: 'Неверный логин или пароль',
    life: 3000
  })
}

const onFormSubmit = async () => {
  try {
    const data = await request('/login', {
      method: 'POST',
      body: { login: login.value, password: password.value }
    })

    console.log(data)
    
    localStorage.setItem('token', data.token)
    localStorage.setItem('roleId', data.role_id)
    showSuccess()
    router.push('/')
    
  } catch (error) {
    showError()
  }
}
</script>

<template>
  <div class="auth-container">
    <div class="header">
      <h1>Авторизуйтесь в системе</h1>
    </div>

    <form @submit.prevent="onFormSubmit" class="auth-form">
      <div class="input-wrapper">
        <InputText 
          v-model="login"
          type="text" 
          placeholder="Введите ваш логин" 
        />
      </div>
      <div class="input-wrapper">
        <InputText 
          v-model="password"
          type="password" 
          placeholder="Введите пароль" 
        />
      </div>
      <Button type="submit" severity="primary" label="Войти" />
    </form>
  </div>
</template>

<style scoped>
.auth-container {
  min-height: 80vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.header {
  text-align: center;
  margin-bottom: 2rem;
}

.auth-form {
  width: 100%;
  max-width: 400px;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(255, 255, 255, 0.1);
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