<template>
  <application-layout>
    <el-col :span="12" :offset="6">
      <h2>{{$t(title)}}</h2>
      <hr />
      <slot />
      <ul class="list-group">
        <li class="list-group-item" v-if="l" v-for="(l, i) in links" :key="i">
          <fa-icon :name="l.icon" />&nbsp;
          <router-link :to="{name: l.href}">
            {{$t(`${l.href}.title`)}}
          </router-link>
        </li>
      </ul>
    </el-col>
  </application-layout>
</template>

<script>
import {nonSignInLinks as links} from '@/constants'

export default {
  props: ['title'],
  data () {
    return {
      links
    }
  }
}
</script>
