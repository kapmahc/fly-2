<template>
  <non-sign-in-layout :title="`auth.users.${action}.title`">
    <form>
      <div class="form-group">
        <label for="email">{{$t('attributes.email')}}</label>
        <input v-model="email" type="email" class="form-control" id="email" />
      </div>
      <button v-on:click="onSubmit" type="submit" class="btn btn-primary">{{$t('buttons.submit')}}</button>
    </form>
  </non-sign-in-layout>
</template>

<script>
import {post} from '@/ajax'

export default {
  props: ['action'],
  data () {
    return {
      email: ''
    }
  },
  methods: {
    onSubmit (e) {
      e.preventDefault()
      var data = new URLSearchParams()
      data.append('email', this.email)
      post(`/users/${this.action}`, data)
        .then(
          function (rst) {
            alert(this.$t(`auth.messages.email-for-${this.action}`))
            this.$router.push({name: 'auth.users.sign-in'})
          }.bind(this)
        ).catch(alert)
    }
  }
}
</script>
