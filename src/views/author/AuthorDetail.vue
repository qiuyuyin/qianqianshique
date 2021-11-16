<template>
  <div>
    <h1>简单的作者页面</h1>
    <h2>{{ $route.params.id }}</h2>
    <h2>{{ authorDetail.name }}</h2>
    <h2>{{ authorDetail.lifetime }}</h2>
  </div>
</template>
<script>
import { getAuthorByID } from '@/api/getAuthor'

export default {
  data() {
    return {
      authorDetail: {},
    }
  },
  mounted() {
    this.getAuthor()
  },
  methods: {
    getAuthor() {
      getAuthorByID(this.$route.params.id)
        .then(res => {
          this.authorDetail = res.data
          console.log(this.authorDetail)
        })
    },
  },
}
</script>
