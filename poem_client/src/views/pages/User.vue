<template>
  <div id="account-setting-card">
    <p
      class="text-3xl primary--text font-weight-semibold"
    >
      {{ $store.state.user.userInfo.nickName }}的个人空间
    </p>
    <account-settings-account :account-data="accountSettingData.account"></account-settings-account>
  </div>
</template>

<script>

// demos
import AccountSettingsAccount from './account-settings/AccountSettingsAccount.vue'

export default {
  components: {
    AccountSettingsAccount,
  },
  setup() {
    const accountSettingData = {
      account: {
        avatarImg: require('@/assets/images/avatars/1.png'),
        username: 'johnDoe',
        name: 'john Doe',
        email: 'johnDoe@example.com',
        role: 'Admin',
        status: 'Active',
        company: 'Google.inc',
      },
      information: {
        bio: '欢迎来到千千诗阕',
        birthday: '2020-7-1',
        phone: '1234567',
        website: 'https://qianqianshique.com/',
        country: 'CN',
        languages: ['English', 'CN'],
        gender: 'male',
      },
    }

    return {
      accountSettingData,
    }
  },
}
</script>
