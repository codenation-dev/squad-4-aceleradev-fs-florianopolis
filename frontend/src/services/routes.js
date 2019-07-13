import React from 'react';

import {BrowserRouter, Switch, Route} from 'react-router-dom';

import Warning from '../pages/warning';
import Dashboard from '../pages/dashboard';


const Routes = () => (
    <BrowserRouter>
        <Switch>
            <Route exact path="/" component={Dashboard} />
            <Route path="/alertas" component={Warning} />
        </Switch>
    </BrowserRouter>
);

export default Routes;