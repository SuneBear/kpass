export const createEmptyPromise = () => {
  let callbacks
  const promiseInstance = new Promise((resolve, reject) => {
    callbacks = { resolve, reject }
  })
  promiseInstance.resolve = callbacks.resolve
  promiseInstance.reject = callbacks.reject
  return promiseInstance
}
