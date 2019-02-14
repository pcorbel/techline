<template>
  <v-layout row justify-center>
    <v-dialog
      :value="editDialog"
      @input="updateEditDialog"
      max-width="600px"
      persistent
    >
      <v-btn slot="activator" color="primary" dark>Add an event</v-btn>
      <v-card>
        <v-card-title>
          <span v-if="isNew" class="headline">New event</span>
          <span v-if="!isNew" class="headline">Edit event</span>
        </v-card-title>
        <v-card-text>
          <v-form ref="form">
            <v-container grid-list-md>
              <v-layout row wrap>
                <v-flex xs12>
                  <v-select
                    label="Title*"
                    :value="event.title"
                    :items="titles"
                    :rules="inputRules"
                    @input="updateTitle"
                  ></v-select>
                </v-flex>
                <v-flex xs12>
                  <v-text-field
                    label="Author*"
                    :value="event.author"
                    :rules="inputRules"
                    @input="updateAuthor"
                  ></v-text-field>
                </v-flex>
                <v-flex xs12>
                  <v-textarea
                    label="Message*"
                    :value="event.message"
                    :rules="inputRules"
                    @input="updateMessage"
                  ></v-textarea>
                </v-flex>
              </v-layout>
            </v-container>
            <small>*required field</small>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" flat @click="reset">Cancel</v-btn>
          <v-btn color="blue darken-1" flat @click="saveEvent">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>

<script>
import { mapState } from "vuex";

export default {
  methods: {
    updateEditDialog(editDialog) {
      this.$store.commit("updateEditDialog", editDialog);
    },
    updateIsNew(isNew) {
      this.$store.commit("updateIsNew", isNew);
    },
    updateEvent(event) {
      this.$store.commit("updateEvent", event);
    },
    updateAuthor(author) {
      this.event.author = author;
      this.updateEvent(this.event);
    },
    updateTitle(title) {
      let theme = this.themes.find(theme => theme.title === title);
      this.event.title = title;
      this.event.color = theme.color;
      this.event.icon = theme.icon;
      this.updateEvent(this.event);
    },
    updateMessage(message) {
      this.event.message = message;
      this.updateEvent(this.event);
    },
    reset() {
      this.$store.commit("reset");
    },
    async saveEvent() {
      if (this.$refs.form.validate()) {
        if (this.isNew) {
          let date = new Date();
          date.setMilliseconds(0);
          this.event.created_at = date.toISOString();
          this.updateEvent(this.event);
        }
        await this.$store.dispatch("saveEvent");
        await this.$store.dispatch("fetchData");
        this.reset();
      }
    }
  },
  computed: {
    ...mapState({
      editDialog: state => state.editDialog,
      isNew: state => state.isNew,
      themes: state => state.themes,
      event: state => state.event,
      titles: state =>
        state.themes.map(theme => {
          return theme.title;
        })
    })
  },
  data() {
    return {
      inputRules: [v => !!v || "This field is required"]
    };
  }
};
</script>
