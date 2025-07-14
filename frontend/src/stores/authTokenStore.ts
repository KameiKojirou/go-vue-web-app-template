import { defineStore } from "pinia"
import { ref } from "vue"
import Cookies from "js-cookie"

export const useAuthTokenStore = defineStore('auth', () => {
  const token = ref()

  function getToken() {
    token.value = Cookies.get('token')
    console.log(token.value)
    return token.value
  }

  async function setToken() {
    await fetch('http://localhost:1323/api/login', {
    }).then((response) => {
      return response.json()
    }).then((data) => {
      return data
    })
    token.value = Cookies.get('token')
    console.log(token.value)
  }

  function removeToken() {
    Cookies.remove('token')
    console.log(token.value)
    token.value = ''
  }

  return { token, getToken, setToken, removeToken }
})
