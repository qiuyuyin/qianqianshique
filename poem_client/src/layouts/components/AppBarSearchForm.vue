<template>
  <v-menu
    offset-y
    :right="$vuetify.breakpoint.mdAndUp"
    nudge-bottom="10"
    max-width="300px"
    :close-on-content-click="close"
    class="mt-2"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-text-field
        ref="content"
        v-model="searchText"
        rounded
        dense
        outlined
        placeholder="全局搜索"
        :prepend-inner-icon="icons.mdiMagnify"
        class="app-bar-search flex-grow-0 ml-2 mr-1"
        hide-details
        v-bind="attrs"
        v-on="on"
      ></v-text-field>
    </template>
    <v-list>
      <template v-if="searchText == ''">
        <v-list-item v-if="$store.state.search.searchRecords.length != 0">
          <v-list-item-icon class="me-2">
            <v-icon size="20">
              {{ icons.mdiTextSearch }}
            </v-icon>
          </v-list-item-icon>
          <v-list-item-content class="secondary--text font-weight-black">
            搜索记录
          </v-list-item-content>
          <v-btn
            text
            small
            color="primary"
            @click="ClearRecord"
          >
            删除
          </v-btn>
        </v-list-item>
        <v-list-item v-else>
          <v-list-item-icon class="me-2">
            <v-icon size="20">
              {{ icons.mdiTextSearch }}
            </v-icon>
          </v-list-item-icon>
          <v-list-item-content class="secondary--text font-weight-black">
            暂无搜索记录
          </v-list-item-content>
        </v-list-item>
        <v-list-item
          v-for="(sentence,id) in $store.state.search.searchRecords"
          :key="id+10"
          @click="setText(sentence)"
        >
          <v-list-item-icon class="me-2">
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ sentence }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </template>
      <template v-else>
        <template v-if="code === 0">
          <template v-if="searchRes.SearchAuthorList">
            <v-list-item>
              <v-list-item-icon class="me-2">
                <v-icon size="20">
                  {{ icons.mdiAccountSearchOutline }}
                </v-icon>
              </v-list-item-icon>
              <v-list-item-content class="secondary--text font-weight-black">
                作者
              </v-list-item-content>
            </v-list-item>
            <v-list-item
              v-for="(author,i) in searchRes.SearchAuthorList"
              :key="i"
              @click="authorRouter(author.ID)"
            >
              <v-list-item-icon class="me-2">
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>{{ author.name }} · [{{ author.dynasty }}]</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-divider></v-divider>
          </template>
          <template v-if="searchRes.SearchPoemList">
            <v-list-item>
              <v-list-item-icon class="me-2">
                <v-icon size="20">
                  {{ icons.mdiTextBoxSearchOutline }}
                </v-icon>
              </v-list-item-icon>
              <v-list-item-content class="secondary--text font-weight-black">
                诗句
              </v-list-item-content>
            </v-list-item>
            <v-list-item
              v-for="poem in searchRes.SearchPoemList"
              :key="poem.id"
              @click="poemRouter(poem.ID,poem.poemType)"
            >
              <v-list-item-icon class="me-2">
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>{{ poem.title }} - [{{ poem.authorName }}]</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <v-divider></v-divider>
          </template>
          <template v-if="searchRes.SearchPoemList">
            <v-list-item>
              <v-list-item-icon class="me-2">
                <v-icon size="20">
                  {{ icons.mdiTextSearch }}
                </v-icon>
              </v-list-item-icon>
              <v-list-item-content class="secondary--text font-weight-black">
                诗句
              </v-list-item-content>
            </v-list-item>
            <v-list-item
              v-for="sentence in searchRes.SearchSentenceList"
              :key="sentence.id"
              link
            >
              <v-list-item-icon class="me-2">
              </v-list-item-icon>
              <v-list-item-content>
                <v-list-item-title>{{ sentence.content }}</v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </template>
        </template>
        <template v-else>
          <v-list-item>
            <v-list-item-icon class="me-2">
              <v-icon size="20">
                {{ icons.mdiTextSearch }}
              </v-icon>
            </v-list-item-icon>
            <v-list-item-content class="error--text font-weight-black">
              暂无信息
            </v-list-item-content>
          </v-list-item>
        </template>
      </template>
    </v-list>
  </v-menu>
</template>
<script>
import {
  mdiMagnify, mdiAccountSearchOutline, mdiStickerTextOutline, mdiTextSearch, mdiTextBoxSearchOutline,
} from '@mdi/js'
import { mapActions } from 'vuex'
import { searchAll } from '@/api/search'

export default {
  data() {
    return {
      searchText: '',
      searchRes: {},
      code: 0,
      close: false,
    }
  },
  watch: {
    // 如果 `question` 发生改变，这个函数就会运行
    searchText() {
      this.search()
      if (this.searchText === '') {
        this.close = false
      }
      setInterval(() => {
        if (this.searchText !== '') {
          this.close = true
        }
      }, 500)
    },
  },
  methods: {
    ...mapActions('search', ['ClearRecord']),
    search() {
      searchAll({ keyword: this.searchText })
        .then(res => {
          console.log(res)
          this.searchRes = res.data
          if (res.code === 0) {
            console.log(this.searchText)
            this.$store.dispatch('search/AppendRecord', this.searchText)
          }
          this.code = res.code
        })
    },
    setText(sentence) {
      this.searchText = sentence
      this.$refs.content.focus()
    },
    inoutFocus() {
      this.$refs.content.focus()
    },
    authorRouter(id) {
      this.inoutFocus()
      this.$router.push({ path: 'author/:id', name: 'author-detail', params: { id } })
    },
    poemRouter(id, poemType) {
      this.inoutFocus()
      this.$router.push({ path: 'poem/:id/:poemType', name: 'poem-detail', params: { id, poemType } })
    },
  },
  setup() {
    const icons = {
      mdiMagnify,
      mdiAccountSearchOutline,
      mdiStickerTextOutline,
      mdiTextSearch,
      mdiTextBoxSearchOutline,
    }

    return {
      icons,
    }
  },

}
</script>
