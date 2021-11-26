<template>
  <v-card
    class="greeting-card"
    height="180px"
  >
    <v-row class="ma-0 pa-0">
      <v-col cols="12">
        <div class="text-xl font-weight-semibold primary--text pt-2 ps-2 pb-2 mx-auto ">
          🥳一阙小诗
        </div>
      </v-col>
    </v-row>
    <div class="pl-10 text-xl font-weight-semibold pr-4">
      {{ everyDayDate.content }}
    </div>
    <p class="text-xl font-weight-semibold primary--text pr-5 mt-3 text-right">
      {{ getAuthor }}
    </p>
  </v-card>
</template>

<script>
const jinrishici = require('jinrishici')

jinrishici.load(result => {
  console.log(result)
})
export default {
  data() {
    return {
      everyDayDate: {
        content: '',
        origin: {
          author: '',
          dynasty: '',
        },
      },
    }
  },
  computed: {
    getAuthor() {
      if (this.everyDayDate.origin.author === '') { return '' }

      return `${this.everyDayDate.origin.author}(${this.everyDayDate.origin.dynasty})`
    },
  },

  // 浏览器初始化之后使用这个格式。
  mounted() {
    this.loadSentence()
  },
  methods: {
    loadSentence() {
      jinrishici.load(result => {
        this.everyDayDate = result.data
      }, errDate => {
        console.log(errDate)
      })
    },
  },
}
</script>

<style lang="scss" scoped>
.greeting-card {
  position: relative;
  .greeting-card-bg {
    position: absolute;
    bottom: 0;
    right: 0;
  }
  .greeting-card-trophy {
    position: absolute;
    bottom: 10%;
    right: 8%;
  }
}
// rtl
.v-application {
  &.v-application--is-rtl {
    .greeting-card-bg {
      right: initial;
      left: 0;
      transform: rotateY(180deg);
    }
    .greeting-card-trophy {
      left: 8%;
      right: initial;
    }
  }
}
</style>
