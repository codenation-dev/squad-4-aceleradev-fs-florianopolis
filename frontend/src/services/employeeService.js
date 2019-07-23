import api from './api';

export const listEmployee = (
    nome, funcao, orgao, wage, page) => {
    
    return api.get(`public_func?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&page=${page}`)
    .then(response => response)
    .catch(err => {throw new Error(err)})
}

export default {
  listEmployee
}