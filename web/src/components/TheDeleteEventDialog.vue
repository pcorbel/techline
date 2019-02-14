<template>
  <v-layout row justify-center>
    <v-dialog
      :value="deleteDialog"
      @input="updateDeleteDialog"
      max-width="600px"
      persistent
    >
      <v-card>
        <v-card-title>
          <span class="headline">Delete event</span>
        </v-card-title>
        <v-card-text>
          Are you really sure?
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" flat @click="reset">Cancel</v-btn>
          <v-btn color="blue darken-1" flat @click="deleteEvent">Delete</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>

<script>
import { mapState } from "vuex";

export default {
  methods: {
    updateDeleteDialog(deleteDialog) {
      this.$store.commit("updateDeleteDialog", deleteDialog);
    },
    updateIsNew(isNew) {
      this.$store.commit("updateIsNew", isNew);
    },
    updateEvent(event) {
      this.$store.commit("updateEvent", event);
    },
    reset() {
      this.$store.commit("reset");
    },
    async deleteEvent() {
      await this.$store.dispatch("deleteEvent");
      await this.$store.dispatch("fetchData");
      this.reset();
    }
  },
  computed: {
    ...mapState({
      deleteDialog: state => state.deleteDialog,
      isNew: state => state.isNew,
      event: state => state.event
    })
  }
};
</script>
