import React, { PureComponent } from 'react';
import {
  BarChart, Bar, Cell, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer
} from 'recharts';


export default class EmployeeBar extends PureComponent {

  render() {
    return (
      <div style={{width: '100%', height: 300}}>
      Histograma de salário do mês:
    <ResponsiveContainer>
      <BarChart
        data={this.props.data}
        margin={{
          top: 5, right: 30, left: 20, bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="floor" />
        <YAxis scale="log" domain={[0.01, 'auto']} allowDataOverflow />
        <Tooltip />
        <Legend />
        <Bar dataKey="qtd" fill="#8884d8" name="quantidade de pessoas" />
      </BarChart>
      </ResponsiveContainer>
      </div>
    );
  }
}
