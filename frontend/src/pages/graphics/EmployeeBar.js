import React, { PureComponent } from 'react';
import {
  BarChart, Bar, Cell, XAxis, YAxis, CartesianGrid, Tooltip, Legend,
} from 'recharts';


export default class EmployeeBar extends PureComponent {
  static jsfiddleUrl = 'https://jsfiddle.net/alidingling/30763kr7/';

  render() {
    return (
      <BarChart
        width={500}
        height={300}
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
        <Bar dataKey="qtd" fill="#8884d8" />
      </BarChart>
    );
  }
}
