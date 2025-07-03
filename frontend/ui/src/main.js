import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/main.css'

import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faLinkedin, faGithub, faLinux, faWindows, faApple, faGolang, faPython, faJava, faVuejs,
  faJs, faHtml5, faDocker, faGitAlt, faBootstrap
} from '@fortawesome/free-brands-svg-icons'

import {
  faFileArrowDown, faDesktop, faCode, faScrewdriverWrench, faLayerGroup, faDatabase,
  faGraduationCap
} from '@fortawesome/free-solid-svg-icons'

import { createPinia } from 'pinia'

library.add(
  faPython, faJava, faGolang, faVuejs, faLinux, faCode, faJs, faHtml5, faScrewdriverWrench,
  faApple, faWindows, faGithub, faLinkedin, faFileArrowDown, faDesktop, faDocker, faGitAlt, faLayerGroup,
  faBootstrap, faDatabase, faGraduationCap
)


const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)

app.component('font-awesome-icon', FontAwesomeIcon)

app.mount('#app')
