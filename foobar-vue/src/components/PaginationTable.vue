<template>
  <div class="row " style="width: 100%; margin-left: 0; margin-bottom: 20px; margin-top: 15px">
    <div style="display: flex; flex-direction: row;">
      <v-btn v-if="showPrevBtn"
             :disabled="isDisabled"
             @click="emitPrevious"
             outlined
             color="#e2136e"
             style="margin-right: 10px">
        <v-icon left>
          mdi-arrow-left
        </v-icon>
        Previous
      </v-btn>
      <v-btn
          :disabled="isDisabled"
          @click="emitNext"
          v-if="showNextBtn"
          outlined
          color="#e2136e">
        Next
        <v-icon
            right
            dark>
          mdi-arrow-right
        </v-icon>
      </v-btn>
    </div>
    <v-spacer></v-spacer>
    <div v-if="itemCount > 0">Showing items from {{ firstItemCursor }} to {{ lastItemCursor }}</div>
  </div>
</template>

<script>
export default {
  name: "PaginationTable",
  props: {
    pageNumber: Number,
    itemCount: Number,
    pageSize: Number,
    isDisabled: Boolean,
  },
  computed: {
    showNextBtn() {
      return this.itemCount === this.pageSize
    },
    showPrevBtn() {
      return this.pageNumber > 0;
    },
    firstItemCursor() {
      return (this.pageNumber*this.pageSize) + 1;
    },
    lastItemCursor() {
      return (this.pageNumber*this.pageSize)+ this.itemCount;
    }
  },
  methods: {
    emitNext() {
      this.$emit('next', this.pageNumber)
    },
    emitPrevious() {
      this.$emit('previous', this.pageNumber);
    }
  }
}
</script>