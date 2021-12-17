<template>
  <div>
    <p
      class="text-3xl primary--text font-weight-semibold"
    >
      作者
    </p>
    <v-row>
      <v-col
        v-if="authorList"
        cols="12"
        md="8"
      >
        <template
          v-for="data in authorList"
        >
          <v-hover
            v-slot="{ hover }"
            :key="data.id"
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
              class="mb-5 mt-2"
            >
            </author-card>
          </v-hover>
        </template>
        <v-pagination
          v-if="authorList"
          v-model="currentPage"
          class="mt-5"
          :length="pageLength"
          :total-visible="pageVisible"
          @input="onPageChange"
        >
        </v-pagination>
      </v-col>
      <v-col
        v-if="authorList"
        cols="12"
        md="4"
      >
        <v-card class="mt-2 pa-4">
          <p
            class="text-2xl primary--text font-weight-semibold"
          >
            诗文数目排名
          </p>
          <div
            v-for="(data,i) in poemCountList"
            :key="i"
            class="d-flex align-start pt-2"
          >
            <router-link
              :to="{ name: 'author-detail',
                     params: {
                       id: data.ID,
                     },
              }"
              class="author-router-link"
            >
              <p class="text-xl pr-2 poem-count-name">
                {{ data.name }}
              </p>
            </router-link>

            <v-progress-linear
              :color="colorList[i%10]"
              height="10"
              :value="data.poemCount / 100"
              striped
              class="mt-2 process-link mr-2"
            ></v-progress-linear>
            <p class="text-lg">
              {{ data.poemCount }}
            </p>
          </div>
          <br>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// eslint-disable-next-line object-curly-newline
import AuthorCard from './AuthorCard.vue'
import { getAuthorList, getAuthorPoemCount } from '@/api/author'

export default {
  components: {
    AuthorCard,
  },
  data() {
    return {
      authorList: undefined,
      currentPage: 1,
      pageSize: 5,
      pageLength: 20,
      pageVisible: 10,
      poemCountList: [],
    }
  },
  mounted() {
    this.getAuthorListByPage()
    this.getAuthorPoemCount()
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
    getAuthorPoemCount() {
      getAuthorPoemCount(15).then(res => {
        console.log(res)
        this.poemCountList = res.data
        console.log(this.poemCountList)
      })
    },
  },
  setup() {
    const colorList = ['#fab1a0', '#e17055', '#fdcb6e', 'primary', '#99A799', '#9D84B7', '#FF7777', '#B97A95', '#D9CAB3', '#B97A95']

    return {
      colorList,
    }
  },
}
</script>
<style type="text/css">
  .author-router-link {
    text-decoration: none;
  }
  .author-router-link:hover {
    cursor: pointer;
    text-decoration: underline;
    text-decoration-line: underline;
    text-decoration-thickness: initial;
    text-decoration-style: initial;
    text-decoration-color: initial;
  }

  .process-link {
    width: 230px;
  }
  .poem-count-name {
    min-width: 70px;
  }
  .poem-count-count {
    max-width: 210px;
  }

</style>
