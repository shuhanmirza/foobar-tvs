<template>
  <div data-app>
    <v-data-table
        :headers="headers"
        :items="eventList"
        class="elevation-21"
        style="border-radius: 16px"
    >

      <template v-slot:top>
        <v-card flat style="border-radius: 16px; margin-bottom: -10px">
          <v-card-title color="grey darken-2 white--text">
            {{ tableTitle }}
          </v-card-title>
          <v-card-text style="margin-top: -12px">
            List of Events! - <span class="text-button" @click="createItem()"> Create </span>
          </v-card-text>
        </v-card>

        <v-divider/>

        <v-dialog
            v-model="createEventDialog"
            max-width="500px"
        >
          <v-card>
            <v-card-title>
              <span class="text-h5">{{ formTitle }}</span>
            </v-card-title>

            <v-card-text>
              <v-container>
                <v-row>
                  <v-col
                      cols="12"
                      sm="6"
                      md="4"
                  >
                    <v-text-field
                        v-model="editedItem.name"
                        label="Dessert name"
                    ></v-text-field>
                  </v-col>
                  <v-col
                      cols="12"
                      sm="6"
                      md="4"
                  >
                    <v-text-field
                        v-model="editedItem.calories"
                        label="Calories"
                    ></v-text-field>
                  </v-col>
                  <v-col
                      cols="12"
                      sm="6"
                      md="4"
                  >
                    <v-text-field
                        v-model="editedItem.fat"
                        label="Fat (g)"
                    ></v-text-field>
                  </v-col>
                  <v-col
                      cols="12"
                      sm="6"
                      md="4"
                  >
                    <v-text-field
                        v-model="editedItem.carbs"
                        label="Carbs (g)"
                    ></v-text-field>
                  </v-col>
                  <v-col
                      cols="12"
                      sm="6"
                      md="4"
                  >
                    <v-text-field
                        v-model="editedItem.protein"
                        label="Protein (g)"
                    ></v-text-field>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                  color="blue darken-1"
                  text
                  @click="close"
              >
                Cancel
              </v-btn>
              <v-btn
                  color="blue darken-1"
                  text
                  @click="save"
              >
                Save
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <v-dialog v-model="dialogDelete" max-width="500px">
          <v-card>
            <v-card-title class="text-h5">Are you sure you want to delete this item?</v-card-title>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" text @click="closeDelete">Cancel</v-btn>
              <v-btn color="blue darken-1" text @click="deleteItemConfirm">OK</v-btn>
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
      <template v-slot:no-data>
        <v-btn
            color="primary"
            @click="initialize"
        >
          Reset
        </v-btn>
      </template>
    </v-data-table>
  </div>
</template>

<script>
export default {
  name: "EventTable",
  data: () => ({
    tableTitle: "Events",
    createEventDialog: false,
    dialogDelete: false,
    headers: [
      {text: 'Name', align: 'start', sortable: false, value: 'name'},
      {text: 'Location', value: 'location'},
      {text: 'Time', value: 'datetime'},
      {text: 'Actions', value: 'actions', sortable: false},
    ],
    eventList: [],
    editedIndex: -1,
    editedItem: {
      name: '',
      calories: 0,
      fat: 0,
      carbs: 0,
      protein: 0,
    },
    defaultItem: {
      name: '',
      calories: 0,
      fat: 0,
      carbs: 0,
      protein: 0,
    },
  }),
  methods: {
    initialize() {
      this.eventList = [
        {
          name: 'test event 1',
          location: "Russia",
          datetime: "1667404666"
        },
        {
          name: 'test event 2',
          location: "Russia",
          datetime: "1667404666"
        },
        {
          name: 'test event 3',
          location: "Russia",
          datetime: "1667404666"
        },
        {
          name: 'test event 4',
          location: "Russia",
          datetime: "1667404666"
        },
      ]
    },

    createItem() {
      this.createEventDialog = true
    },
    editItem(item) {
      this.editedIndex = this.eventList.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.createEventDialog = true
    },

    deleteItem(item) {
      this.editedIndex = this.eventList.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
    },

    deleteItemConfirm() {
      this.eventList.splice(this.editedIndex, 1)
      this.closeDelete()
    },

    close() {
      this.createEventDialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    closeDelete() {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    save() {
      if (this.editedIndex > -1) {
        Object.assign(this.eventList[this.editedIndex], this.editedItem)
      } else {
        this.eventList.push(this.editedItem)
      }
      this.close()
    },
    getLocationList() {
      this.$API_CLIENT.get(this.$API_PATH.LOCATION_LIST).then(({data}) => {
        console.log(data)
        this.url = data.url
      }).catch(({response}) => {
        console.log(response)
      });
    }
  },
  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
    },
  },
  watch: {
    dialog(val) {
      val || this.close()
    },
    dialogDelete(val) {
      val || this.closeDelete()
    },
  },

  created() {
    this.initialize()
    this.getLocationList()
  }
}
</script>


<style scoped>
.text-button {
  color: blue;
  font-weight: bolder;
  text-decoration: underline ;
}
</style>