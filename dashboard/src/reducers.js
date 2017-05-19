import {USERS_SIGN_IN, USERS_SIGN_OUT, SITE_REFRESH} from './actions'
import jwtDecode from 'jwt-decode'


const currentUser = (state={}, action) => {
  switch(action.type){
    case USERS_SIGN_IN:
      try{
        return jwtDecode(action.token)
      }catch(e){
        console.error(e)
      }
      return {}
    case USERS_SIGN_OUT:
      return {}
    default:
      return state
  }
}

const siteInfo = (state={languages:[]}, action) => {
  switch(action.type){
    case SITE_REFRESH:
      return Object.assign({}, action.info)
    default:
      return state;
  }
}

export default {
  currentUser,
  siteInfo
}
