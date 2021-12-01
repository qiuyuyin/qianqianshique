<template>
  <div>
    <p class="text-3xl primary--text font-weight-semibold">
      诗句
    </p>

    <v-row>
      <v-col
        cols="12"
        md="8"
      >
        <div class="d-flex align-center">
          <p class="mt-3 mr-4">
            诗句筛选:
          </p>
          <v-menu
            offset-y
            :close-on-content-click="false"
          >
            <template v-slot:activator="{ on }">
              <v-text-field
                full-width
                height="35px"
                dense
                outlined
                class="app-bar-search"
                hide-details
                placeholder="点击筛选"
                v-on="on"
              >
                <template v-slot:prepend-inner>
                  <v-chip
                    v-if="typeOne"
                    filter
                    outlined
                    small
                    close
                    class="mr-2"
                    @click:close="typeOne = ''"
                  >
                    {{ typeOne }}
                  </v-chip>
                  <v-chip
                    v-if="typeTwo"
                    filter
                    outlined
                    small
                    class="mr-2"
                    close
                    @click:close="typeTwo = ''"
                  >
                    {{ typeTwo }}
                  </v-chip>
                  <v-chip
                    v-if="typeThree"
                    filter
                    outlined
                    small
                    close
                    @click:close="typeThree = ''"
                  >
                    {{ typeThree }}
                  </v-chip>
                </template>
              </v-text-field>
            </template>

            <v-list>
              <v-chip-group
                v-model="typeOne"
                column
                active-class="primary--text"

                class="pa-3 mt-n3 mb-n1"
              >
                <span class="pr-5 pt-1">题材:</span>
                <v-chip
                  v-for="(chip,i) in poemGroup.chipGroupTypeOne"
                  :key="i"
                  filter
                  outlineds
                  small
                  :value="chip"
                >
                  {{ chip }}
                </v-chip>
              </v-chip-group>
              <v-divider></v-divider>
              <v-chip-group
                v-model="typeTwo"
                active-class="primary--text"
                column
                class="pa-3 mt-n1 mb-n2"
              >
                <span class="pr-5 pt-1">类型:</span>
                <v-chip
                  v-for="(chip,i) in poemGroup.chipGroupTypeTwo"
                  :key="i"
                  filter
                  outlineds
                  small
                  :value="chip"
                >
                  {{ chip }}
                </v-chip>
              </v-chip-group>
              <v-divider></v-divider>
              <v-chip-group
                v-model="typeThree"
                active-class="primary--text"
                class="pa-3 mt-n1 mb-n2"
              >
                <span class="pr-5 pt-1">朝代:</span>
                <v-chip
                  v-for="(chip,i) in poemGroup.chipGroupTypeThree"
                  :key="i"
                  filter
                  outlineds
                  small
                  :value="chip"
                >
                  {{ chip }}
                </v-chip>
              </v-chip-group>
            </v-list>
          </v-menu>
        </div>

        <v-divider class="mb-2 mt-3"></v-divider>
        <span class="pr-5">诗句类型:</span>
        <v-btn-toggle
          v-model="poemType"
          dense
          color="primary"
          text-xl
          class="mb-2 mt-2"
        >
          <v-btn value="shi">
            <span>诗</span>
          </v-btn>
          <v-btn
            value="ci"
          >
            <span>词</span>
          </v-btn>
          <v-btn value="qu">
            <span>曲</span>
          </v-btn>
        </v-btn-toggle>
        <v-row v-if="poemList != null">
          <v-col
            v-for="data in poemList"
            :key="data.id"
            cols="12"
            class="mt-2"
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
          v-if="poemList"
          v-model="page.currentPage"
          class="mt-5"
          :length="page.pageLength"
          :total-visible="page.pageVisible"
          @input="onPageChange"
        >
        </v-pagination>
        <template v-if="!poemList">
          <p></p>
          <p></p>
          <p>暂无信息</p>
        </template>
      </v-col>
    </v-row>
  </div>
</template>

<script>
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
        pageSize: 6, // 页面大小
        pageLength: 3, // 设置的显示分页数目
        pageVisible: 10, // 可视化的分页数目
      },
      poemType: 'shi',
      typeOne: undefined,
      typeTwo: undefined,
      typeThree: undefined,
      poemGroup: {},
    }
  },
  watch: {
    typeThree() {
      this.page.currentPage = 1
      this.getPoemListByPage()
    },
    typeTwo() {
      this.page.currentPage = 1
      this.getPoemListByPage()
    },
    typeOne() {
      this.page.currentPage = 1
      this.getPoemListByPage()
    },
    poemType(newType, oldType) {
      this.page.currentPage = 1
      this.getPoemListByPage()
      console.log(oldType)
      if (newType === 'shi') {
        this.poemGroup = this.poemShi
      }
      if (newType === 'ci') {
        this.poemGroup = this.poemCi
      }
      if (newType === 'qu') {
        this.poemGroup = this.poemQu
      }
    },
  },
  mounted() {
    this.getPoemListByPage()
    this.poemGroup = this.poemShi
  },
  methods: {
    getPoemListByPage() {
      getPoemList({
        page: this.page.currentPage,
        pageSize: this.page.pageSize,
        poemType: this.poemType,
        tag: this.typeOne,
        poemStyle: this.typeTwo,
        dynasty: this.typeThree,
      }).then(res => {
        console.log(res.data)
        this.poemList = res.data.list
        const pageLength = Math.ceil(res.data.total / this.page.pageSize)
        this.page.pageLength = pageLength <= 1000 ? pageLength : 1000
      })
    },
    async onPageChange(page) {
      this.page.currentPage = page
      await this.getPoemListByPage()
    },
  },
  setup() {
    const poemShi = {
      chipGroupTypeOne: ['悼亡', '离别', '山水', '田园'],
      chipGroupTypeTwo: ['乐府', '七律', '五律', '七绝', '五绝'],
      chipGroupTypeThree: ['唐', '宋', '元', '明', '清', '五代十国', '汉', '秦', '隋', '三国', '南北朝', '金', '晋'],
    }
    const poemCi = {
      chipGroupTypeOne: ['悼亡', '离别', '山水', '田园'],
      chipGroupTypeTwo: ['水调歌头', '浣溪沙', '临江仙', '蝶恋花', '西江月', '满江红'],
      chipGroupTypeThree: ['唐', '宋', '元', '明', '清', '五代十国', '隋', '南北朝', '金'],
    }
    const poemQu = {
      chipGroupTypeOne: ['悼亡', '离别', '山水', '田园'],
      chipGroupTypeTwo: [],
      chipGroupTypeThree: ['元', '明', '清'],
    }

    return {
      poemShi,
      poemCi,
      poemQu,
    }
  },
}
</script>
