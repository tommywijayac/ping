<template>
	<div>
    <RoomCard 
      v-for="(room, index) in rooms"
      :title="room.title"
      :index="index"
      :key="room.id" />
    <button v-on:click="sendMessage('hello')">Send Message</button>
  </div>
</template>

<script>
import RoomCard from './components/RoomCard.vue'

export default {
  name: 'App',
  components: {
    RoomCard
  },
  data: function() {
    return {
      rooms: {},
      connection: null
    }
  },
  methods: {
    sendMessage: function(message) {
      this.connection.send(message);
    }
  },
  created: function(){
    console.log("Starting connection to WebSocket Server")
    this.connection = new WebSocket("ws://localhost:3000/ping")

    //store app context
    const app = this

    this.connection.onmessage = function(event) {
      var msg = JSON.parse(event.data)
      var rooms = []

      for (var i = 0; i < msg.length; i++) {
        var room = {}
        room.id = msg[i].id
        room.title = msg[i].title
        rooms.push(room)
      }

      app.rooms = rooms
    }

    this.connection.onopen = function(event) {
      console.log(event)
      console.log("Successfully connected to the echo websocket server...")
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
