<template>
  <div>
    <p
      class="text-3xl primary--text font-weight-semibold"
    >
      作者
    </p>
    <v-row>
      <v-col
        cols="12"
        md="8"
      >
        <v-row>
          <v-col
            v-for="data in authorList"
            :key="data.id"
            cols="12"
          >
            <v-hover
              v-slot="{ hover }"
              close-delay="50"
              open-delay="50"
            >
              <author-card
                :id="data.ID"
                :elevation="hover ? 16 : 2"
                :class="{ 'on-hover': hover }"
                :dynasty="data.dynasty"
                :name="data.name"
                :content="data.lifetime"
                :img-src="'https://song.gushiwen.cn/'+data.avatar"
              >
              </author-card>
            </v-hover>
          </v-col>
        </v-row>
        <v-pagination
          v-model="currentPage"
          class="mt-5"
          :length="pageLength"
          :total-visible="pageVisible"
          @input="onPageChange"
        >
        </v-pagination>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// eslint-disable-next-line object-curly-newline
import AuthorCard from './AuthorCard.vue'
import { getAuthorList } from '@/api/getAuthor'

export default {
  components: {
    AuthorCard,
  },
  data() {
    return {
      authorList: [],
      currentPage: 1,
      pageSize: 5,
      pageLength: 20,
      pageVisible: 10,
    }
  },
  mounted() {
    this.getAuthorListByPage()
  },
  methods: {
    getAuthorListByPage() {
      getAuthorList({
        page: this.currentPage,
        pageSize: this.pageSize,
      }).then(res => {
        console.log(res)
        this.authorList = res.data.list
        console.log(this.authorList)
        this.pageLength = Math.ceil(res.data.total / this.pageSize)
      })
    },
    async onPageChange(page) {
      this.currentPage = page
      await this.getAuthorListByPage()
    },
  },
}
</script>
