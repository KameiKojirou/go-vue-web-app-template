<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useAuthTokenStore } from './stores/authTokenStore';

const auth = useAuthTokenStore();
</script>

<template>
  <header>
    <div>
     Token: {{ auth.token }}
      <nav class="flex flex-row gap-2 justify-between bg-base-300 p-4">
        <RouterLink class="btn btn-ghost" to="/">
          <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="64" height="64" />
        </RouterLink>
        <div class="flex flex-row gap-2 [&>a]:btn [&>a]:btn-ghost items-center">
          <RouterLink to="/">Home</RouterLink>
          <RouterLink to="/about">About</RouterLink>
          <RouterLink v-if="auth.token" to="/dashboard">Dashboard</RouterLink>
          <RouterLink v-if="!auth.token" to="/login">Login</RouterLink>
        </div>
      </nav>
    </div>
  </header>
  <div class="w-full flex flex-row justify-center">
    <div class="p-4 max-w-7xl w-full">
      <RouterView />
    </div>
  </div>
</template>
