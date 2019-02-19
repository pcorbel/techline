import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import format from "date-fns/format";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    events: [],
    themes: [],
    editDialog: false,
    deleteDialog: false,
    isNew: true,
    filter: "",
    event: {
      id: null,
      title: "",
      color: "",
      icon: "",
      message: "",
      author: "",
      created_at: null
    }
  },
  getters: {
    // Apply filter
    events: state => {
      return state.events.filter(event => event.title.includes(state.filter));
    },
    // Format latest event createdAt timestamp
    lastEventCreatedAt: state => {
      return state.events.length === 0
        ? "1970-01-01 00 00 00"
        : format(
            state.events[0].created_at,
            "YYYY-MM-DD HH:mm:ss"
          );
    }
  },
  mutations: {
    updateEvents(state, events) {
      state.events = events;
    },
    updateThemes(state, themes) {
      state.themes = themes;
    },
    updateEditDialog(state, editDialog) {
      state.editDialog = editDialog;
    },
    updateDeleteDialog(state, deleteDialog) {
      state.deleteDialog = deleteDialog;
    },
    updateIsNew(state, isNew) {
      state.isNew = isNew;
    },
    updateEvent(state, event) {
      state.event = event;
    },
    updateFilter(state, filter) {
      state.filter = filter;
    },
    reset(state) {
      state.event = {};
      state.editDialog = false;
      state.deleteDialog = false;
      state.isNew = true;
    }
  },
  actions: {
    // Get Data from the API
    async fetchData() {
      let themes = await axios.get(
        process.env.VUE_APP_ROOT_API + process.env.VUE_APP_THEME_ENDPOINT
      );
      this.commit("updateThemes", themes.data.themes ? themes.data.themes : []);

      let events = await axios.get(
        process.env.VUE_APP_ROOT_API + process.env.VUE_APP_EVENT_ENDPOINT
      );
      this.commit("updateEvents", events.data ? events.data : []);
    },
    // Post Data to the API
    async saveEvent() {
      await axios.post(
        process.env.VUE_APP_ROOT_API + process.env.VUE_APP_EVENT_ENDPOINT,
        this.state.event,
        {
          headers: { "Content-Type": "application/json" }
        }
      );
    },
    // Delete Data to the API
    async deleteEvent() {
      await axios.delete(
        process.env.VUE_APP_ROOT_API +
          process.env.VUE_APP_EVENT_ENDPOINT +
          "/" +
          this.state.event.id
      );
    },
  }
});
