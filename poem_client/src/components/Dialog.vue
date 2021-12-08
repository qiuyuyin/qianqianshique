<template>
  <v-snackbar
    v-model="visible"
    top
    right
    text
    class="pt-12"
    :color="color"
    elevation="24"
  >
    {{ this.$store.state.snackbar.msg }}
    <!-- 关闭按钮 -->
    <template v-slot:action="{ attrs }">
      <v-btn
        v-if="showClose"
        v-bind="attrs"
        icon
        :color="color"
        @click="close"
      >
        <v-icon>{{ mdiClose }}</v-icon>
      </v-btn>
    </template>
  </v-snackbar>
</template>
<script>
import {
  mdiClose,
} from '@mdi/js'

export default {
  data() {
    return {}
  },
  computed: {
    visible() {
      return this.$store.state.snackbar.visible
    },
    showClose() {
      return this.$store.state.snackbar.showClose
    },
    color() {
      return this.$store.state.snackbar.color
    },
  },
  methods: {
    close() {
      this.$store.commit('snackbar/CLOSE_SNACKBAR')
    },
  },
  setup() {
    return {
      mdiClose,
    }
  },
}
</script>
// 调用试例
// this.$store.dispatch('snackbar/openSnackbar', {
//       msg: '欢迎来到千千诗阕',
//       color: 'primary',
//     })
