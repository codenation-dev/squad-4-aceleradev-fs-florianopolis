import React from 'react';

import {BrowserRouter, Switch, Route} from 'react-router-dom';

import Warning from '../pages/warning';
import Dashboard from '../pages/dashboard';
import Employee from '../pages/employee';
import Customer from '../pages/customer';

import Login from '../pages/login'
import {PrivateRoute} from '../helpers/PrivateRoute';

const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route path="/login" component={Login} />
            <PrivateRoute exact path="/" component={Dashboard} />            
            <PrivateRoute exact Route path="/alertas" component={Warning} />
            <PrivateRoute exact path="/employee" component={Employee} />
            <PrivateRoute exact path="/customer" component={Customer} />
        </Switch>
    </BrowserRouter>
);

export default Routes;