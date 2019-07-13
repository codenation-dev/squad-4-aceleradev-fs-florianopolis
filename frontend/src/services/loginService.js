import api from './api';
import sha256 from 'crypto-js/sha256';
const HASH = 'b340b3c69a177bc5508c842aa038afc252692dee';

export const login = ({
    email,
    password
}) => {
    if (!email){
        throw new Error('Você precisa informar o email.')
    }
    if (!password){
        throw new Error('Você precisa informar a senha.')
    }
    
    const shaPassword = sha256(HASH + password);
    api.post('login', 
      {
        email: email,
        pass: shaPassword.toString(),
      }
    )
    .then(response => console.log(response))
    .catch(err => {throw new Error('Invalid User/Password')})    
}

export const register = ({
    email,
    password
}) => {
    if (!email){
        throw new Error('Você precisa informar o email.')
    }
    if (!password){
        throw new Error('Você precisa informar a senha.')
    }

    const shaPassword = sha256(HASH + password);
    api.post('PostUser', 
      {
        email: email,
        pass: shaPassword,
      }
    )
    .then(response => console.log(response))
    .catch(err => {throw new Error('Invalid User/Password')})    
    return user
}

export const isLogged = () => !!localStorageWrapper.get(NS_LOGGED_USER)

export const logout = () => localStorageWrapper.set(NS_LOGGED_USER, null)

export const getUser = () => isLogged && localStorageWrapper.get(NS_LOGGED_USER)

export default {
    register,
    login,
    isLogged,
    getUser
}