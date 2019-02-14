<template>
  <v-app>
    <the-app-toolbar />
    <v-content>
      <router-view />
    </v-content>
    <the-footer />
  </v-app>
</template>

<script>
import TheFooter from "@/components/TheFooter";
import TheAppToolbar from "@/components/TheAppToolbar";

export default {
  name: "App",
  components: {
    TheFooter,
    TheAppToolbar
  }
};
</script>
