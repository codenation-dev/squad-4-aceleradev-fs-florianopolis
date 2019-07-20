import React from 'react';

const RenderEmployee = (props) => {
    if(props.list.length === 0){
        return <p>Nenhum </p>
    }
    return props.list.map((customer, key) => {
        const { id, nome, cargo, orgao, salario } = customer
        return (
            <tr key={id} className="">
                <td className="col">{nome}</td>
                <td className="col">{cargo}</td>
                <td className="col">{orgao}</td>
                <td className="col text-right">{salario.toLocaleString('pt-BR', {"minimumFractionDigits": 2})}</td>
            </tr>
        )
    })
}

export default RenderEmployee