<template>
  <non-sign-in-layout title="auth.users.reset-password.title">
    <form>
      <div class="form-group">
        <label for="password">{{$t('attributes.password')}}</label>
        <input v-model="password" type="password" class="form-control" id="password" aria-describedby="passwordHelp"/>
        <small id="passwordHelp" class="form-text text-muted">{{$t('helpers.password')}}</small>
      </div>
      <div class="form-group">
        <label for="passwordConfirmation">{{$t('attributes.passwordConfirmation')}}</label>
        <input v-model="passwordConfirmation" type="password" class="form-control" id="passwordConfirmation" aria-describedby="passwordConfirmationHelp" />
        <small id="passwordConfirmationHelp" class="form-text text-muted">{{$t('helpers.passwordConfirmation')}}</small>
      </div>
      <button v-on:click="onSubmit" type="submit" class="btn btn-primary">{{$t('buttons.submit')}}</button>
    </form>
  </non-sign-in-layout>
</template>

<script>
import {post} from '@/ajax'

export default {
  data () {
    return {
      password: '',
      passwordConfirmation: ''
    }
  },
  methods: {
    onSubmit (e) {
      e.preventDefault()
      var data = new URLSearchParams()
      data.append('token', this.$route.params.token)
      data.append('password', this.password)
      data.append('passwordConfirmation', this.passwordConfirmation)
      post('/users/reset-password', data)
        .then(function (rst) {
          alert(this.$t('auth.messages.reset-password-success'))
          this.$router.push({name: 'auth.users.sign-in'})
        }.bind(this)).catch(alert)
    }
  }
}
</script>
