<template>
	<div class="grid">
    <RoomCard v-for="(room, index) in rooms"
      :index="index"
      :key="room.id"
      :title="room.title"
      :state="room.state"
      :icon_path="room.icon_path"
      @pingAck="onPingAck(room)"
    />
  </div>
</template>

<script>
import RoomCard from './components/RoomCard.vue'
import { reactive } from 'vue'

export default {
  name: 'App',
  components: {
    RoomCard,
  },
  data() {
    return {
      rooms: reactive([{
        id: 0,
        title: 'Ruangan A',
        state: '',
        icon_path: 'default.png'
      },{
        id: 1,
        title: 'Ruangan B',
        state: 'active',
        icon_path: 'default.png'
      },{
        id: 2,
        title: 'Ruangan C',
        state: 'warning',
        icon_path: 'default.png'
      }]),
      connection: null
    }
  },
  methods: {
    sendMessage: function(message) {
      this.connection.send(message);
    },
    onPingAck: function(room) {
      //set clicked room state to inactive
      room.state = '';
    },
  },
  created: function(){
    console.log("Starting connection to WebSocket Server")
    this.connection = new WebSocket("ws://localhost:3000/ping")

    //store app context
    const self = this

    this.connection.onmessage = function(event) {
      var msg = JSON.parse(event.data)
      self.rooms = msg;
      console.log(msg)

      // var rooms = []

      // for (var i = 0; i < msg.length; i++) {
      //   var room = {}
      //   room.id = msg[i].id
      //   room.title = msg[i].title
      //   room.state = msg[i].state
      //   rooms.push(room)
      // }

      // app.rooms = rooms
    }

    this.connection.onopen = function(event) {
      console.log(event)
      console.log("Successfully connected to the echo websocket server...")
    }
  }
}
</script>

<style>
body {
  /* abs height value */
  height: 100vh;
  /* old way size (padding&margin) calculation */
  box-sizing: border-box;
  /* override vue default 8px margin */
  margin: 0px !important;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  /* relative height to body */
  height: 100%;
  /* give whitespace */
  box-sizing: border-box;
  padding: 3%;
}

/*****************************************************/
/********************* LAYOUT ************************/
.grid {
  display: grid;
  gap: 30px 30px;
  grid-template-columns: repeat(3, 1fr);
  grid-template-rows: repeat(3, 1fr);
  height: 100%;
}
</style>
