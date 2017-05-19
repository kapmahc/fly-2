<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="personalBar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{user ? $t('personal-bar.welcome', {name: user.name}) : $t('personal-bar.sign-in-or-up')}}
    </a>
    <div v-if="user" class="dropdown-menu" aria-labelledby="personalBar">
      <router-link class="dropdown-item" :to="{name: 'site.dashboard'}">
        {{$t('personal-bar.dashboard')}}
      </router-link>
      <div class="dropdown-divider"></div>
      <a v-on:click="onSignOut" class="dropdown-item">
        {{$t('personal-bar.sign-out')}}
      </a>
    </div>
    <div v-else class="dropdown-menu" aria-labelledby="personalBar">
      <router-link v-if="l" class="dropdown-item" v-for="l in links" v-bind:key="l" :to="{name: l.href}">
        {{$t(`${l.href}.title`)}}
      </router-link>
      <div v-else class="dropdown-divider"></div>
    </div>
  </li>
</template>

<script>
import {nonSignInLinks as links, TOKEN} from '@/constants'
import {_delete} from '@/ajax'

export default {
  data () {
    return {
      links
    }
  },
  beforeCreate () {
    var token = sessionStorage.getItem(TOKEN)
    if (token) {
      this.$store.commit('signIn', token)
    }
  },
  computed: {
    user () {
      return this.$store.state.currentUser
    }
  },
  methods: {
    onSignOut () {
      _delete('/users/sign-out', function (rst) {
        sessionStorage.clear()
        this.$store.commit('signOut')
        this.$router.push({name: 'site.home'})
      }.bind(this))
    }
  }
}
</script>
