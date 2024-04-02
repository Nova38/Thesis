import VueJsonPretty from 'vue-json-pretty'


import JsonEditorVue from 'json-editor-vue'

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.vueApp.use(JsonEditorVue, {
    // global props & attrs (one-way data flow)
  })
  nuxtApp.vueApp.component('vue-json-pretty', VueJsonPretty)


})
