import api from './api';

export const listCustomer = (
    nome ) => {
    
    return api.get(`cliente?nome=${nome}`)
    .then(response => response)
    .catch(err => {throw new Error(err)})
}

export default {
  listCustomer
}