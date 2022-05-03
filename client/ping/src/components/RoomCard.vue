<template>
  <div
    ref='card'
    :class="[
      'card',
      { 'card-pulse': localState === 'active' },
      { 'card-pulse-warn': localState === 'warning' }
    ]"
    v-on="localState === 'active' || localState === 'warning' ? 
      { click: pingAck } : {}"
  >
    <div ref='content' class='content'>
      <span class="title">{{ title }}</span>
      <span>image</span>
      <span>time</span>
    </div>
    <div ref='checkmarkcontainer' class='checkmark_container' hidden>
      <svg ref='checkmark' class="checkmark" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 52 52">
        <circle class="checkmark__circle" cx="26" cy="26" r="25" fill="none"/>
        <path class="checkmark__check" fill="none" d="M14.1 27.2l7.1 7.2 16.7-16.8"/>
      </svg>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: { required: true, type: String},
    state: { default: '', type: String}
  },
  data() {
    return {
      localState: this.state,
      
      //flag to make sure only one animation is played at a time
      isAnimeFinished: true,
    }
  },
  watch: {
    localState(oldState, newState) {
      console.log("inside watcher")
      if (oldState !== newState) {
        console.log("changed..?")
        this.localState = newState;
      }
    }
  },
  methods: {
    pingAck: function(){
      if (!this.isAnimeFinished) {
        return
      }
      this.isAnimeFinished = false;

      var card = this.$refs.card;
      card.classList.remove('card-pulse');
      card.classList.remove('card-pulse-warn');

      var content = this.$refs.content;
      content.classList.add('blur');

      //trigger check animation
      var check = this.$refs.checkmarkcontainer;
      check.removeAttribute('hidden');

      //easy approach: wait until we're sure all animation are done
      //with big enough value
      setTimeout(this.pingAckAnimationFinish, 2500);
      
      //notify server
      this.$emit('pingAck');
    },
    pingAckAnimationFinish: function() {
      //clean up class
      var card = this.$refs.card;
      card.classList.remove('card-transition')

      var content = this.$refs.content;
      content.classList.remove('blur')

      var check = this.$refs.checkmarkcontainer;
      check.setAttribute('hidden', 'hidden');

      //enable animation
      this.isAnimeFinished = true;
    }
  }
}
</script>

<style scoped>
/*****************************************************/
/********************* CARD **************************/
.card {
  border: 1px solid black;
  background-color: #333333;

  /* disable text selection */
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;

  /* configure elements to be centered vertically & horizontally */
  /* need to do here as well (not only in content)*/
  align-items: center;
  display: flex;
  flex-direction: column;
  justify-content: center;
  
  /* so we can place div.check on top of div.content*/
  position: relative;

  transition: background-color 0.2s linear;
}

.card-pulse {
  background-color: white;
  box-shadow: 0 0 0 0 rgba(0, 81, 255, 1);
  animation: pulse 1s infinite;
}

.card-pulse-warn {
  background-color: rgba(255, 0, 0, 0.3);
  box-shadow: 0 0 0 0 rgba(255,0 ,0, 1);
  animation: pulse-warn 1s infinite;
}

@keyframes pulse {
  0% {
    box-shadow: 0 0 0 0 rgba(0, 81, 255, 0.7);
  }

  70% {
    box-shadow: 0 0 0 10px rgba(0, 0, 0, 0);
  }

  100% {
    box-shadow: 0 0 0 0 rgba(0, 0, 0, 0);
  }
}

@keyframes pulse-warn {
  0% {
    box-shadow: 0 0 0 0 rgba(255, 0, 0, 0.7);
  }

  70% {
    box-shadow: 0 0 0 10px rgba(0, 0, 0, 0);
  }

  100% {
    box-shadow: 0 0 0 0 rgba(0, 0, 0, 0);
  }
}

</style>

<style scoped>
/*****************************************************/
/******************* COMPONENT ***********************/
/* local variable */
* {
  --checkmark_color: #03a629;
  --checkmark_size: 150px;
}

.content {
  /* configure elements to be centered vertically & horizontally */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  box-sizing: border-box;
  margin: 1rem;
}

.checkmark_container {
  /* position at parent's center */
  /* https://stackoverflow.com/questions/2941189/how-to-overlay-one-div-over-another-div */
  /* https://stackoverflow.com/questions/15328416/position-by-center-point-rather-than-top-left-point */
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%)
}

.checkmark__circle {
  stroke-dasharray: 166;
  stroke-dashoffset: 166;
  stroke-width: 2;
  stroke-miterlimit: 10;
  stroke: var(--checkmark_color);
  fill: none;
  animation: stroke-circle 1.2s cubic-bezier(0.65, 0, 0.45, 1) forwards;
}

.checkmark {
  width: var(--checkmark_size);
  height: var(--checkmark_size);
  border-radius: 50%;
  display: block;
  stroke-width: 2;
  stroke: #fff;
  stroke-miterlimit: 10;
  margin: 10% auto;
  box-shadow: inset 0px 0px 0px var(--checkmark_color);
  animation: fill 1.6s ease-in-out .4s forwards, scale .3s ease-in-out .9s both;
}

.checkmark__check {
  transform-origin: 50% 50%;
  stroke-dasharray: 48;
  stroke-dashoffset: 48;
  animation: stroke-check 1s cubic-bezier(0.65, 0, 0.45, 1) 0.8s forwards;
}

@keyframes stroke-circle {
  50% {
    stroke-dashoffset: 0;
  }
  100% {
    stroke-dashoffset: 166;
  }
}
@keyframes stroke-check {
  40%, 80% {
    stroke-dashoffset: 0;
  }
  100% {
    stroke-dashoffset: 48;
  }
}
@keyframes scale {
  0%, 100% {
    transform: none;
  }
  50% {
    transform: scale3d(1.1, 1.1, 1);
  }
}
@keyframes fill {
  50% {
    box-shadow: inset 0px 0px 0px var(--checkmark_size) var(--checkmark_color);
  }
  100% {
    box-shadow: inset 0px 0px 0px 0px var(--checkmark_color);
  }
}

/* https://stackoverflow.com/questions/11977103/blur-effect-on-a-div-element */
.blur {
  filter: blur(0px);
  animation: blur 2s linear;

  /* transition duration & keyframes is tightly-tied to check animation: */
  /* make sure it triggers after check vanishes, but before class is removed */
}
@keyframes blur {
  0%, 100% {
    filter: blur(0px);
  }
  10%, 80% {
    filter: blur(5px);
  }
}

.title {
  font-size: 3rem;
  font-family: Arial, Helvetica, sans-serif;
}
</style>