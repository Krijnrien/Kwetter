import Vue from 'vue';
import Vuex from 'vuex';
import axios from 'axios';
import VueNativeSock from 'vue-native-websocket';

const BACKEND_URL = 'http://192.168.99.100:8080';
const PUSHER_URL = 'ws://192.168.99.100:8080/pusher';

const SET_KWEETS = 'SET_KWEETS';
const CREATE_KWEET = 'CREATE_KWEET';
const SEARCH_SUCCESS = 'SEARCH_SUCCESS';
const SEARCH_ERROR = 'SEARCH_ERROR';

const MESSAGE_MEOW_CREATED = 1;

Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    kweets: [],
    searchResults: [],
  },
  mutations: {
    SOCKET_ONOPEN(state, event) {},
    SOCKET_ONCLOSE(state, event) {},
    SOCKET_ONERROR(state, event) {
      console.error(event);
    },
    SOCKET_ONMESSAGE(state, message) {
      switch (message.kind) {
        case MESSAGE_MEOW_CREATED:
          this.commit(CREATE_KWEET, { id: message.id, body: message.body });
      }
    },
    [SET_KWEETS](state, kweets) {
      state.kweets = kweets;
    },
    [CREATE_KWEET](state, kweet) {
      state.kweets = [kweet, ...state.kweets];
    },
    [SEARCH_SUCCESS](state, kweets) {
      state.searchResults = kweets;
    },
    [SEARCH_ERROR](state) {
      state.searchResults = [];
    },
  },
  actions: {
    getKweets({ commit }) {
      axios
        .get(`${BACKEND_URL}/poster`)
        .then(({ data }) => {
          commit(SET_KWEETS, data);
        })
        .catch((err) => console.error(err));
    },
    async createKweet({ commit }, kweet) {
      const { data } = await axios.post(`${BACKEND_URL}/poster`, null, {
        params: {
          body: kweet.body,
        },
      });
    },
    async searchKweets({ commit }, query) {
      if (query.length == 0) {
        commit(SEARCH_SUCCESS, []);
        return;
      }
      axios
        .get(`${BACKEND_URL}/search`, {
          params: { query },
        })
        .then(({ data }) => commit(SEARCH_SUCCESS, data))
        .catch((err) => {
          console.error(err);
          commit(SEARCH_ERROR);
        });
    },
  },
});

Vue.use(VueNativeSock, PUSHER_URL, { store, format: 'json' });

store.dispatch('getKweets');

export default store;
