<template>
  <div>
    <v-dialog
      v-model="dialog"
      persistent
      max-width="600px"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          color="primary"
          v-bind="attrs"
          v-on="on"
        >
          我也要写诗
        </v-btn>
      </template>
      <v-card>
        <v-row>
          <v-col
            cols="12"
          >
            <v-row>
              <v-col
                cols="12"
                md="6"
                class="pb-0"
              >
                <v-radio-group
                  v-model="type"
                  row
                  class="mb-n5 ml-4"
                >
                  <template v-slot:label>
                    <div class="text-xl font-weight-semibold">
                      类型:
                    </div>
                  </template>
                  <v-radio
                    color="primary"
                    label="藏头诗"
                    value="1"
                  ></v-radio>
                  <v-radio
                    color="secondary"
                    label="主题诗"
                    value="2"
                  ></v-radio>
                </v-radio-group>
              </v-col>
              <v-col
                cols="12"
                md="6"
                class="ml-4"
              >
                <v-radio-group
                  v-model="poemYan"
                  row
                  class="mb-n2"
                >
                  <template v-slot:label>
                    <div class="text-xl font-weight-semibold">
                      绝句:
                    </div>
                  </template>
                  <v-radio
                    color="primary"
                    label="五言"
                    value="5"
                  ></v-radio>
                  <v-radio
                    color="secondary"
                    label="七言"
                    value="7"
                  ></v-radio>
                </v-radio-group>
              </v-col>
            </v-row>
          </v-col>
          <v-col
            cols="12"
            class=" pb-0"
          >
            <v-text-field
              v-model="keys"
              clearable
              outlined
              type="text"
              label=""
              dense
              class="ml-4"
              :rules="rules"
            >
              <template v-slot:append-outer>
                <v-btn
                  color="primary"
                  bottom
                  class="mt-n2 mr-3"
                  :disabled="disabled"
                  @click="getPoem"
                >
                  生成
                </v-btn>
              </template>
            </v-text-field>
          </v-col>
        </v-row>
        <div v-if="poems.length !== 0">
          <v-row class="mt-3 justify-center mb-3">
            <v-col
              cols="12"
              md="11"
            >
              <v-card
                class="text-center pt-4 pb-4"
              >
                <div v-if="!progressOn">
                  <h3 class="font-weight-bold mb-2">
                    《{{ copyKeys }}》
                  </h3>
                  <h3
                    v-for="poem in copyPoems"
                    :key="poem"
                    class="poem-text"
                  >
                    {{ poem }}
                  </h3>
                </div>
                <div v-if="progressOn">
                  <h2
                    color="primary"
                    class="primary-text"
                  >
                    请等待
                  </h2>
                  <v-progress-circular
                    indeterminate
                    color="primary"
                    :size="60"
                    class="mt-10 pt-15"
                  ></v-progress-circular>
                </div>
              </v-card>
            </v-col>
          </v-row>
        </div>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue darken-1"
            text
            @click="dialog = false"
          >
            关闭
          </v-btn>
          <v-btn
            color="blue darken-1"
            text
            @click="uploadPoem"
          >
            提交
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-row class="mt-3">
      <template v-for="poem in userPoemList">
        <v-col
          :key="poem.ID"
          cols="12"
          md="4"
        >
          <v-card
            max-width="400"
          >
            <v-card-title>
              <span class="title font-weight-light">{{ poem.title }}</span>
            </v-card-title>

            <v-card-text class="headline font-weight-bold">
              {{ poem.content }}
            </v-card-text>
            <v-card-actions>
              <v-list-item>
                <v-list-item-avatar color="grey darken-3 ml-n3">
                  <v-img
                    src="https://avataaars.io/?avatarStyle=Transparent&topType=ShortHairShortCurly&accessoriesType=Prescription02&hairColor=Black&facialHairType=Blank&clotheType=Hoodie&clotheColor=White&eyeType=Default&eyebrowType=DefaultNatural&mouthType=Default&skinColor=Light"
                  ></v-img>
                </v-list-item-avatar>

                <v-list-item-content>
                  <v-list-item-title>{{ poem.authorName }}</v-list-item-title>
                </v-list-item-content>

                <v-row
                  v-if="false"
                  align="center"
                  justify="end"
                >
                  <v-icon class="mr-1">
                    {{ mdiHeart }}
                  </v-icon>
                  <span class="subheading mr-2">256</span>
                  <span class="mr-1">·</span>
                  <v-icon class="mr-1">
                    {{ mdiShareVariant }}
                  </v-icon>
                  <span class="subheading">45</span>
                </v-row>
              </v-list-item>
            </v-card-actions>
          </v-card>
        </v-col>
      </template>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-pagination
          v-if="userPoemList"
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
import { mdiHeart, mdiShareVariant } from '@mdi/js'
import { getTokenOfPoem, getAIPoem } from '@/api/AIPoem'
import { postUserPoem, findUserPoemByPage } from '@/api/user'

export default {
  data() {
    return {
      rules: [
        value => !!value || '最少输入一个字',
        value => (value && value.length <= 4) || '不能大于四个字',
      ],
      everyDayDate: {
        content: '',
        origin: {
          author: '',
          dynasty: '',
        },
      },
      tokenData: {},
      poemYan: '5',
      keys: '交错进行',
      poems: [],
      type: '1',
      copyPoems: [],
      copyKeys: '',
      progressOn: false,
      dialog: false,
      userPoemList: [], // 得到的userPoem集合
      page: {
        currentPage: 1, // 当前页面
        pageSize: 6, // 页面大小
        pageLength: 3, // 设置的显示分页数目
        pageVisible: 10, // 可视化的分页数目
      },
    }
  },
  computed: {
    disabled() {
      return false
    },
  },
  mounted() {
    this.getToken()
    this.getUserPoemListByPage()
  },
  methods: {
    uploadPoem() {
      postUserPoem({
        authorID: this.$store.state.user.userInfo.ID,
        authorName: this.$store.state.user.userInfo.nickName,
        title: this.copyKeys,
        content: this.poems[0],
      }).then(res => {
        console.log(res.data)
      })
      this.getUserPoemListByPage()
      this.dialog = false
    },
    getToken() {
      getTokenOfPoem().then(res => {
        this.tokenData = res.data
      })
    },
    getPoem() {
      if (!this.keys) return
      this.progressOn = true
      getAIPoem(this.tokenData.access_token, {
        keys: this.keys,
        poemtype: this.type,
        yan: this.poemYan,
      }).then(res => {
        this.poems = res.data.data.poems
        this.copyKeys = this.keys
        this.copyPoems[0] = `${this.poems[0]}，${this.poems[1]}。`
        this.copyPoems[1] = `${this.poems[2]}，${this.poems[3]}。`
      })
      setTimeout(() => { this.progressOn = false }, 1000)
    },
    getUserPoemListByPage() {
      findUserPoemByPage({
        page: this.page.currentPage,
        pageSize: this.page.pageSize,
      }).then(res => {
        console.log(res.data)
        this.userPoemList = res.data.list
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
    return {
      mdiHeart,
      mdiShareVariant,
    }
  },
}
</script>

<style>
.poem-text {
  letter-spacing:1px;
  font-weight: 440;
}
</style>
