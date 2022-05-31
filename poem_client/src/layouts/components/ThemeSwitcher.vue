<template>
  <v-tooltip
    bottom
    color="primary"
  >
    <template v-slot:activator="{ on, attrs }">
      <theme-switcher
        v-bind="attrs"
        v-on="on"
      >
        <v-fade-transition
          mode="out-in"
        >
          <v-icon
            :key="$vuetify.theme.dark"
            @click="changeTheme"
          >
            {{ $vuetify.theme.dark ? icons.mdiWeatherSunny : icons.mdiWeatherNight }}
          </v-icon>
        </v-fade-transition>
      </theme-switcher>
    </template>
    <span>主题切换</span>
  </v-tooltip>
</template>

<script>
import { mdiWeatherNight, mdiWeatherSunny } from '@mdi/js'
import { mapActions } from 'vuex'

export default {
  methods: {
    ...mapActions('user', ['SetDark']),
    changeTheme() {
      this.$vuetify.theme.dark = !this.$vuetify.theme.dark
      this.SetDark(this.$vuetify.theme.dark)
    },
  },
  setup() {
    return {
      icons: {
        mdiWeatherNight,
        mdiWeatherSunny,
      },
    }
  },
}
</script>

<style>
</style>
