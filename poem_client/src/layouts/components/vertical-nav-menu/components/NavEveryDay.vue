<template>
  <div class="every-padding">
    <v-card color="text-center">
      <v-card-text class="d-flex flex-column justify-center align-center">
        <h6 class="text-xl mt-4 font-weight-medium">
          Support
        </h6>
      </v-card-text>
      <v-card-text>
        According to us blisters are a very common thing and we come across them very often in our daily lives. It is a
        very common occurrence like cold or fever depending upon your lifestyle.
      </v-card-text>
    </v-card>
  </div>
</template>
<script>
import { ref } from '@vue/composition-api'
import {
  mdiChevronUp,
  mdiChevronDown,
  mdiCartPlus,
  mdiShareVariantOutline,
  mdiLockOpenOutline,
  mdiStarOutline,
  mdiAccountOutline,
  mdiTrendingUp,
  mdiHelpCircleOutline,
} from '@mdi/js'

export default {
  setup() {
    const isCardDetailsVisible = false
    const rating = ref(5)

    return {
      isCardDetailsVisible,
      rating,

      // icons
      icons: {
        mdiChevronUp,
        mdiChevronDown,
        mdiCartPlus,
        mdiShareVariantOutline,
        mdiLockOpenOutline,
        mdiStarOutline,
        mdiAccountOutline,
        mdiTrendingUp,
        mdiHelpCircleOutline,
      },
    }
  },
}
</script>

<style lang="scss" scoped>

.every-padding{
  padding: 2vm;
  margin: 20px;
  bottom: 10px;
}

@import '@/plugins/vuetify/default-preset/preset/mixins.scss';

// membership pricing
@include theme--child(memberpricing-bg) using ($material) {
  background-color: rgba(map-deep-get($material, 'primary-shade'), map-deep-get($material, 'states', 'hover'));
}
</style>
