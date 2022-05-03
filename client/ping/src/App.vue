<template>
	<div class="grid">
    <RoomCard v-for="(room, index) in rooms"
      :index="index"
      :key="room.id"
      :title="room.title"
      :state="room.state"
      @pingAck="onPingAck(room, room.id)"
    />
    <!-- <RoomCard :state="'active'" :title="'test'" :key="333" :id='333'
      @pingAck="onPingAck(room)"
    /> -->
    <button v-on:click="sendMessage('hello')">Dummy trigger (simulate BE trigger)</button>
  </div>
</template>

<script>
import RoomCard from './components/RoomCard.vue'

export default {
  name: 'App',
  components: {
    RoomCard,
  },
  data() {
    return {
      rooms: [{
        id: 0,
        title: '<placeholder-title>',
        state: 'active',
      },{
        id: 1,
        title: '<placeholder-title>',
        state: 'warning',
      },{
        id: 2,
        title: '<placeholder-title>',
        state: '',
      },{
        id: 3,
        title: '<placeholder-title>',
        state: 'active',
      },{
        id: 4,
        title: '<placeholder-title>',
        state: '',
      }],
      connection: null
    }
  },
  watch: {
    rooms: {
      handler(newValue, oldValue) {
        console.log("hello from watcher", newValue, oldValue)
      },
    }
  },
  methods: {
    sendMessage: function(message) {
      this.connection.send(message);
    },
    onPingAck: function(room, roomId) {
      console.log(room);
      console.log(roomId);
      console.log(this.rooms);

      // const toChangeRoom = this.rooms.find(r => r.id === roomId);
      // toChangeRoom.state = '';

      //this.rooms.push({id:333, title:"pushed from FE", state: ""})

      //this.connection.send('hellaw from onPingAck');
    },
  },
  created: function(){
    console.log("Starting connection to WebSocket Server")
    this.connection = new WebSocket("ws://localhost:3000/ping")

    //store app context
    const app = this

    this.connection.onmessage = function(event) {
      var msg = JSON.parse(event.data)
      console.log(msg)

      //TODO: try mutate directly here
      //Result: 
      // var ref = msg[0].id
      // const room = app.rooms.find(r => r.id === ref)
      // room.state = "active"

      //TODO: try modifying by push
      //Result: 
      app.rooms.push({id:333, title:"pushed from FE", state: ""})

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
  height: 100vh;
  box-sizing: border-box;
  margin: 0px !important; /* override vue default 8px margin */
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100%;
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
