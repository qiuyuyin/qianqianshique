<template>
  <div>
    <p
      class="text-3xl primary--text font-weight-semibold"
    >
      诗句
    </p>
    <v-row>
      <v-col
        cols="12"
        md="7"
      >
        <v-row>
          <v-col
            v-for="data in poemList"
            :key="data.id"
            cols="12"
          >
            <v-hover
              v-slot="{ hover }"
              close-delay="50"
              open-delay="50"
            >
              <poem-card
                :id="data.ID"
                :elevation="hover ? 16 : 2"
                :class="{ 'on-hover': hover }"
                :dynasty="data.dynasty"
                :author-name="data.authorName"
                :content="data.content"
                :poem-type="data.poemType"
                :title="data.title"
              >
              </poem-card>
            </v-hover>
          </v-col>
        </v-row>
        <v-pagination
          v-if="poemList.length !== 0 "
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
import PoemCard from './PoemCard.vue'
import { getPoemList } from '@/api/poemType'

export default {
  components: {
    PoemCard,
  },
  data() {
    return {
      poemList: [], // 得到的poem集合
      page: {
        currentPage: 1, // 当前页面
        pageSize: 5, // 页面大小
        pageLength: 3, // 设置的显示分页数目
        pageVisible: 10, // 可视化的分页数目
      },
      poemType: 'shi',
    }
  },
  mounted() {
    this.getPoemListByPage()
  },
  methods: {
    getPoemListByPage() {
      getPoemList({
        page: this.page.currentPage,
        pageSize: this.page.pageSize,
        poemType: this.poemType,
      }).then(res => {
        console.log(res.data)
        this.poemList = res.data.list
        this.page.pageLength = Math.ceil(res.data.total / this.page.pageSize)
      })
    },
    async onPageChange(page) {
      this.page.currentPage = page
      await this.getPoemListByPage()
    },
  },
}
</script>
