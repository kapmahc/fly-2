// import $ from 'jquery'

// const $ = window.jQuery
const $ = require('jquery')

// import {TOKEN} from './constants'

// $.ajaxSetup({
//   beforeSend: function (xhr) {
//     xhr.setRequestHeader('Authorization', `BEARER ${window.sessionStorage.getItem(TOKEN)}`)
//   }
// })

export const api = (path) => {
  return `${process.env.API_HOST}${path}`
}

const _fail = (e) => alert(e.responseText || e.statusText)

export const get = (path, success, fail) => {
  if (!fail) {
    fail = _fail
  }
  $.get(api(path), success).fail(fail)
}

export const _delete = (path, success, fail) => {
  if (!fail) {
    fail = _fail
  }
  $.ajax({
    url: api(path),
    type: 'DELETE',
    success
  }).fail(fail)
}

export const post = (path, body, success, fail) => {
  if (!fail) {
    fail = _fail
  }
  $.post(api(path), body, success).fail(fail)
}

// ---------------------------------

// const parse = (res) => {
//   // res.status === 200 || res.status === 0
//   return res.ok ? res.json() : res.text().then(err => { throw err })
// }
//
// const options = (method) => {
//   return {
//     method: method,
//     mode: 'cors',
//     credentials: 'include',
//     headers: {
//       'Authorization': `BEARER ${window.sessionStorage.getItem(TOKEN)}`
//     }
//   }
// }
//
// export const get = (path) => {
//   return fetch(api(path), options('get')).then(parse)
// }

// export const _delete = (path) => {
//   return fetch(api(path), options('delete')).then(parse)
// }

// export const post = (path, body) => {
//   var data = options('post')
//   data.body = body
//   console.log(data)
//   return fetch(api(path), data).then(parse)
// }
