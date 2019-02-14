<template>
  <v-toolbar color="primary" fixed dark app>
    <v-toolbar-title class="ml-0 pl-3">
      <v-icon left>far fa-calendar-check</v-icon>
      <span>Techline</span>
    </v-toolbar-title>
    <v-spacer />
    <div class="text-xs-center">
      <v-menu offset-x offset-y>
        <v-btn slot="activator" flat>
          <v-icon left>expand_more</v-icon>
          <span>FILTERS</span>
        </v-btn>
        <v-list>
          <v-list-tile
            v-for="theme in themes"
            :key="theme.title"
            @click="updateFilter(theme.title)"
          >
            <v-icon left>{{ theme.icon }}</v-icon>
            <v-list-tile-title>{{ theme.title }}</v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-menu>
    </div>
  </v-toolbar>
</template>

<script>
import { mapState } from "vuex";

export default {
  methods: {
    updateFilter(filter) {
      this.$store.commit("updateFilter", filter);
    }
  },
  computed: {
    ...mapState({
      filter: state => state.filter,
      themes: state => state.themes
    })
  }
};
</script>
