<template>
  <v-card
    v-bind="$attrs"
    class="align-self-start"
  >
    <div>
      <div class="d-flex flex-no-wrap justify-space-between">
        <div class="mt-4 ml-6">
          <v-img
            v-if="$vuetify.breakpoint.mdAndUp"
            :src="imgSrc"
            class="rounded-lg"
            width="80px"
            height="115px"
          ></v-img>
        </div>
        <div class="mt-3">
          <v-list-item>
            <v-list-item-avatar
              v-if="!$vuetify.breakpoint.mdAndUp"
              size="35"
            >
              <v-img
                :src="imgSrc"
              ></v-img>
            </v-list-item-avatar>

            <v-list-item-content>
              <v-list-item-title class="font-weight-bold text-xl">
                {{ name }}
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
          <v-card-text>
            <div
              class="recommand-text-div text--primary"
            >
              <span class="font-weight-bold">
                {{ content }}
              </span>
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
        >
          详情
        </v-btn>
      </v-card-text>
    </div>
  </v-card>
</template>
<script>

export default {
  props: {
    content: {
      type: String,
      required: true,
    },
    name: {
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
    }
  },
}
</script>
<style type="text/css">
.recommand-text-div {
  height: 83px;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 4;
  overflow: hidden;
}

</style>
