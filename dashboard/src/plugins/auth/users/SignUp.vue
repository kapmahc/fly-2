<template>
  <non-sign-in-layout title="auth.users.sign-up.title">
    <form>
      <div class="form-group">
        <label for="name">{{$t('attributes.fullName')}}</label>
        <input v-model="name" type="text" class="form-control" id="name" />
      </div>
      <div class="form-group">
        <label for="email">{{$t('attributes.email')}}</label>
        <input v-model="email" type="email" class="form-control" id="email" aria-describedby="emailHelp" />
        <small id="emailHelp" class="form-text text-muted">{{$t('helpers.email')}}</small>
      </div>
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
      name: '',
      email: '',
      password: '',
      passwordConfirmation: ''
    }
  },
  methods: {
    onSubmit (e) {
      e.preventDefault()
      post(
        '/users/sign-up',
        {
          name: this.name,
          email: this.email,
          password: this.password,
          passwordConfirmation: this.passwordConfirmation
        },
        function (rst) {
          alert(this.$t('auth.messages.email-for-confirm'))
          this.$router.push({name: 'auth.users.sign-in'})
        }.bind(this)
      )
    }
  }
}
</script>
