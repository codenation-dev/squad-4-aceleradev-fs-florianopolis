import api from './api';

export const sendEmail = (
    nome, funcao, orgao, wage, ehCliente, relevancia) => {
    
    return api.get(`email_to?nome=${nome}&cargo=${funcao}&orgao=${orgao}&salario=${wage}&customer=${ehCliente}&relevancia=${relevancia}`)
    .then(response => response.data)
    .catch(error => [])
}

export default {
  sendEmail
}