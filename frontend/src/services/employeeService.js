import api from './api';
import {logout} from './loginService';

export const listEmployee = (
    nome, funcao, orgao, wage, ehCliente, page, campoOrdenacao, ordenacao) => {
    
    return api.get(`public_func?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&customer=${ehCliente}&page=${page}&sortby=${campoOrdenacao}&desc=${ordenacao}`)
    .then(response => response.data)
    .catch(error => [])
}

export default {
  listEmployee
}