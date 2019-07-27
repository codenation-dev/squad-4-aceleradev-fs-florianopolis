import api from './api';
import sha256 from 'crypto-js/sha256';

const HASH = 'b340b3c69a177bc5508c842aa038afc252692dee';

export const login = (
    email,
    password
) => {
    if (!email){
        throw new Error('Você precisa informar o email.')
    }
    if (!password){
        throw new Error('Você precisa informar a senha.')
    }
    
    const shaPassword = sha256(HASH + password);
    return api.post('login', 
      {
        email: email,
        password: shaPassword.toString(),   
      },
      {withCredentials: false}
    )
    .then(response => {
        localStorage.setItem("authToken", response.data.token)
        document.cookie = `token=${response.data.token};max-age=604800;`
    })
    .catch(err => console.log(err))    
}

export const register = (
    email,
    password
) => {
    if (!email){
        throw new Error('Você precisa informar o email.')
    }
    if (!password){
        throw new Error('Você precisa informar a senha.')
    }

    const shaPassword = sha256(HASH + password);
    return api.post('user',
    {        
        email: email,
        password: shaPassword.toString(),   
      }
    )
    .then(response => response)
    .catch(err => {throw new Error(err)})    
}

export const logout = () => {
    localStorage.clear();
    document.cookie.remove('token');
}

export default {
    register,
    login,
    logout
}