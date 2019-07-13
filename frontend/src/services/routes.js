import React from 'react';

import {BrowserRouter, Switch, Route} from 'react-router-dom';

import Alertas from '../pages/alertas';
import Dashboard from '../pages/dashboard';


const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Dashboard} />
            <Route path="/alertas" component={Alertas} />
        </Switch>
    </BrowserRouter>
);

export default Routes;