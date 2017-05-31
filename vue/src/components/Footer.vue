<template>
  <footer>
    <el-row>
      <el-col :offset="1" :span="22">
        &copy; {{ $t("site.copyright") }}
        &middot; <a href="#">Privacy</a>
        &middot; <a href="#">Terms</a>
        {{$t('footer.other-languages')}}:
        <el-button type="text" v-bind:key="l" v-on:click="setLocale(l)" v-for="l in languages">
          {{$t(`languages.${l}`)}}
        </el-button>
      </el-col>
    </el-row>
  </footer>
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
