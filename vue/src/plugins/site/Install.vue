<template>
  <non-sign-in-layout title="site.install.title">
    <form>
      <div class="form-group">
        <label for="title">{{$t('site.attributes.title')}}</label>
        <input v-model="title" type="text" class="form-control" id="title" />
      </div>
      <div class="form-group">
        <label for="subTitle">{{$t('site.attributes.subTitle')}}</label>
        <input v-model="subTitle" type="text" class="form-control" id="subTitle" />
      </div>
      <div class="form-group">
        <label for="email">{{$t('attributes.email')}}</label>
        <input v-model="email" type="email" class="form-control" id="email" aria-describedby="emailHelp">
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
      title: '',
      subTitle: '',
      email: '',
      password: '',
      passwordConfirmation: ''
    }
  },
  methods: {
    onSubmit (e) {
      e.preventDefault()
      var data = new URLSearchParams()
      data.append('title', this.title)
      data.append('subTitle', this.subTitle)
      data.append('email', this.email)
      data.append('name', this.name)
      data.append('password', this.password)
      data.append('passwordConfirmation', this.passwordConfirmation)
      post('/install', data)
        .then(function (rst) {
          alert(this.$t('success'))
          this.$router.push({name: 'auth.users.sign-in'})
        }.bind(this)).catch(alert)
    }
  }
}
</script>
