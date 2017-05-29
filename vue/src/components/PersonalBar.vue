<template>
  <b-nav-item-dropdown right v-if="user">
    <template slot="text">
      {{$t('personal-bar.welcome', {name: user.name})}}
    </template>
    <b-dropdown-item to="#">Profile</b-dropdown-item>
    <b-dropdown-item to="#">Signout</b-dropdown-item>
  </b-nav-item-dropdown>
  <b-nav-item-dropdown right v-else="user">
    <template slot="text">
      {{$t('personal-bar.sign-in-or-up')}}
    </template>
    <b-dropdown-item v-if="l" v-for="l in links" v-bind:key="l" :to="{name: l.href}">
      {{$t(`${l.href}.title`)}}
    </b-dropdown-item>
    <b-dropdown-divider v-else />
  </b-nav-item-dropdown>
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
