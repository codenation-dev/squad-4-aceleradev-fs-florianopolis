import React, { Component } from 'react';
import './App.css';
import Routes from './services/routes';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import '../node_modules/startbootstrap-sb-admin-2/css/sb-admin-2.min.css'

class App extends Component {
  render() {
    return (
      <Routes />
    );
  }
}

export default App;
