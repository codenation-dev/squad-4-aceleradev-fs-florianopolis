import React, { PureComponent } from 'react';
import {
  BarChart, Bar, Cell, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer
} from 'recharts';


export default class QtdDepart extends PureComponent {
  static jsfiddleUrl = 'https://jsfiddle.net/alidingling/31s5e83y/';

  render() {
    return (
      <div style={{width: '100%', height: 300}}>
        Funcionários por função e salário médio:
      <ResponsiveContainer>
      <BarChart
        data={this.props.data}
        margin={{
          top: 20, right: 30, left: 20, bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="cargo" />
        <YAxis yAxisId="left" orientation="left" stroke="#8884d8" />
        <YAxis yAxisId="right" orientation="right" stroke="#82ca9d" />
        <Tooltip />
        <Legend />
        <Bar yAxisId="left" dataKey="qtd" name="número de pessoas" fill="#8884d8" />
        <Bar yAxisId="right" dataKey="min" name="mínimo" stackId="a" fill="#ccc" />
        <Bar yAxisId="right" dataKey="avg" name="média salarial" stackId="a" fill="#82ca9d" />
        <Bar yAxisId="right" dataKey="max" name="máximo" stackId="a" fill="#82ca" />
      </BarChart>
      </ResponsiveContainer>
      </div>
    );
  }
}
