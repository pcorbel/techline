<template>
  <v-footer height="auto" class="white pa-3">
    <span class="caption"> Last event : {{ getLastEvent }} </span>
    <v-spacer />
    <a href="https://github.com/pcorbel/techline">
      <span class="caption mr-1"> Contribute on pcorbel/techline </span>
      <v-icon color="black" small right>fab fa-github</v-icon>
    </a>
  </v-footer>
</template>

<script>


export default {
  computed: {
    getLastEvent() {
      return this.$store.getters.lastEventCreatedAt;
    }
  }
};
</script>
