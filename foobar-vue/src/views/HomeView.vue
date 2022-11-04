<template>
  <div class="home">
    <EventTable class="event-table"></EventTable>
  </div>
</template>

<script>
import EventTable from "@/components/EventTable";
export default {
  name: 'HomeView',
  components: {
    EventTable
  }
}
</script>

<style scoped>
.home{
  padding: 10px;
}
.event-table{
  margin: 100px;
}
</style>
