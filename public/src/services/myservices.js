import axios from 'axios'
//const URL = 'http://localhost:3000'
const URL = '.'

const apiClient = axios.create({
    baseURL: URL,
    withCredentials: false,
    timeout: 1000,
   
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
       // 'Authorization': 'Barier ' + btoa('api_username_here' + ':' + 'api_password_here')     
    }
})

export default apiClient