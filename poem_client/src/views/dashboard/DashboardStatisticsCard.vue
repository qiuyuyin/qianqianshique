<template>
  <v-card min-height="180px">
    <v-card-title>
      <span class="font-weight-black text--secondary">本站所保存的诗词数据</span>
    </v-card-title>

    <v-card-subtitle class="mb-8 mt-n4">
      <span>中华文化，博大精深！</span>
    </v-card-subtitle>

    <v-card-text>
      <v-row>
        <v-col
          v-for="data in statisticsData"
          :key="data.title"
          cols="6"
          md="3"
          class="d-flex align-center"
        >
          <v-avatar
            size="44"
            :color="resolveStatisticsIconVariation (data.title).color"
            rounded
            class="elevation-1"
          >
            <v-icon
              dark
              color="white"
              size="30"
            >
              {{ resolveStatisticsIconVariation (data.title).icon }}
            </v-icon>
          </v-avatar>
          <div class="ms-3">
            <p class="text-xs mb-0">
              {{ data.title }}
            </p>
            <h3 class="text-xl font-weight-semibold">
              {{ data.total }}
            </h3>
          </div>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>

<script>
// eslint-disable-next-line object-curly-newline
import { mdiAccountOutline, mdiCurrencyUsd, mdiTrendingUp, mdiDotsVertical, mdiLabelOutline,
  mdiDeathStarVariant, mdiFileDocumentMultipleOutline,
} from '@mdi/js'

export default {
  setup() {
    const statisticsData = [
      {
        title: '诗人数目',
        total: '6926',
      },
      {
        title: '唐诗数目',
        total: '24万',
      },
      {
        title: '宋词数目',
        total: '2.7万',
      },
      {
        title: '名句数目',
        total: '3360',
      },
    ]

    const resolveStatisticsIconVariation = data => {
      if (data === '唐诗数目') return { icon: mdiTrendingUp, color: 'primary' }
      if (data === '诗人数目') return { icon: mdiAccountOutline, color: 'secondary' }
      if (data === '宋词数目') return { icon: mdiDeathStarVariant, color: 'warning' }
      if (data === '名句数目') return { icon: mdiFileDocumentMultipleOutline, color: 'info' }

      return { icon: mdiAccountOutline, color: 'success' }
    }

    return {
      statisticsData,
      resolveStatisticsIconVariation,

      // icons
      icons: {
        mdiDotsVertical,
        mdiTrendingUp,
        mdiAccountOutline,
        mdiLabelOutline,
        mdiCurrencyUsd,
        mdiDeathStarVariant,
        mdiFileDocumentMultipleOutline,
      },
    }
  },
}
</script>
