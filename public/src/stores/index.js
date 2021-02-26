import { createStore } from 'vuex'
import axios from '../services/myservices'

import auth from './auth'

export const store = createStore({
    state() {
        return {
            axios: axios,
        }
    },
    modules: {
        auth
    }
})