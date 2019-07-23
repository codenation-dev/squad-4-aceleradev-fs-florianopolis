import api from './api';

export const listCustomer = (
  name, page ) => {
    
    return api.get(`customer?name=${name}&page=${page}`)
    .then(response => response)
    .catch(err => {throw new Error(err)})
}

export default {
  listCustomer
}