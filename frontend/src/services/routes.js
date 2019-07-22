import React from 'react';

import {BrowserRouter, Switch, Route} from 'react-router-dom';

import Warning from '../pages/warning';
import Dashboard from '../pages/dashboard';
import Employee from '../pages/employee';
import Customer from '../pages/customer';

import Login from '../pages/login'


const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Dashboard} />
            <Route path="/login" component={Login} />
            <Route path="/alertas" component={Warning} />
            <Route path="/employee" component={Employee} />
            <Route path="/customer" component={Customer} />
        </Switch>
    </BrowserRouter>
);

export default Routes;