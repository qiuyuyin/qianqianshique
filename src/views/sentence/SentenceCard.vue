<template>
  <v-list
    v-bind="$attrs"
    class="align-self-start"
  >
    <v-list-item>
      <v-list-item-content>
        <v-list-item-title class="font-weight-bold text-xl mb-2 ml-2">
          {{ content }}
        </v-list-item-title>
        <v-list-item-subtitle class="text-right mr-2">
          ———
          <router-link
            :to="PoemRouter"
            class="sentence-datail-router-link"
          >
            <span>《{{ title }}》</span>
          </router-link>
          ·
          <router-link
            :to="AuthorRouter"
            class="sentence-datail-router-link"
          >
            <span :to="PoemRouter">{{ authorName }}</span>
          </router-link>
        </v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
  </v-list>
</template>
<script>

export default {
  props: {
    poemId: {
      type: String,
      required: true,
    },
    content: {
      type: String,
      required: true,
    },
    authorName: {
      type: String,
      required: true,
    },
    authorId: {
      type: String,
      required: true,
    },
    id: {
      type: String,
      required: true,
    },
    dynasty: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
    poemType: {
      type: String,
      required: true,
    },
  },
  computed: {
    AuthorRouter() {
      return {
        name: 'author-detail',
        params: {
          id: this.authorId,
        },
      }
    },
    PoemRouter() {
      return {
        name: 'poem-detail',
        params: {
          id: this.poemId,
          poemType: this.poemType,
        },
      }
    },
  },
}
</script>
<style type="text/css">
  .sentence-datail-router-link {
    text-decoration: none;
  }
  .sentence-datail-router-link:hover {
    cursor: pointer;
    text-decoration: underline;
    text-decoration-line: underline;
    text-decoration-thickness: initial;
    text-decoration-style: initial;
    text-decoration-color: initial;
}

</style>
