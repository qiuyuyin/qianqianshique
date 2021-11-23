<template>
  <div>
    <v-row>
      <v-col
        cols="12"
        md="8"
      >
        <v-card class="pl-4">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary-text">
                {{ poemDetail.title }}
              </div>
              <div class="text-sm">
                {{ poemDetail.authorName }}
              </div>
              <div
                class="text--primary mt-2"
              >
                <p
                  v-for="(contentLine,i) in poemDetail.content"
                  :key="i"
                  class="poem-detail-content"
                >
                  {{ contentLine }}
                </p>
              </div>
            </v-list-item-content>
          </v-list-item>
        </v-card>
      </v-col>
      <v-col
        cols="12"
        md="8"
      >
        <v-card class="pl-4">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary-text">
                注释
              </div>
              <div
                class="text--primary mt-2"
              >
                <p
                  v-for="(contentLine,i) in poemDetail.annotation"
                  :key="i"
                  class="poem-detail-content"
                >
                  {{ contentLine }}
                </p>
              </div>
            </v-list-item-content>
          </v-list-item>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>
<script>
import { getPoemByID } from '@/api/poemType'

export default {
  data() {
    return {
      poemDetail: {},
    }
  },
  mounted() {
    this.getPoem()
  },
  methods: {
    getPoem() {
      getPoemByID(this.$route.params.id, this.$route.params.poemType)
        .then(res => {
          this.poemDetail = res.data
        })
    },
  },
}
</script>
<style type="text/css">
  .poem-detail-content{
    font-size: 16px;
  }
</style>
