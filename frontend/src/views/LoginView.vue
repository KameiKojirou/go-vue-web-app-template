<template>
  <Title>Login</Title>
  <div>
    <button class="btn btn-primary" v-on:click="login">Login</button>
  </div>
</template>

<script lang="ts" setup>
import Title from '@/components/Title.vue';
import router from '@/router';
import { useAuthTokenStore } from '@/stores/authTokenStore';

const authTokenStore = useAuthTokenStore();

const login = () => {
  authTokenStore.setToken()
  router.push('/dashboard')
}

if (authTokenStore.getToken()) {
  router.push('/dashboard');
}
</script>
