<template>
  <div>
    <p
      class="text-3xl  primary--text font-weight-semibold"
    >
      简介
    </p>
    <v-row>
      <v-col
        cols="12"
        md="4"
      >
        <dashboard-every-day-poem></dashboard-every-day-poem>
      </v-col>
      <v-col
        cols="12"
        md="8"
      >
        <dashboard-statistics-card></dashboard-statistics-card>
      </v-col>
    </v-row>
    <p
      class="text-3xl mt-6 primary--text font-weight-semibold"
    >
      推荐
    </p>
    <v-row>
      <v-col
        cols="12"
        md="4"
      >
        <dashboard-recommand-poem-card></dashboard-recommand-poem-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// eslint-disable-next-line object-curly-newline
import { mdiPoll, mdiLabelVariantOutline, mdiCurrencyUsd, mdiHelpCircleOutline } from '@mdi/js'

// demos
import DashboardEveryDayPoem from './DashboardEveryDayPoem.vue'
import DashboardStatisticsCard from './DashboardStatisticsCard.vue'
import DashboardRecommandPoemCard from './DashboardRecommandPoemCard.vue'

export default {
  components: {
    DashboardEveryDayPoem,
    DashboardStatisticsCard,
    DashboardRecommandPoemCard,
  },
  setup() {
    const totalProfit = {
      statTitle: 'Total Profit',
      icon: mdiPoll,
      color: 'success',
      subtitle: 'Weekly Project',
      statistics: '$25.6k',
      change: '+42%',
    }

    const totalSales = {
      statTitle: 'Refunds',
      icon: mdiCurrencyUsd,
      color: 'secondary',
      subtitle: 'Past Month',
      statistics: '$78',
      change: '-15%',
    }

    // vertical card options
    const newProject = {
      statTitle: 'New Project',
      icon: mdiLabelVariantOutline,
      color: 'primary',
      subtitle: 'Yearly Project',
      statistics: '862',
      change: '-18%',
    }

    const salesQueries = {
      statTitle: 'Sales Quries',
      icon: mdiHelpCircleOutline,
      color: 'warning',
      subtitle: 'Last week',
      statistics: '15',
      change: '-18%',
    }

    return {
      totalProfit,
      totalSales,
      newProject,
      salesQueries,
    }
  },
}
</script>
