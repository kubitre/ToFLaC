import React from 'react';
import {Switch, Route} from 'react-router-dom';
import MainPageComponent from '../../Containers/MainPage';

export default function Routing(){
    return (
        <Switch>
            <Route exact to='/' component={MainPageComponent}/>
        </Switch>
    )
}