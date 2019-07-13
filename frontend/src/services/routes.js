import React from 'react';

import {BrowserRouter, Switch, Route} from 'react-router-dom';

import Warning from '../pages/warning';
import Dashboard from '../pages/dashboard';

import Login from '../pages/login'


const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Dashboard} />
            <Route path="/login" component={Login} />
            <Route path="/alertas" component={Warning} />
        </Switch>
    </BrowserRouter>
);

export default Routes;