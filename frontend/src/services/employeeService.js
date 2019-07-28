import api from './api';
import {logout} from './loginService';

export const listEmployee = (
    nome, funcao, orgao, wage, ehCliente, relevancia, page, campoOrdenacao, ordenacao) => {
    
    return api.get(`public_func?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&customer=${ehCliente}&relevancia=${relevancia}&page=${page}&sortby=${campoOrdenacao}&desc=${ordenacao}`)
    .then(response => response.data)
    .catch(error => [])
}

export const statEmployee = (
  nome, funcao, orgao, wage, ehCliente, relevancia) => {
  
  return api.get(`public_func/stats?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&customer=${ehCliente}&relevancia=${relevancia}`)
  .then(response => response.data)
  .catch(error => [])
}

export const departStatEmployee = (
  nome, funcao, orgao, wage, ehCliente, relevancia) => {
  
  return api.get(`public_func/distribution?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&customer=${ehCliente}&relevancia=${relevancia}`)
  .then(response => response.data)
  .catch(error => [])
}

export default {
  listEmployee,
  statEmployee,
  departStatEmployee,
}