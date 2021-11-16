<template>
  <div>
    <v-row>
      <v-col
        cols="12"
        md="9"
      >
        <v-card>
          <div class="d-flex flex-no-wrap justify-space-between">
            <div
              v-if="authorDetail.avatar!=='https://song.gushiwen.cn/'&&$vuetify.breakpoint.mdAndUp"
              class="mt-6 ml-6 mb-6"
            >
              <v-img
                :src="authorDetail.avatar"
                class="rounded-lg"
                width="88px"
                height="100%"
              ></v-img>
            </div>
            <div class="mt-3">
              <v-list-item>
                <v-list-item-avatar
                  v-if="authorDetail.avatar!=='https://song.gushiwen.cn/'&&!$vuetify.breakpoint.mdAndUp"
                  size="35"
                >
                  <v-img
                    :src="authorDetail.avatar"
                  ></v-img>
                </v-list-item-avatar>

                <v-list-item-content>
                  <v-list-item-title class="font-weight-bold text-xl">
                    {{ authorDetail.name }}
                  </v-list-item-title>
                </v-list-item-content>
              </v-list-item>
              <v-card-text>
                <div
                  class="text--primary"
                >
                  <span class="font-weight-bold">
                    {{ authorDetail.lifetime }}
                  </span>
                </div>
              </v-card-text>
            </div>
          </div>
        </v-card>
      </v-col>

      <v-col
        v-for="data in authorDetail.describe"
        :key="data.id"
        cols="12"
        md="9"
      >
        <v-card>
          <v-card-title>
            {{ data.type }}
          </v-card-title>
          <div
            v-for="(content,index) in data.content"
            :key="index"
          >
            <v-card-text>
              {{ content }}
            </v-card-text>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>
<script>
import { getAuthorByID } from '@/api/getAuthor'

export default {
  data() {
    return {
      authorDetail: {},
    }
  },
  mounted() {
    this.getAuthor()
  },
  methods: {
    getAuthor() {
      getAuthorByID(this.$route.params.id)
        .then(res => {
          this.authorDetail = res.data

          // 对读取的数据进行筛选
          this.authorDetail.avatar = `https://song.gushiwen.cn/${this.authorDetail.avatar}`
          this.authorDetail.describe = this.authorDetail.describe.filter(data => data.content.length !== 0)
          console.log(this.authorDetail)
        })
    },
  },
}
</script>
