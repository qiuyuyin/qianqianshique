<template>
  <v-card
    v-bind="$attrs"
  >
    <v-img
      height="140px"
      :src="imgSrc"
    ></v-img>
    <v-card-title>
      {{ title }}
    </v-card-title>
    <v-card-text>
      <div
        class="recommand-text-div text--primary"
      >
        <p class="text-end primary--text">
          {{ author }}
        </p>
        <span>
          {{ content }}
        </span>
      </div>
    </v-card-text>
    <v-card-actions class="primary pa-0">
      <v-btn
        color="primary"
        block
        dark
        large
      >
        查看详情
      </v-btn>
    </v-card-actions>
  </v-card>
</template>
<script>
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
