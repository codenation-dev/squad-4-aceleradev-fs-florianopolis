import api from './api';

export const listCustomer = (
    nome ) => {
    
    const list = [{id: 1, nome: 'Customer 1', cargo: 'Cargo', orgao: 'órgao', salario: 1000.1}, 
{id: 2, nome: 'Customer 2', cargo: 'Cargo', orgao: 'órgao', salario: 10.1}]

      if (nome) {
        return list.filter(function(customer) {
          return customer.nome.indexOf(nome) > 0;
        });
      }

      return list;

    /*return api.get('login', 
      {
        nome, orgao, funcao
      }
    )
    .then(response => response)
    .catch(err => {throw new Error(err)})*/
}

export default {
  listCustomer
}