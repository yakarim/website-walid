import axios from '../services/myservices'
import tokens from '../services/auth-header';

export default {
    state: {
        token: localStorage.getItem('access_token') || null,
        user: null
    },
    getters: {
        loggedIn(state) {
            return state.token !== null
        }
    },
    mutations: {
        signIn(state, token) {
            state.token = token
        },
        destroyToken(state) {
            state.token = null
        }
    },
    actions: {
        destroyToken(context) {
            if (context.getters.loggedIn) {
                return new Promise((resolve, reject) => {
                    axios.get('/logoutdest', {
                        headers : tokens()
                      })
                        .then(res => {
                            localStorage.removeItem('access_token')
                            localStorage.removeItem('_id')
                            context.commit('destroyToken')
                            resolve(res)
                        }).catch(error => {
                            localStorage.removeItem('access_token')
                            localStorage.removeItem('_id')
                            context.commit('destroyToken')
                            console.log(error)
                            reject(error)
                        });
                })
            }
        },
        signIn({ commit }, credentials) {
            return new Promise((resolve, reject) => {
                axios.post('login', {
                    email: credentials.email,
                    password: credentials.password
                })
                    .then(res => {
                        const token = res.data.token
                        const _id = res.data.id
                        console.log(res)
                        localStorage.setItem('access_token', token)
                        localStorage.setItem('_id', _id)
                        commit('signIn', token)
                        resolve(res)
                    }).catch(error => {
                        console.log(error)
                        reject(error)
                    });
            })


        }
    }
}