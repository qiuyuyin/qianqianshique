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
        v-for="data in recommandPoem"
        :key="data.id"
        cols="12"
        md="4"
      >
        <dashboard-recommand-poem-card
          :title="data.title"
          :content="data.content"
          :author="data.author"
          img-src="https://api.ixiaowai.cn/gqapi/gqapi.php"
        >
        </dashboard-recommand-poem-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// eslint-disable-next-line object-curly-newline

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
    const recommandPoem = [
      {
        id: 1,
        title: '不知道',
        author: '李白',
        content: 'asfasdfsad',
      },
      {
        id: 2,
        title: '不知道',
        author: '李白',
        content: 'asfasdfsad',
      },
      {
        id: 3,
        title: '不知道',
        author: '李白',
        content: 'asfasdfsad',
      },

    ]

    return {
      recommandPoem,
    }
  },
}
</script>
