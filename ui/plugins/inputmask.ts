import { defineNuxtPlugin } from '#app'
import Inputmask from 'inputmask'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.directive('inputmask', {
    mounted(el, binding) {
      new Inputmask(binding.value).mask(el)
    },
    updated(el, binding) {
      new Inputmask(binding.value).mask(el)
    },
  })
})
