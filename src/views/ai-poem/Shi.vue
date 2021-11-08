<template>
  <div>
    <v-card>
      <v-container fluid>
        <v-row>
          <v-col
            cols="12"
            md="7"
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
                  class="mb-n5"
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
                class="pb-0"
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
            md="5"
            class="mt-3   pb-0"
          >
            <v-text-field
              v-model="keys"
              clearable
              outlined
              type="text"
              label=""
              dense
              class=""
              :rules="rules"
            >
              <template v-slot:append-outer>
                <v-btn
                  color="primary"
                  bottom
                  class="mt-n2 ml-8"
                  :disabled="disabled"
                  @click="getPoem"
                >
                  生成
                </v-btn>
              </template>
            </v-text-field>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
    <div v-if="poems.length !== 0">
      <v-row class="mt-4 justify-center">
        <v-col
          cols="12"
          md="5"
          class="align-center"
        >
          <v-card
            height="470px"
            class="text-center pt-5 align-center"
          >
            <div v-if="!progressOn">
              <h1 class="text-xl font-weight-medium mb-3 mt-15">
                《{{ copyKeys }}》
              </h1>
              <h2
                v-for="poem in copyPoems"
                :key="poem"
                class="poem-text "
              >
                {{ poem }}
              </h2>
            </div>
            <div v-if="progressOn">
              <h2
                color="primary"
                class="mt-12 primary-text"
              >
                请等待
              </h2>
              <v-progress-circular
                indeterminate
                color="primary"
                :size="60"
                class="mt-15 pt-15"
              ></v-progress-circular>
            </div>
          </v-card>
        </v-col>
      </v-row>
    </div>
  </div>
</template>

<script>
import { getTokenOfPoem, getAIPoem } from '@/api/getAIPoem'

export default {
  setup() {
    return {
      icons: {
      },
    }
  },
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
      keys: '风雪漫天',
      poems: [],
      type: '1',
      copyPoems: [],
      copyKeys: '',
      progressOn: false,
    }
  },
  computed: {
    disabled() {
      return false
    },
  },
  mounted() {
    this.getToken()
  },
  methods: {
    getToken() {
      getTokenOfPoem().then(res => {
        console.log(res)
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
        console.log(res)
        this.poems = res.data.data.poems
        console.log(this.poems)
        this.copyKeys = this.keys
        this.copyPoems = this.poems
      })
      setTimeout(() => { this.progressOn = false }, 1000)

      // this.progressOn = false
    },
  },
}
</script>

<style>
.poem-text {
  letter-spacing:10px;
}
</style>
