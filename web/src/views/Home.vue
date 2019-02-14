<template>
  <v-content>
    <the-edit-event-dialog />
    <the-delete-event-dialog />
    <the-timeline />
  </v-content>
</template>

<script>
import TheTimeline from "@/components/TheTimeline";
import TheEditEventDialog from "@/components/TheEditEventDialog";
import TheDeleteEventDialog from "@/components/TheDeleteEventDialog";

export default {
  components: {
    TheTimeline,
    TheEditEventDialog,
    TheDeleteEventDialog
  }
};
</script>
