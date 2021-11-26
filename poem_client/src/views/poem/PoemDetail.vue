<template>
  <div>
    <v-row>
      <!-- 诗句 -->
      <v-col
        cols="12"
        md="8"
      >
        <v-card class="pl-2 pr-1">
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
      <!-- 注释 -->
      <v-col
        v-if="poemDetail.annotation"
        cols="12"
        md="8"
      >
        <v-card class="pl-1">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary--text">
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
      <!-- 译文 -->
      <v-col
        v-if="poemDetail.translation"
        cols="12"
        md="8"
      >
        <v-card class="pl-1">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary--text">
                译文
              </div>
              <div
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
      <!-- 鉴赏 -->
      <v-col
        v-if="poemDetail.appreciation"
        cols="12"
        md="8"
      >
        <v-card class="pl-1">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary--text">
                鉴赏
              </div>
              <div
                class="mt-2 poem-detail-intro"
              >
                <p
                  v-for="(contentLine,i) in poemDetail.appreciation"
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
      <!-- 文评 -->
      <v-col
        v-if="poemDetail.comment"
        cols="12"
        md="8"
      >
        <v-card class="pl-1">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary--text">
                文评
              </div>
              <div
                class="mt-2 "
              >
                <p
                  v-for="(contentLine,i) in poemDetail.comment"
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
      <!-- 创作背景 -->
      <v-col
        v-if="poemDetail.intro"
        cols="12"
        md="8"
      >
        <v-card class="pl-1">
          <v-list-item>
            <v-list-item-content>
              <div class="font-weight-bold text-xl mb-2 pt-2 primary--text">
                创作背景
              </div>
              <div

                class="mt-2 poem-detail-intro"
              >
                <p
                  class="poem-detail-content text--primary font-weight-blod"
                >
                  {{ poemDetail.intro }}
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
  .poem-detail-intro {
    text-indent:2em;
    /* letter-spacing: 0.4px; */
    line-height: 1.4rem;
  }
</style>
