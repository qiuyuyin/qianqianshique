<template>
  <v-navigation-drawer
    :value="isDrawerOpen"
    app
    floating
    width="220"
    class="app-navigation-menu"
    :right="$vuetify.rtl"
    @input="val => $emit('update:is-drawer-open', val)"
  >
    <!-- Navigation Header -->
    <div class="vertical-nav-header d-flex items-center ps-6 pe-5 pt-5 pb-2">
      <router-link
        to="/"
        class="d-flex align-center text-decoration-none"
      >
        <v-img
          :src="require('@/assets/images/logos/logo.svg')"
          max-height="30px"
          max-width="30px"
          alt="logo"
          contain
          eager
          class="app-logo me-3"
        ></v-img>
        <v-slide-x-transition>
          <h2 class="app-title text--primary">
            千千诗阙
          </h2>
        </v-slide-x-transition>
      </router-link>
    </div>

    <!-- Navigation Items -->
    <v-list
      expand
      rounded
      class="vertical-nav-menu-items pr-5"
    >
      <nav-menu-link
        title="主页"
        :to="{ name: 'dashboard' }"
        :icon="icons.mdiHomeOutline"
      ></nav-menu-link>
      <nav-menu-link
        title="作者"
        :to="{ name: 'author'}"
        :icon="icons.mdiAccountHardHat"
      ></nav-menu-link>
      <nav-menu-link
        title="诗句"
        :to="{ name: 'poem'}"
        :icon="icons.mdiScriptTextOutline"
      ></nav-menu-link>
      <nav-menu-link
        title="梁园"
        :to="{ name: 'ai-poem-shi'}"
        :icon="icons.mdiAccountCogOutline"
      ></nav-menu-link>
    </v-list>
    <v-divider class="mb-10"></v-divider>

    <div class="pt-15">
    </div>
    <v-footer
      absolute
      class="font-weight-medium"
      padless
      color="transparent"
    >
      <v-col
        class="text-center"
        cols="12"
      >
        <strong>切换主题：</strong><theme-switcher class="pb-1"></theme-switcher>
      </v-col>
    </v-footer>
  </v-navigation-drawer>
</template>

<script>
// eslint-disable-next-line object-curly-newline
import {
  mdiHomeOutline,
  mdiAlphaTBoxOutline,
  mdiEyeOutline,
  mdiCreditCardOutline,
  mdiTable,
  mdiFileOutline,
  mdiFormSelect,
  mdiAccountCogOutline,
  mdiAccountHardHat,
  mdiScriptTextOutline,
} from '@mdi/js'
import NavMenuLink from './components/NavMenuLink.vue'
import ThemeSwitcher from '../ThemeSwitcher.vue'

export default {
  components: {
    NavMenuLink,
    ThemeSwitcher,
  },
  props: {
    isDrawerOpen: {
      type: Boolean,
      default: null,
    },
  },
  setup() {
    return {
      icons: {
        mdiHomeOutline,
        mdiAlphaTBoxOutline,
        mdiEyeOutline,
        mdiCreditCardOutline,
        mdiTable,
        mdiFileOutline,
        mdiFormSelect,
        mdiAccountCogOutline,
        mdiAccountHardHat,
        mdiScriptTextOutline,
      },
    }
  },
}
</script>

<style lang="scss" scoped>
.app-title {
  font-size: 1.25rem;
  font-weight: 700;
  font-stretch: normal;
  font-style: normal;
  line-height: normal;
  letter-spacing: 0.3px;
}

// ? Adjust this `translateX` value to keep logo in center when vertical nav menu is collapsed (Value depends on your logo)
.app-logo {
  transition: all 0.18s ease-in-out;
  .v-navigation-drawer--mini-variant & {
    transform: translateX(-4px);
  }
}

@include theme(app-navigation-menu) using ($material) {
  background-color: map-deep-get($material, 'background');
}

.app-navigation-menu {
  .v-list-item {
    &.vertical-nav-menu-link {
      ::v-deep .v-list-item__icon {
        .v-icon {
          transition: none !important;
        }
      }
    }
  }
}

// You can remove below style
// Upgrade Banner
.app-navigation-menu {
  .upgrade-banner {
    position: absolute;
    bottom: 13px;
    left: 50%;
    transform: translateX(-50%);
  }
}
</style>
