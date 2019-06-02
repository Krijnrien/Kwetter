<template>
  <div>
    <form v-on:submit.prevent="createKweet">
      <div class="input-group">
        <input v-model.trim="kweetBody" type="text" class="form-control" placeholder="Supperino">
        <div class="input-group-append">
          <button class="btn btn-primary" type="submit">Kweet</button>
        </div>
      </div>
    </form>

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
        kweetBody: '',
      };
    },
    computed: mapState({
      kweets: (state) => state.kweets,
    }),
    methods: {
      createKweet() {
        if (this.kweetBody.length != 0) {
          this.$store.dispatch('createKweet', { body: this.kweetBody });
          this.kweetBody = '';
        }
      },
    },
    components: {
      Kweet,
    },
  };
</script>
