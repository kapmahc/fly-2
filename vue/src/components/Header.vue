<template>
  <nav class="navbar navbar-toggleable-md navbar-inverse fixed-top bg-inverse">
    <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <router-link class="navbar-brand" :to="{name: 'site.home'}">
      {{$t('site.subTitle')}}
    </router-link>
    <div class="collapse navbar-collapse" id="navbarCollapse">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item active">
          <router-link class="nav-link" :to="{name: 'site.home'}">
            {{$t('header.home')}}
            <span class="sr-only">(current)</span>
          </router-link>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Link</a>
        </li>
        <LanguageBar />
        <PersonalBar />
      </ul>
      <SearchForm />
    </div>
  </nav>
</template>

<script>
import PersonalBar from './PersonalBar'
import LanguageBar from './LanguageBar'
import SearchForm from './SearchForm'

export default {
  data () {
    return {
    }
  },
  components: {
    PersonalBar,
    LanguageBar,
    SearchForm
  }
}
</script>
