<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="languageBar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{$t('language-bar.switch')}}
    </a>
    <div class="dropdown-menu" aria-labelledby="languageBar">
      <a class="dropdown-item" v-bind:key="l" v-on:click="setLocale(l)" v-for="l in languages">
        {{$t(`languages.${l}`)}}
      </a>
    </div>
  </li>
</template>

<script>
import {LANGUAGES, load as setLocale} from '@/i18n'
export default {
  data () {
    return {
      languages: LANGUAGES
    }
  },
  methods: {
    setLocale
  }
}
</script>
