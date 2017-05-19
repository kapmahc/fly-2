<template>
  <non-sign-in-layout title="site.leave-words.new.title">
    <form>
      <div class="form-group">
        <label for="body">{{$t('attributes.body')}}</label>
        <textarea v-model="body" class="form-control" id="body" rows="6" aria-describedby="bodyHelp"></textarea>
        <small id="bodyHelp" class="form-text text-muted">{{$t('site.helpers.leave-word.body')}}</small>
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
      body: ''
    }
  },
  methods: {
    onSubmit (e) {
      e.preventDefault()
      var data = new URLSearchParams()
      data.append('body', this.body)
      data.append('type', 'text')
      post('/leave-words', data)
        .then(function (rst) {
          alert(this.$t('success'))
          this.body = ''
        }.bind(this)).catch(alert)
    }
  }
}
</script>
