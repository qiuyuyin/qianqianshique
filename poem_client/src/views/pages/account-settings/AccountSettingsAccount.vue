<template>
  <v-card
    flat
    class="pa-3 mt-2"
  >
    <v-alert
      dense
      class="pb-3 ml-4 mr-4"
      outlined
      type="error"
    >
      请不要上传违法违规图片，后果自负
    </v-alert>
    <v-card-text class="d-flex">
      <v-avatar
        rounded
        size="120"
        class="me-6"
      >
        <v-img :src="imgHeader()"></v-img>
      </v-avatar>

      <!-- upload photo -->
      <div>
        <v-btn
          color="primary"
          class="me-3 mt-5"
          @click="$refs.refInputEl.click()"
        >
          <v-icon class="d-sm-none">
            {{ icons.mdiCloudUploadOutline }}
          </v-icon>
          <span class="d-none d-sm-block">修改头像</span>
        </v-btn>

        <input
          ref="refInputEl"
          type="file"
          accept=".jpeg,.png,.jpg,GIF"
          :hidden="true"
          @change="getFile"
        />

        <p class="text-sm mt-5">
          允许 JPG, PNG, SVG. 最大 1000K
        </p>
      </div>
    </v-card-text>

    <!-- <v-card-text>
      <v-form class="multi-col-validation mt-6">
        <v-row>
          <v-col
            md="6"
            cols="12"
          >
            <v-text-field
              v-model="accountDataLocale.name"
              label="Name"
              dense
              outlined
              class="pb-4"
            ></v-text-field>

            <v-text-field
              v-model="accountDataLocale.email"
              label="E-mail"
              dense
              class="pb-4"
              outlined
            ></v-text-field>

            <v-select
              v-model="accountDataLocale.status"
              dense
              outlined
              class="pb-4"
              label="Status"
              :items="status"
            ></v-select>
          </v-col>

          alert
          <v-col cols="12">
            <v-alert
              color="warning"
              text
              class="mb-0"
            >
              <div class="d-flex align-start">
                <v-icon color="warning">
                  {{ icons.mdiAlertOutline }}
                </v-icon>

                <div class="ms-3">
                  <p class="text-base font-weight-medium mb-1">
                    Your email is not confirmed. Please check your inbox.
                  </p>
                  <a
                    href="javascript:void(0)"
                    class="text-decoration-none warning--text"
                  >
                    <span class="text-sm">Resend Confirmation</span>
                  </a>
                </div>
              </div>
            </v-alert>
          </v-col>

          <v-col cols="12">
            <v-btn
              color="primary"
              class="me-3 mt-4"
            >
              保存修改
            </v-btn>
            <v-btn
              color="secondary"
              outlined
              class="mt-4"
              type="reset"
              @click.prevent="resetForm"
            >
              取消
            </v-btn>
          </v-col>
        </v-row>
      </v-form>
    </v-card-text> -->
  </v-card>
</template>

<script>
import { mdiAlertOutline, mdiCloudUploadOutline } from '@mdi/js'
import { mapActions } from 'vuex'
import uploadFile from '@/api/upload'
import getHeaderImg from '@/utils/headerImg'

export default {
  data() {
    return {
      file: {},
    }
  },
  methods: {
    ...mapActions('user', ['SetHeaderImg']),
    getFile(e) {
      const { files } = e.target
      if (files.length > 0) {
        const file = files[0]
        if (file.size > 1024 * 1024) {
          this.$store.dispatch('snackbar/openSnackbar', {
            msg: '文件太大,请重新上传',
            color: 'error',
          })
        } else {
          console.log(file)
          uploadFile(file, this.$store.state.user.userInfo.ID).then(res => {
            this.SetHeaderImg(res.data.data.url)
            console.log(res.data.data.url)
          })
        }
      }
    },
    imgHeader() {
      return getHeaderImg(this.$store.state.user.userInfo.userName, this.$store.state.user.userInfo.headerImg)
    },
  },

  setup() {
    const status = ['Active', 'Inactive', 'Pending', 'Closed']

    return {
      status,
      icons: {
        mdiAlertOutline,
        mdiCloudUploadOutline,
      },
    }
  },
}
</script>
