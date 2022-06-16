import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import { defineNuxtPlugin } from '#app'

export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.$pinia.use(piniaPluginPersistedstate)
})