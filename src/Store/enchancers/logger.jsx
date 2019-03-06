export const logger = store => next => action => {
    console.log(`Тип события: ${action.type}, данные события: ${action.payload}`)
    return next(action)
}