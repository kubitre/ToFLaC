import {createStore, applyMiddleware} from 'redux';
import rootReducer from './Reducers';
import {logger} from './enchancers/logger';

export default function configureStore(initialState){
    const store = createStore(rootReducer, initialState, applyMiddleware(logger))

    if (module.hot){
        module.hot.accept('./Reducers', () => {
            const nextRootReducer = require('./Reducers')
            store.replaceReducer(nextRootReducer)
        })
    }
    return store
}