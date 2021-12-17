<template>
  <v-card
    v-bind="$attrs"
    class="align-self-start"
  >
    <!-- <v-card-text class="d-flex flex-column justify-center align-center">
      <v-avatar
        color="primary"
        class="mt-5 mx-auto"
        icon
        size="50"
      >
        <v-icon
          size="2rem"
          color="white"
          class="spin-icon"
        >
          {{ mdiBrightness5 }}
        </v-icon>
      </v-avatar>
    </v-card-text> -->
    <v-card-text class="text-center mb-n6 text--primary pt-7">
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
        <span class="font-weight-bold text-md">
          {{ content }}
        </span>
      </div>
    </v-card-text>
    <v-card-text class="text-center">
      <v-btn
        :to="returnRouter"
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
import { mdiBrightness5 } from '@mdi/js'

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
    id: {
      type: Number,
      required: true,
    },
    dynasty: {
      type: String,
      required: true,
    },
  },
  computed: {
    returnRouter() {
      return {
        name: 'poem',
        params: {
          dynasty: this.dynasty,
          id: this.id,
        },
      }
    },
  },
  setup() {
    return {
      mdiBrightness5,
    }
  },
}
</script>
<style type="text/css">
.recommand-text-div {
  height: 100px;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4;
  overflow: hidden;
}
@keyframes changDeg{
  0%{
    transform: rotate(0deg);
  }
  100%{
    transform: rotate(360deg);
  }
}
.spin-icon {
  animation:changDeg 8s linear 1.5s infinite;
}
</style>
