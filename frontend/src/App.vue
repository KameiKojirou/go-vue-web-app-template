<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useAuthTokenStore } from './stores/authTokenStore';

const auth = useAuthTokenStore();
</script>

<template>
  <header>
    <div>
     Token: {{ auth.token }}
      <nav class="flex flex-row gap-2 justify-between">
        <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="64" height="64" />
        <div class="flex flex-row gap-2 [&>a]:btn [&>a]:btn-ghost">
          <RouterLink to="/">Home</RouterLink>
          <RouterLink to="/about">About</RouterLink>
          <RouterLink v-if="auth.token" to="/dashboard">Dashboard</RouterLink>
          <RouterLink v-if="!auth.token" to="/login">Login</RouterLink>
        </div>
      </nav>
    </div>
  </header>

  <RouterView />
</template>
