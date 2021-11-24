<template>
  <div>
    <v-row>
      <v-col
        cols="12"
        md="8"
      >
        <v-card class="pl-2">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary--text">
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
        <v-card class="pl-2 pr-2">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary--text">
                注释
              </div>
              <div
                v-if="poemDetail.annotation"
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
      <v-col
        cols="12"
        md="8"
      >
        <v-card class="pl-2">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary--text">
                注释
              </div>
              <div
                v-if="poemDetail.translation"
                class="mt-2 "
              >
                <p
                  v-for="(contentLine,i) in poemDetail.translation"
                  :key="i"
                  class="poem-detail-content text--primary font-weight-blod"
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
    /* font-size: 16px; */
    line-height: 1.35rem;
  }
</style>
