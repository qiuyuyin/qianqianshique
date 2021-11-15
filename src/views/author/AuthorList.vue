<template>
  <div>
    <p
      class="text-3xl primary--text font-weight-semibold"
    >
      作者
    </p>
    <v-row>
      <v-col
        v-for="data in PoetList"
        :key="data.id"
        cols="12"
        md="8"
      >
        <v-hover
          v-slot="{ hover }"
          close-delay="50"
          open-delay="50"
        >
          <author-card
            :id="data.id"
            :elevation="hover ? 16 : 2"
            :class="{ 'on-hover': hover }"
            :dynasty="data.dynasty"
            :name="data.name"
            :content="data.content"
            :img-src="data.src"
          >
          </author-card>
        </v-hover>
      </v-col>
    </v-row>
  </div>
</template>

<script>
// eslint-disable-next-line object-curly-newline
import AuthorCard from './AuthorCard.vue'

export default {
  components: {
    AuthorCard,
  },
  setup() {
    const PoetList = [
      {
        id: 1,
        dynasty: '宋代',
        name: '辛弃疾',
        content: '少年不识愁滋味，爱上层楼。爱上层楼，为赋新词强说愁。而今识尽愁滋味，欲说还休。欲说还休，却道“天凉好个秋”！',
        src: 'https://song.gushiwen.cn/authorImg/libai.jpg',
      },
      {
        id: 2,
        dynasty: '唐代',
        name: '李白',
        content: '弃我去者，昨日之日不可留；乱我心者，今日之日多烦忧。长风万里送秋雁，对此可以酣高楼。蓬莱文章建安骨，中间小谢又清发。俱怀逸兴壮思飞，欲上青天揽明月。抽刀断水水更流，举杯销愁愁更愁。人生在世不称意，明朝散发弄扁舟。',
        src: 'https://song.gushiwen.cn/authorImg/libai.jpg',
      },
      {
        id: 3,
        dynasty: '清代',
        name: '纳兰性德',
        content: '人生若只如初见，何事秋风悲画扇。等闲变却故人心，却道故人心易变。骊山语罢清宵半，泪雨霖铃终不怨。何如薄幸锦衣郎，比翼连枝当日愿。',
        src: 'https://song.gushiwen.cn/authorImg/libai.jpg',
      },
      {
        id: 4,
        dynasty: '宋代',
        name: '李清照',
        content: '昨夜雨疏风骤，浓睡不消残酒。试问卷帘人，却道海棠依旧。知否，知否？应是绿肥红瘦。',
      },
      {
        id: 5,
        dynasty: '唐代',
        name: '白居易',
        content: '夜来携手梦同游，晨起盈巾泪莫收。漳浦老身三度病，咸阳宿草八回秋。君埋泉下泥销骨，我寄人间雪满头。阿卫韩郎相次去，夜台茫昧得知不？',
        src: 'https://song.gushiwen.cn/authorImg/libai.jpg',
      },

    ]

    return {
      PoetList,
    }
  },
}
</script>
