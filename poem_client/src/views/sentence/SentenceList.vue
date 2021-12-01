<template>
  <div>
    <p
      class="text-3xl primary--text font-weight-semibold"
    >
      名句
    </p>
    <v-row>
      <v-col
        cols="12"
        md="8"
      >
        <v-card>
          <v-hover
            v-for="(data,i) in sentenceList"
            :key="i"
            v-slot="{ hover }"
            close-delay="5"
            open-delay="5"
          >
            <sentence-card
              :id="data.ID"
              :elevation="hover ? 16 : 2"
              :class="{ 'on-hover': hover }"
              :dynasty="data.dynasty"
              :author-name="data.from"
              :content="data.content"
              :title="data.title"
              :author-id="data.authorID"
              :poem-id="data.poemID"
              :poem-type="data.poemtype"
            >
            </sentence-card>
          </v-hover>
        </v-card>
        <v-pagination
          v-if="sentenceList.length !== 0"
          v-model="page.currentPage"
          class="mt-5"
          :length="page.pageLength"
          :total-visible="page.pageVisible"
          @input="onPageChange"
        >
        </v-pagination>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// eslint-disable-next-line object-curly-newline
import SentenceCard from './SentenceCard.vue'
import { getSentenceList } from '@/api/sentence'

export default {
  components: { SentenceCard },
  data() {
    return {
      sentenceList: [], // 得到的sentence集合
      page: {
        currentPage: 1, // 当前页面
        pageSize: 10, // 页面大小
        pageLength: 3, // 设置的显示分页数目
        pageVisible: 10, // 可视化的分页数目
      },
      sentenceType: 'shi',
    }
  },
  mounted() {
    this.getSentenceListByPage()
  },
  methods: {
    getSentenceListByPage() {
      getSentenceList({
        page: this.page.currentPage,
        pageSize: this.page.pageSize,
      }).then(res => {
        console.log(res.data)
        this.sentenceList = res.data.list
        this.page.pageLength = Math.ceil(res.data.total / this.page.pageSize)
      })
    },
    async onPageChange(page) {
      this.page.currentPage = page
      await this.getSentenceListByPage()
    },
  },
}
</script>
