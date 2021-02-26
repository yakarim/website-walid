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
        query(state) {
            state.user = null
        }
    },
    actions: {
        queryUser(context) {
            if (context.getters.loggedIn) {
                return new Promise((resolve, reject) => {
                    axios.get('/user/query', {
                        headers : tokens()
                      })
                        .then(res => {
                            context.commit('query')
                            resolve(res)
                        }).catch(error => {
                            context.commit('query')
                            reject(error)
                        });
                })
            }
        },
        createUser({ commit }, credentials) {
            return new Promise((resolve, reject) => {
                axios.post('/user/query', {
                    email: credentials.email,
                    password: credentials.password
                })
                    .then(res => {
                        const token = res.data.token
                        localStorage.setItem('access_token', token)
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