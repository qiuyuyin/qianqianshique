<template>
  <div>
    <p
      class="text-xl mt-2 primary--text font-weight-semibold"
    >
      Questions
    </p>
    <v-expansion-panels
      focusable
    >
      <v-expansion-panel>
        <v-expansion-panel-header class="primary--text">
          <h3>1. 为什么要做一个诗词类的网站？</h3>
        </v-expansion-panel-header>
        <v-expansion-panel-content class="text-primary">
          <p>
            <strong>浪花有意千重雪，桃李无言一队春。一壶酒，一竿纶，快活如侬有几人。
              一棹春风一叶舟，一纶茧缕一轻钩。花满渚，酒满瓯，万顷波中得自由。</strong>
          </p>

          中国诗词文化渊源流长，但是随之短视频和泛娱乐的兴起，大多数人都被裹挟在互联网的洪流中，无法脱身。
          设计本站的一个想法就是希望我们中华的诗词不仅仅出现在课本和古籍中，也可以被记录在互联网中
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="primary--text">
          <h3>2. 本项目是否完全开源？</h3>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <p>
            本项目保证<strong class="primary--text">完全开源</strong>，因为诗词是没有版权限制的（作者去世50年后就会失去著作权）。
          </p>
          <p>
            所以不用担心任何版权问题，可以放心的进行传播，之后本站也会发布有关部署相关的操作。
          </p>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="primary--text">
          <h3>3. 项目是否会一直维护？</h3>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          本站现在只是初期阶段，很多有趣的功能还没有加入，所以请耐心等待下一个版本的迭代。
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="primary--text">
          <h3>4. 是否可以用作商业用途？</h3>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <p>本项目使用<strong class="primary--text">MIT开源协议</strong>，请勿用作商用，或者其他的牟利用途</p>
          <p>我们应该尊重中华传统文化，而不是让他成为牟利的工具</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header class="primary--text">
          <h3>5. 以后会考虑移动端吗？</h3>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
          <p><strong>本网站已经支持移动浏览器进行访问。</strong></p>
          <p>由于开发本站的还是学生，时间还不是很充裕，但是本站是会考虑使用Uni-App技术搭建小程序模块，以方便用户对诗词进行浏览。</p>
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
    <p
      class="text-2xl primary--text pt-10 font-weight-semibold text-center"
    >
      2021年项目进度
    </p>
    <v-timeline>
      <v-timeline-item
        v-for="(year, i) in years"
        :key="i"
        color="secondary"
        small
      >
        <template v-slot:opposite>
          <span
            :class="`headline font-weight-bold primary--text`"
            :style="{ color:'pink' }"
            v-text="year.year"
          ></span>
        </template>
        <v-card class="pa-2">
          <h4 class="font-weight-blod pl-2 primary--text">
            {{ year.title }}
          </h4>
          <v-card-text
            v-if="$vuetify.breakpoint.mdAndUp"
            class="mt-2"
          >
            {{ year.des }}
          </v-card-text>
        </v-card>
      </v-timeline-item>
    </v-timeline>
  </div>
</template>
<script>
export default {
  data: () => ({
    years: [
      {
        color: 'cyan',
        year: '12-19',
        title: '关于本站',
        des: '为本站提供了一些问答模式，和时间树组件',
      },
      {
        color: 'green',
        year: '12-14',
        title: 'Oss服务',
        des: '开始在后台搭建一些OSS（使用阿里云进行搭建），用户可以上传头像',
      },
      {
        color: 'pink',
        year: '12-11',
        title: '初步修复BUG',
        des: '修复了比如页面跳转，用户登录失败提示不清晰相关的一些BUG',
      },
      {
        color: 'cyan',
        year: '12-9',
        title: '1.0版本上线',
        des: '项目1.0版本上线，邀请了一些用户来进行了内测出现了一部分的BUG，特别是移动端的BUG',
      },
      {
        color: 'green',
        year: '12-6',
        title: '用户写诗',
        des: '用户可以通过输入一个主题来完成写诗的操作，这个诗句可以是藏头诗或者主题诗',
      },
      {
        color: 'pink',
        year: '12-9',
        title: '移动端适配',
        des: '由于项目总体使用flex来进行布局，所以支持响应式操作，但是需要特定的为移动端安装一些组件',
      },
      {
        color: 'cyan',
        year: '12-7',
        title: '写诗',
        des: '后台将用户写的诗进行保存到数据库中，这样所有的用户都会看见你写的诗',
      },
      {
        color: 'green',
        year: '12-6',
        title: '用户写诗',
        des: '用户可以通过输入一个主题来完成写诗的操作，这个诗句可以是藏头诗或者主题诗',
      },
      {
        color: 'pink',
        year: '12-3',
        title: '登录注册模块',
        des: '登录和注册模块都需要使用状态管理，在本项目中状态管理使用的是VUEX，通过vuex-persist插件来保存用户信息到本地，这下次进入网页会自动进行登录',
      },
      {
        color: 'cyan',
        year: '12-1',
        title: '用户数据库搭建',
        des: '由于本站需要考虑到和用户交互的情况所以需要搭建一个用户的数据库',
      },
      {
        color: 'green',
        year: '11-30',
        title: '名句前台',
        des: '由于前台开始熟练，项目速度开始突飞猛进',
      },
      {
        color: 'pink',
        year: '11-25',
        title: '名句后台',
        des: '名句需要和诗人和诗词进行配合，所以需要使对名句的数据进行清洗，之后再搭建名句的后台服务',
      },
      {
        color: 'amber',
        year: '11-20',
        title: '诗词前台',
        des: '采用vuetify组件开始搭建主要模块，包括诗人页面，诗句页面',
      },
      {
        color: 'orange',
        year: '11-16',
        title: '诗词后台',
        des: '构建诗词的后台服务，使用七牛云的MongoDB ORM',
      },
      {
        color: 'cyan',
        year: '11-10',
        title: '数据库构建',
        des: '搭建诗词数据库和用户数据库，分别使用Mysql和MongoDB',
      },
      {
        color: 'green',
        year: '11-4',
        title: '清洗数据',
        des: '由于数据由多个数据源中获取，并且获取的数据中存在诸多问题',
      },
      {
        color: 'pink',
        year: '10-25',
        title: '后台架构',
        des: '使用MVC架构进行搭建，各个层级之间互相不干扰',
      },
      {
        color: 'amber',
        year: '10-26',
        title: '前台demo',
        des: '通过vuetify搭建了一个前台的模型，便于之后开发',
      },
      {
        color: 'orange',
        year: '10-22',
        title: '前台架构',
        des: '选择前台所使用的架构，使用vue全家桶进行开发',
      },
    ],
  }),
}
</script>
