import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import router from './router'
import ApiClient from "@/util/apiclient";


Vue.config.productionTip = false
Vue.prototype.$API_CLIENT = new ApiClient(
    process.env.VUE_APP_SERVER_BASE_PATH)

new Vue({
    router,
    vuetify,
    render: h => h(App)
}).$mount('#app')

