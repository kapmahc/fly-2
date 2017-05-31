<template>
  <div>
    <app-header />
    <el-row>
      <slot />
    </el-row>
    <hr />
    <app-footer />
  </div>
</template>

<script>
import Header from '@/components/Header'
import Footer from '@/components/Footer'

export default {
  data () {
    return {
    }
  },
  components: {
    'app-header': Header,
    'app-footer': Footer
  }
}
</script>
