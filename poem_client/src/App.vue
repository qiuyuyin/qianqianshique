<template>
  <component :is="resolveLayout">
    <Dialog></Dialog>
    <router-view :key="$router.currentRoute.fullPath"></router-view>
  </component>
</template>

<script>
import { computed } from '@vue/composition-api'
import { useRouter } from '@/utils'
import LayoutBlank from '@/layouts/Blank.vue'
import LayoutContent from '@/layouts/Content.vue'
import UpgradeToPro from './components/UpgradeToPro.vue'
import Dialog from '@/components/Dialog.vue'

export default {
  components: {
    LayoutBlank,
    LayoutContent,
    UpgradeToPro,
    Dialog,
  },
  setup() {
    const { route } = useRouter()

    const resolveLayout = computed(() => {
      // Handles initial route
      if (route.value.name === null) return null

      if (route.value.meta.layout === 'blank') return 'layout-blank'

      return 'layout-content'
    })

    return {
      resolveLayout,
    }
  },
}
</script>
