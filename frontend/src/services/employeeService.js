import api from './api';
import {logout} from './loginService';

export const listEmployee = (
    nome, funcao, orgao, wage, page) => {
    
    return api.get(`public_func?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&page=${page}`)
    .then(response => response.data)
    .catch(error => [])
}

export default {
  listEmployee
}