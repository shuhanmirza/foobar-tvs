<template>
  <div></div>

</template>

<script>
export default {
  name: "EventTable",
  methods: {
    async getLocationList() {
      this.$API_CLIENT.get(this.$API_PATH.LOCATION_LIST).then(({data}) => {
        console.log(data)
        this.url = data.url
      }).catch(({response}) => {
        console.log(response)
      });
    }
  },
  created() {
    this.getLocationList()
  }
}
</script>

<style scoped>

</style>