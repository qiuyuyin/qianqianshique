<template>
  <v-card
    v-bind="$attrs"
    class="align-self-start"
  >
    <div>
      <div class="text-center">
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title class="font-weight-bold text-xl mb-2 pt-2">
              {{ title }}
            </v-list-item-title>
            <v-list-item-subtitle class="">
              {{ authorName }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-card-text>
          <div
            class="text--primary"
          >
            <p
              v-for="(contentLine,i) in content"
              :key="i"
              class="poem-card-content"
            >
              {{ contentLine }}
            </p>
          </div>
        </v-card-text>
      </div>
    </div>

    <v-card-text class="text-center">
      <v-btn
        :to="returnRouter"
        text
        color="primary"
        class="text-center font-weight-black"
        large
        target="_blank"
      >
        详情
      </v-btn>
    </v-card-text>
  </v-card>
</template>
<script>

export default {
  props: {
    content: {
      type: Array,
      required: true,
    },
    authorName: {
      type: String,
      required: true,
    },
    id: {
      type: String,
      required: true,
    },
    dynasty: {
      type: String,
      required: true,
    },
    poemType: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
  },
  computed: {
    returnRouter() {
      return {
        name: 'poem-detail',
        params: {
          poemType: this.poemType,
          id: this.id,
        },
      }
    },
  },
  mounted() {
  },
  setup() {
    return {
    }
  },
}
</script>
<style type="text/css">
  .poem-card-content{
    font-size: 16px;
  }
</style>
