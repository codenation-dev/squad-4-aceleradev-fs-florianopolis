import api from './api';

export const listEmployee = (
    nome, orgao, funcao ) => {
    
    return api.get(`funcionario?nome=${nome}`)
    .then(response => response)
    .catch(err => {throw new Error(err)})
}

export default {
  listEmployee
}