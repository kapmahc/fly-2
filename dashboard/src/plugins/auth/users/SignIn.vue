<template>
<non-sign-in-layout title="auth.users.sign-in.title">
  <form>
    <div class="form-group">
      <label for="email">{{$t('attributes.email')}}</label>
      <input v-model="email" type="email" class="form-control" id="email" />
    </div>
    <div class="form-group">
      <label for="password">{{$t('attributes.password')}}</label>
      <input v-model="password" type="password" class="form-control" id="password" />
    </div>
    <button v-on:click="onSubmit" type="submit" class="btn btn-primary">{{$t('buttons.submit')}}</button>
  </form>
</non-sign-in-layout>
</template>

<script>
import {post} from '@/ajax'
import {TOKEN} from '@/constants'

export default {
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    onSubmit (e) {
      e.preventDefault()
      var data = new URLSearchParams()
      data.append('email', this.email)
      data.append('password', this.password)
      post('/users/sign-in', data)
        .then(function (rst) {
          var token = rst.token
          sessionStorage.setItem(TOKEN, token)
          this.$store.commit('signIn', token)
          this.$router.push({name: 'site.dashboard'})
        }.bind(this)).catch(alert)
    }
  }
}
</script>
