<template>
  <div data-app>
    <v-data-table
        :headers="headers"
        :items="eventList"
        class="elevation-2"
        style="border-radius: 16px"
        :loading="loading"
        :hide-default-footer="true"
        loading-text="Loading... Please wait"
        no-data-text="No results found for your query"
    >

      <template v-slot:top>
        <v-card flat style="border-radius: 16px; margin-bottom: -10px">
          <v-card-title color="grey darken-2 white--text">
            {{ tableTitle }}
          </v-card-title>
          <v-card-text style="margin-top: -12px">
            List of Events! - <span class="text-button" @click="showCreateItemDialog()"> Create </span>
          </v-card-text>
        </v-card>

        <v-divider/>

        <v-dialog v-model="createItemDialog" max-width="500px">
          <v-card>
            <v-card-title>
              <span>{{ createEventDialogTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-form v-model="isCreateItemFormValid" ref="createItemForm">
                  <v-row>
                    <v-col cols="12" sm="6" md="4">
                      <v-text-field
                          v-model="editedItem.name"
                          label="Name"
                          :rules="[nameRules.required, nameRules.policy]"
                          required
                      ></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6" md="4">
                      <v-select
                          v-model="editedItem.location"
                          :items="locationList"
                          item-text="name"
                          label="Location"
                          return-object
                      >

                      </v-select>
                    </v-col>
                    <v-col cols="12" sm="6" md="4" style="justify-content: center">
                      <v-text-field
                          v-model="editedItem.datetime"
                          label="Datetime"
                          :rules="[datetimeRules.required, datetimeRules.policy]"
                          required
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-form>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="onCloseButtonClicked">
                Cancel
              </v-btn>
              <v-btn color="success" :disabled="!isCreateItemFormValid" text @click="onSaveButtonClick">
                Create
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <v-dialog v-model="dialogDelete" max-width="510px">
          <v-card>
            <v-card-title class="text-h5">Are you sure you want to delete this event?</v-card-title>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="success" text @click="closeDelete">Cancel</v-btn>
              <v-btn color="warning" text @click="deleteItemConfirm">OK</v-btn>
              <v-spacer></v-spacer>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </template>
      <template v-slot:[`item.actions`]="{ item }">
        <v-icon
            small
            class="mr-2"
            @click="editItem(item)"
        >
          mdi-pencil
        </v-icon>
        <v-icon
            small
            @click="deleteItem(item)"
        >
          mdi-delete
        </v-icon>
      </template>
    </v-data-table>

    <pagination-table
        :item-count="eventList.length"
        :pageNumber="pageNumber"
        :is-disabled="loading"
        :page-size="pageSize"
        @next="getNext"
        @previous="getPrevious"
    ></pagination-table>

  </div>
</template>

<script>
import moment from 'moment';
import CONSTANTS from "@/util/helper";
import PaginationTable from "@/components/PaginationTable";
import API_PATHS from "@/util/apipath";

export default {
  name: "EventTable",
  components: {
    PaginationTable
  },
  data: () => ({
    loading: false,
    pageNumber: 0,
    pageSize: 10,
    tableTitle: "Events",
    createItemDialog: false,
    dialogDelete: false,
    headers: [
      {text: 'Name', align: 'start', sortable: false, value: 'name'},
      {text: 'Location', value: 'location'},
      {text: 'Time', value: 'datetime'},
      {text: 'Actions', value: 'actions', sortable: false},
    ],
    eventList: [],
    locationList: [],
    editedItemId: -1,
    editedItemIndex: -1,
    editedItem: {
      name: '', location: {}, datetime: 0,
    },
    defaultItem: {
      name: '', location: {}, datetime: moment().format(CONSTANTS.TIME_FORMAT)
    },
    nameRules: {
      required: v => !!v || 'Name is required',
      policy: v => v.length >= 3 ||
          'Name must have minimum of three characters'
    },
    datetimeRules: {
      required: v => !!v || 'Datetime is required',
      policy: v => (moment(v, CONSTANTS.TIME_FORMAT).isValid() && v.length === 24) ||
          'Please follow 2022-11-04T08:18:48+0600 format'
    },
    isCreateItemFormValid: true
  }),
  methods: {
    initialize() {
      this.getLocationListApi()
      this.getEventListApi()
    },

    showCreateItemDialog() {
      this.editedItemId = -1
      this.editedItem.location = this.locationList.at(0)
      this.editedItem.datetime = moment().format(CONSTANTS.TIME_FORMAT)
      this.createItemDialog = true
    },
    editItem(item) {
      this.editedItemId = this.eventList.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.createItemDialog = true
    },

    deleteItem(item) {
      this.editedItemId = item.id
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
    },

    deleteItemConfirm() {
      this.deleteEventApi(this.editedItemId)
      this.$nextTick( () => {
        this.getEventListApi()
      })
      this.closeDelete()
    },

    onCloseButtonClicked() {
      this.createItemDialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedItemId = -1
      })
    },

    closeDelete() {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedItemId = -1
      })
    },
    onSaveButtonClick() {
      if (this.editedItemId > -1) {
        Object.assign(this.eventList[this.editedItemId], this.editedItem)
      } else {
        this.createEventApi(this.editedItem)
      }
      this.onCloseButtonClicked()
    },
    getNext() {
      this.pageNumber++
      this.getEventListApi()
    },
    getPrevious() {
      this.pageNumber--
      this.getEventListApi()
    },
    createEventApi(item) {
      console.log(item)
      let payload = {
        "name": item.name,
        "location_id": item.location.id,
        "datetime": moment(item.datetime, CONSTANTS.TIME_FORMAT).unix()
      }

      this.$API_CLIENT.post(API_PATHS.CREATE_EVENT, payload).then(({data}) => {
        console.log(data)
        this.url = data.url
      }).catch(({response}) => {
        console.log(response)
        alert("API FAILED")
      });
    },
    getLocationListApi() {
      this.loading = true
      this.$API_CLIENT.get(API_PATHS.LOCATION_LIST).then(({data}) => {
        this.locationList = data.location_list
        console.log(this.locationList)
      }).catch(({response}) => {
        console.log(response)
        alert("API FAILED")
      });
      this.loading = false
    },
    getEventListApi() {
      this.loading = true

      let paramData = {
        page_number: this.pageNumber,
        page_size: this.pageSize
      }
      let eventListResponse = []
      this.$API_CLIENT.getWithParam(API_PATHS.GET_EVENT_LIST, paramData).then(({data}) => {
        console.log(data)

        eventListResponse = data.events;
        this.eventList = []
        if (eventListResponse !== null) {
          for (let index = 0; index < eventListResponse.length; index += 1) {
            eventListResponse.at(index).datetime = moment.unix(eventListResponse.at(index).datetime).format(CONSTANTS.TIME_FORMAT)
            this.eventList.push(eventListResponse.at(index))
          }
        }

      }).catch(({response}) => {
        console.log(response)
        alert("API FAILED")
      });

      this.loading = false
    },

    deleteEventApi(eventId) {
      this.loading = true
      this.$API_CLIENT.delete(API_PATHS.DELETE_EVENT + eventId).then(({data}) => {
        console.log(data)
      }).catch(({response}) => {
        console.log(response)
        alert("API FAILED")
      });
      this.loading = false
    },

    test() {
      console.log(CONSTANTS.TIME_FORMAT)
      let now = moment();
      console.log(now.format("YYYY-MM-DDTHH:mm:ssZZ"))

      let date = moment('2022-11-04T08:07:05+0600', "YYYY-MM-DDTHH:mm:ssZZ")
      console.log(date.unix())

      console.log(moment.unix(1667537276).format(CONSTANTS.TIME_FORMAT))
    }
  },
  computed: {
    createEventDialogTitle() {
      return this.editedItemId === -1 ? 'Create Event' : 'Edit Event'
    },
  },
  watch: {
    dialog(val) {
      val || this.onCloseButtonClicked()
    },
    dialogDelete(val) {
      val || this.closeDelete()
    },
  },

  created() {
    this.test()
    this.initialize()
  }
}
</script>


<style scoped>
.text-button {
  color: blue;
  font-weight: bolder;
  text-decoration: underline;
}
</style>