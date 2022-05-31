<template>
  <v-app>
    <vertical-nav-menu
      v-if="!upAndMd"
      :is-drawer-open.sync="isDrawerOpen"
    ></vertical-nav-menu>

    <v-app-bar
      app
      flat
      absolute
      color="transparent"
    >
      <div class="boxed-container w-full">
        <div class="d-flex align-center mx-3">
          <!-- Left Content -->
          <router-link
            v-if="$vuetify.breakpoint.mdAndUp"
            to="/"
          >
            <v-img
              :src="require('@/assets/images/logos/logo.svg')"
              max-height="40px"
              max-width="40px"
              alt="logo"
              contain
              eager
              class="app-logo me-3"
            ></v-img>
          </router-link>
          <router-link
            v-if="$vuetify.breakpoint.mdAndUp"
            to="/"
          >
            <v-app-bar-title class="text--primary mr-5 ml-n1 font-weight-bold">
              千千诗阙
            </v-app-bar-title>
          </router-link>

          <v-app-bar-nav-icon
            class="d-block d-md-none ml-n2"
            @click="isDrawerOpen = !isDrawerOpen"
          ></v-app-bar-nav-icon>
          <v-tabs
            v-if="$vuetify.breakpoint.mdAndUp"
            background-color="transparent"
            class="app-bar-tabs"
          >
            <v-tab
              :to="{ name: 'dashboard'}"
              class="font-weight-semibold"
            >
              主页
            </v-tab>
            <v-tab
              :to="{ name: 'poem'}"
              class="font-weight-semibold"
            >
              诗句
            </v-tab>
            <v-tab
              class="font-weight-semibold"
              :to="{ name: 'author'}"
            >
              作者
            </v-tab>
            <v-tab
              class="font-weight-semibold"
              :to="{ name: 'sentence'}"
            >
              名句
            </v-tab>
            <v-tab
              class="font-weight-semibold"
              :to="{ name: 'ai-poem-shi'}"
            >
              梁园
            </v-tab>
          </v-tabs>
          <!-- Right Content -->
          <app-bar-search-form></app-bar-search-form>
          <template v-if="$vuetify.breakpoint.mdAndUp">
            <v-tooltip
              bottom
              color="primary"
            >
              <template v-slot:activator="{ on, attrs }">
                <a
                  href="https://github.com/qiuyuyin/qianqianshique"
                  target="_blank"
                  rel="nofollow"
                  v-bind="attrs"
                  v-on="on"
                >
                  <v-icon class="ms-6 me-4">
                    {{ icons.mdiGithub }}
                  </v-icon>
                </a>
              </template>
              <span>项目源码</span>
            </v-tooltip>
            <theme-switcher>
            </theme-switcher>
            <v-tooltip
              bottom
              color="primary"
            >
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  v-bind="attrs"
                  icon
                  small
                  class="ms-3"
                  :to="{name: 'about'}"
                  v-on="on"
                >
                  <v-icon>
                    {{ icons.mdiHelpCircleOutline }}
                  </v-icon>
                </v-btn>
              </template>
              <span>关于本站</span>
            </v-tooltip>
          </template>

          <app-bar-user-menu></app-bar-user-menu>
        </div>
      </div>
    </v-app-bar>

    <v-main>
      <v-container
        fluid
        class="pa-0"
      >
        <div class="app-content-container">
          <slot></slot>
        </div>
      </v-container>
    </v-main>
    <v-footer
      padless
      color="transparent"
    >
      <v-col
        class="text-center mt-2 mb-4"
        cols="12"
      >
        {{ new Date().getFullYear() }} — <strong>千千诗阙</strong><a
          class="ml-2"
          href="https://beian.miit.gov.cn/"
        >
          豫ICP备2021012626号-2</a>
      </v-col>
    </v-footer>
    <v-bottom-navigation
      v-if="!$vuetify.breakpoint.mdAndUp"
      v-model="bottomNav"
      color="primary"
      fixed
      grow
      mandatory
      max-height="38px"
    >
      <v-btn
        text
        to="/"
      >
        <span>主页</span>
        <v-icon>
          {{ icons.mdiHomeCircleOutline }}
        </v-icon>
      </v-btn>

      <v-btn
        text
        to="/poem"
      >
        <span>诗句</span>
        <v-icon>{{ icons.mdiScriptTextOutline }}</v-icon>
      </v-btn>

      <v-btn
        text
        :to="{ name: 'author'}"
      >
        <span>作者</span>
        <v-icon>{{ icons.mdiAccountHardHat }}</v-icon>
      </v-btn>

      <v-btn
        text
        :to="{ name: 'ai-poem-shi'}"
      >
        <span>梁园</span>
        <v-icon>{{ icons.mdiLeadPencil }}</v-icon>
      </v-btn>
    </v-bottom-navigation>
    <p></p>
  </v-app>
</template>

<script>
import { ref } from '@vue/composition-api'
import {
  mdiMagnify, mdiBellOutline, mdiGithub, mdiAccountHardHat,
  mdiHomeCircleOutline, mdiHomeAssistant, mdiLeadPencil,
  mdiScriptTextOutline, mdiHelpCircleOutline,
} from '@mdi/js'
import VerticalNavMenu from './components/vertical-nav-menu/VerticalNavMenu.vue'
import ThemeSwitcher from './components/ThemeSwitcher.vue'
import AppBarUserMenu from './components/AppBarUserMenu.vue'
import AppBarSearchForm from './components/AppBarSearchForm.vue'

export default {
  components: {
    VerticalNavMenu,
    ThemeSwitcher,
    AppBarUserMenu,
    AppBarSearchForm,
  },
  data() {
    return {
      bottomNav: 0,
    }
  },
  computed: {
    upAndMd() {
      return this.$vuetify.breakpoint.mdAndUp
    },
  },
  setup() {
    const isDrawerOpen = ref(null)

    return {
      isDrawerOpen,

      // Icons
      icons: {
        mdiMagnify,
        mdiBellOutline,
        mdiGithub,
        mdiHomeCircleOutline,
        mdiHomeAssistant,
        mdiLeadPencil,
        mdiAccountHardHat,
        mdiScriptTextOutline,
        mdiHelpCircleOutline,
      },
    }
  },
}
</script>

<style lang="scss" scoped>
.v-app-bar ::v-deep {
  .v-toolbar__content {
    padding: 0;

    .app-bar-search {
      .v-input__slot {
        padding-left: 18px;
      }
    }

    .app-bar-tabs {
      box-shadow: 0 0px 0px !important;
      width: 60%;

      .v-tab {
        min-width: 75px;
        max-width: 80px;
        font-size: 16px;
      }
    }

  }
}
.app-content-container{
  width: 88%;
  max-width: 1230px;
  margin-left: auto;
  margin-right: auto;
}

.boxed-container {
  max-width: 1300px;
  margin-left: auto;
  margin-right: auto;
}
</style>
