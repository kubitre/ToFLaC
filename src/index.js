import React from 'react';
import ReactDOM from 'react-dom';

import './index.css';
import * as serviceWorker from './serviceWorker';

import {BrowserRouter} from 'react-router-dom';
import Routing from './Components/Router';

import {Provider} from 'react-redux';
import configureStore from './Store/index';

import 'bootstrap/dist/css/bootstrap.min.css';

const store = configureStore();

ReactDOM.render(
<Provider store={store}>
    <BrowserRouter>
        <Routing />
    </BrowserRouter>
</Provider>, document.getElementById('root'));

serviceWorker.unregister();
