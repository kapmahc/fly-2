<template>
  <el-menu theme="dark" class="el-menu-demo" mode="horizontal" @select="handleSelect">
    <el-menu-item index="home">{{$t('site.subTitle')}}</el-menu-item>
    <LanguageBar />
    <PersonalBar />
  </el-menu>
</template>

<script>
import PersonalBar from './PersonalBar'
import LanguageBar from './LanguageBar'

export default {
  data () {
    return {
    }
  },
  components: {
    PersonalBar,
    LanguageBar
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>
