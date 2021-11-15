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
        md="6"
      >
        <dashboard-recommand-poem-card
          :id="data.id"
          :dynasty="data.dynasty"
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
        dynasty: 'song',
        title: '丑奴儿·书博山道中壁',
        author: '辛弃疾',
        content: '少年不识愁滋味，爱上层楼。爱上层楼，为赋新词强说愁。而今识尽愁滋味，欲说还休。欲说还休，却道“天凉好个秋”！',
      },
      {
        id: 2,
        dynasty: 'tang',
        title: '宣州谢脁楼饯别校书叔云',
        author: '李白',
        content: '弃我去者，昨日之日不可留；乱我心者，今日之日多烦忧。长风万里送秋雁，对此可以酣高楼。蓬莱文章建安骨，中间小谢又清发。俱怀逸兴壮思飞，欲上青天揽明月。抽刀断水水更流，举杯销愁愁更愁。人生在世不称意，明朝散发弄扁舟。',
      },
      {
        id: 3,
        dynasty: 'qing',
        title: '木兰花·拟古决绝词柬友',
        author: '纳兰性德',
        content: '人生若只如初见，何事秋风悲画扇。等闲变却故人心，却道故人心易变。骊山语罢清宵半，泪雨霖铃终不怨。何如薄幸锦衣郎，比翼连枝当日愿。',
      },
      {
        id: 4,
        dynasty: 'qing',
        title: '浣溪沙',
        author: '纳兰性德',
        content: '谁念西风独自凉，萧萧黄叶闭疏窗，沉思往事立残阳。被酒莫惊春睡重，赌书消得泼茶香，当时只道是寻常。',
      },
      {
        id: 5,
        dynasty: 'song',
        title: '如梦令',
        author: '李清照',
        content: '昨夜雨疏风骤，浓睡不消残酒。试问卷帘人，却道海棠依旧。知否，知否？应是绿肥红瘦。',
      },
      {
        id: 6,
        dynasty: 'tang',
        title: '梦微之',
        author: '白居易',
        content: '夜来携手梦同游，晨起盈巾泪莫收。漳浦老身三度病，咸阳宿草八回秋。君埋泉下泥销骨，我寄人间雪满头。阿卫韩郎相次去，夜台茫昧得知不？',
      },

    ]

    return {
      recommandPoem,
    }
  },
}
</script>
