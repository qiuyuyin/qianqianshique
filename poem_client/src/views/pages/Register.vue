<template>
  <div class="auth-wrapper auth-v1">
    <div class="auth-inner">
      <v-card class="auth-card">
        <!-- logo -->
        <v-card-title class="d-flex align-center justify-center py-7">
          <router-link
            to="/"
            class="d-flex align-center"
          >
            <v-img
              :src="require('@/assets/images/logos/logo.svg')"
              max-height="30px"
              max-width="30px"
              alt="logo"
              contain
              class="me-3 "
            ></v-img>

            <h2 class="text-2xl font-weight-semibold">
              千千诗阙
            </h2>
          </router-link>
        </v-card-title>

        <!-- title -->
        <v-card-text>
          <p class="text-2xl font-weight-semibold text--primary mb-2">
            快速注册一个新账号 🚀
          </p>
          <p class="mb-2">
            恭喜你即将成为本站的第 {{ 1 }} 个用户
          </p>
        </v-card-text>

        <!-- login form -->
        <v-card-text>
          <v-form>
            <v-text-field
              v-model="username"
              outlined
              label="用户名"
              placeholder="用户名是唯一的哦，长度大于4"
              hide-details
              class="mb-3"
            ></v-text-field>

            <v-text-field
              v-model="nickname"
              outlined
              label="昵称"
              placeholder="填一个你喜欢的中文名，不唯一"
              hide-details
              class="mb-3"
            ></v-text-field>

            <v-text-field
              v-model="password"
              outlined
              :type="isPasswordVisible ? 'text' : 'password'"
              label="密码"
              placeholder="密码大于六位"
              :append-icon="isPasswordVisible ? icons.mdiEyeOffOutline : icons.mdiEyeOutline"
              hide-details
              @click:append="isPasswordVisible = !isPasswordVisible"
            ></v-text-field>
            <div class="d-flex ">
              <v-checkbox
                hide-details
                class="mt-1"
                v-model="checkbox"
              >
                <template #label>
                  <div class="d-flex align-center flex-wrap ">
                    <span class="me-2">我同意</span>
                  </div>
                </template>
              </v-checkbox>
              <v-dialog
                v-model="dialog"
                width="500"
              >
                <template v-slot:activator="{ on, attrs }">
                  <p
                    class="text-lg user-text primary--text "
                    v-bind="attrs"
                    v-on="on"
                  >
                    用户协议
                  </p>
                </template>

                <v-card>
                  <v-card-title
                    class="headline grey lighten-2"
                    primary-title
                  >
                    用户协议
                  </v-card-title>

                  <v-card-text>
                    <p>1.请不要在本站发布任何违反中华人民共和国法律相关的信息</p>
                    <p>2.请不要在本站中输入任何您的真实信息，如姓名，手机号，身份证ID</p>
                    <p>3.本网站为开源项目，请勿用作商务用途</p>
                    <p>4.如果网站中有任何版权相关问题，可联系邮箱qiansongyin@gmail.com</p>
                  </v-card-text>

                  <v-divider></v-divider>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                      color="primary"
                      text
                      @click="dialog = false"
                    >
                      我接受
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </div>

            <v-btn
              block
              color="primary"
              class="mt-6"
              @click="register"
            >
              注册
            </v-btn>
          </v-form>
        </v-card-text>

        <!-- create new account  -->
        <v-card-text class="d-flex align-center justify-center flex-wrap mt-2">
          <span class="me-2">
            已经有一个账号？
          </span>
          <router-link :to="{ name:'pages-login' }">
            登录
          </router-link>
        </v-card-text>

        <!-- divider -->
        <v-card-text class="d-flex align-center mt-2">
          <v-divider></v-divider>
        </v-card-text>
      </v-card>
    </div>

    <!-- background triangle shape  -->
    <img
      class="auth-mask-bg"
      height="190"
      :src="require(`@/assets/images/misc/mask-${$vuetify.theme.dark ? 'dark':'light'}.png`)"
    >

    <!-- tree -->
    <v-img
      class="auth-tree"
      width="247"
      height="185"
      src="@/assets/images/misc/tree.png"
    ></v-img>

    <!-- tree  -->
    <v-img
      class="auth-tree-3"
      width="377"
      height="289"
      src="@/assets/images/misc/tree-3.png"
    ></v-img>
  </div>
</template>

<script>
// eslint-disable-next-line object-curly-newline
import { mdiFacebook, mdiTwitter, mdiGithub, mdiGoogle, mdiEyeOutline, mdiEyeOffOutline } from '@mdi/js'
import { ref } from '@vue/composition-api'
import { mapActions } from 'vuex'
import { register } from '@/api/user'

export default {
  data() {
    return {
      dialog: false,
      checkbox: false,
    }
  },
  methods: {
    ...mapActions('user', ['LoginIn']),
    login() {
      this.LoginIn({
        username: this.username,
        password: this.password,
      })
    },
    register() {
      if (!this.checkbox) {
        this.$store.dispatch('snackbar/openSnackbar', {
          msg: '未同意用户协议',
          color: 'error',
        })
      } else {
        register({
          username: this.username,
          password: this.password,
          nickname: this.nickname,
        }).then(res => {
          console.log(res)
          if (res.code === -1) {
            this.$store.dispatch('snackbar/openSnackbar', {
              msg: res.msg,
              color: 'error',
            })
          } else {
            this.login()
          }
        })
      }
    },
  },
  setup() {
    const isPasswordVisible = ref(false)
    const username = ref('')
    const email = ref('')
    const password = ref('')
    const socialLink = [
      {
        icon: mdiFacebook,
        color: '#4267b2',
        colorInDark: '#4267b2',
      },
      {
        icon: mdiTwitter,
        color: '#1da1f2',
        colorInDark: '#1da1f2',
      },
      {
        icon: mdiGithub,
        color: '#272727',
        colorInDark: '#fff',
      },
      {
        icon: mdiGoogle,
        color: '#db4437',
        colorInDark: '#db4437',
      },
    ]

    return {
      isPasswordVisible,
      username,
      email,
      password,
      socialLink,

      icons: {
        mdiEyeOutline,
        mdiEyeOffOutline,
      },
    }
  },
}
</script>

<style lang="scss">
@import '~@/plugins/vuetify/default-preset/preset/pages/auth.scss';

.user-text {
  margin-top: 6px;
}
</style>
