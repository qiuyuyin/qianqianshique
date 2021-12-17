<template>
  <v-card
    class="greeting-card"
    height="180px"
  >
    <v-card-title class="pl-n2">
      <div class="text-xl font-weight-semibold primary--text mx-auto ">
        🥳一阙小诗
      </div>
      <v-spacer></v-spacer>
      <v-btn
        icon
        small
        class="mt-n2 me-n3"
        @click="loadSentence"
      >
        <v-icon size="22">
          {{ mdiCached }}
        </v-icon>
      </v-btn>
    </v-card-title>

    <div class="pl-10 text-xl font-weight-semibold pr-4">
      {{ everyDayDate.content }}
    </div>
    <p class="text-xl font-weight-semibold primary--text pr-5 mt-3 text-right">
      {{ getAuthor }}
    </p>
  </v-card>
</template>

<script>
import { mdiCached } from '@mdi/js'

const jinrishici = require('jinrishici')

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

    this.$store.dispatch('snackbar/openSnackbar', {
      msg: '欢迎来到千千诗阕',
      color: 'primary',
    })
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
  setup() {
    return {
      mdiCached,
    }
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
