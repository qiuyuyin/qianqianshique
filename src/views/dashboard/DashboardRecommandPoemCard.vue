<template>
  <v-card
    v-bind="$attrs"
    class="align-self-start"
  >
    <v-card-text class="d-flex flex-column justify-center align-center">
      <v-avatar
        color="primary"
        class="mt-5 mx-auto"
        icon
        size="50"
      >
        <v-icon
          size="2rem"
          color="white"
        >
          {{ mdiFuriganaVertical }}
        </v-icon>
      </v-avatar>
    </v-card-text>
    <v-card-text class="text-center mb-n6 text--primary">
      <p class="font-weight-bold text-xl">
        {{ title }}
      </p>
    </v-card-text>
    <v-card-text>
      <div
        class="recommand-text-div text--primary"
      >
        <p class="secondary--text text-end ">
          {{ author }}
        </p>
        <span class="font-weight-bold">
          {{ content }}
        </span>
      </div>
    </v-card-text>
    <v-card-text class="text-center">
      <v-btn
        text
        color="primary"
        class="text-center font-weight-black"
        large
      >
        详情
      </v-btn>
    </v-card-text>
  </v-card>
</template>
<script>
import { mdiFuriganaVertical } from '@mdi/js'

export default {
  props: {
    content: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
    author: {
      type: String,
      required: true,
    },
    imgSrc: {
      type: String,
      required: true,
    },
  },
  setup() {
    return {
      mdiFuriganaVertical,
    }
  },
}
</script>
<style>
.recommand-text-div {
  height: 100px;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4;
  overflow: hidden;
}

</style>
