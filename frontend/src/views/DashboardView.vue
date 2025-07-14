<template>
  <h1>DashboardView</h1>
  <div>
    Welcome
    <button class="btn btn-error" v-on:click="logout">Logout</button>
  </div>
</template>

<script lang="ts" setup>
import router from '@/router';
import { useAuthTokenStore } from '@/stores/authTokenStore';

const auth = useAuthTokenStore();

const logout = async () => {
  await fetch('http://localhost:1323/api/logout', {

  }).then(() => {

  })
  auth.removeToken()
  router.push('/login')
}

if (!auth.getToken()) {
  router.push('/login');
}

</script>
