import React from 'react';

import {BrowserRouter, Switch, Route} from 'react-router-dom';

import Alertas from '../pages/alertas';
import Dashboard from '../pages/dashboard';

import Login from '../pages/login'


const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Dashboard} />
            <Route path="/alertas" component={Alertas} />
            <Route path="/login" component={Login} />
        </Switch>
    </BrowserRouter>
);

export default Routes;