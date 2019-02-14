<template>
  <v-timeline align-top class="ma-2">
    <v-timeline-item
      v-for="event in events"
      :color="event.color"
      :icon="event.icon"
      :key="event.id"
      fill-dot
    >
      <span slot="opposite" :class="`headline font-weight-bold text`">
        <font :color="event.color">{{ formatDate(event.created_at) }}</font>
        <p class="grey--text">{{ event.author }}</p>
      </span>
      <v-card :color="event.color" dark>
        <v-card-title class="title">
          {{ event.title }}
          <v-spacer />
          <v-btn class="mx-0 edit-btn" flat icon @click="edit(event)">
            <v-icon small>fas fa-pen</v-icon>
          </v-btn>
          <v-btn class="mx-0 edit-btn" flat icon @click="deleteEvent(event)">
            <v-icon small>fas fa-trash-alt</v-icon>
          </v-btn>
        </v-card-title>
        <v-card-text class="white text--primary" row>
          <span style="white-space: pre-line;">{{ event.message }}</span>
        </v-card-text>
      </v-card>
    </v-timeline-item>
  </v-timeline>
</template>

<script>
import format from "date-fns/format";

export default {
  created() {
    this.$store.dispatch("fetchData");
  },
  methods: {
    edit(event) {
      this.$store.commit("updateEvent", JSON.parse(JSON.stringify(event)));
      this.$store.commit("updateIsNew", false);
      this.$store.commit("updateEditDialog", true);
    },
    deleteEvent(event) {
      this.$store.commit("updateEvent", JSON.parse(JSON.stringify(event)));
      this.$store.commit("updateIsNew", false);
      this.$store.commit("updateDeleteDialog", true);
    },
    formatDate(created_at) {
      return format(created_at, "YYYY-MM-DD HH:MM:SS");
    }
  },
  computed: {
    events() {
      return this.$store.getters.events;
    }
  }
};
</script>
