import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/main.css'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { fab } from '@fortawesome/free-brands-svg-icons'
import { fas } from '@fortawesome/free-solid-svg-icons'

library.add(fab, fas)


const app = createApp(App)

app.use(router)

app.component('font-awesome-icon', FontAwesomeIcon)

app.mount('#app')
