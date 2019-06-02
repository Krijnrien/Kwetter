<template>
  <div>
    <input @keyup="searchKweets" v-model.trim="query" type="text" class="form-control" placeholder="Search...">
    <div class="mt-4">
      <Kweet v-for="kweet in kweets" :key="kweet.id" :kweet="kweet" />
    </div>
  </div>
</template>

<script>
  import { mapState } from 'vuex';
  import Kweet from '@/components/Kweet';

  export default {
    data() {
      return {
        query: '',
      };
    },
    computed: mapState({
      kweets: (state) => state.searchResults,
    }),
    methods: {
      searchKweets() {
        if (this.query != this.lastQuery) {
          this.$store.dispatch('searchKweets', this.query);
          this.lastQuery = this.query;
        }
      },
    },
    components: {
      Kweet,
    },
  };
</script>
