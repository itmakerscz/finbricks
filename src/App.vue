<template>
  <div>
    <div class="chat-window">
      <h1>Finbricks Chat App</h1>
      <div class="chat-messages">
        <div v-for="(message, index) in messages" :key="index" class="message" @keydown.enter="sendMessage">
          {{ message }}
        </div>
      </div>
      <div class="chat-input">
        <input type="text" v-model="inputMessage" placeholder="Type a message...">
        <button class="send-button" @click="sendMessage" :disabled="!inputMessage && !audioBlob">
        </button>
        <button class="record-button" @click="toggleRecording">
          <span v-if="!recording"></span>
          <span v-else></span>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',
  data() {
    return {
      messages: [],
      recording: false,
      mediaRecorder: null,
      chunks: [],
      stream: null,
      apiUrl: 'http://localhost:8000/transcribe',
      botUrl: 'http://localhost:8000/chatbot',
      inputMessage: '',
      audioBlob: null
    };
  },
  methods: {
    async startRecording() {
      this.recording = true;
      this.chunks = [];
      try {
        this.stream = await navigator.mediaDevices.getUserMedia({ audio: true });
        this.mediaRecorder = new MediaRecorder(this.stream);
        this.mediaRecorder.addEventListener('dataavailable', event => {
          this.chunks.push(event.data);
        });
        this.mediaRecorder.addEventListener('stop', async () => {
          this.recording = false;
          const blob = new Blob(this.chunks, { type: 'audio/webm' });
          this.audioBlob = blob;
        });
        this.mediaRecorder.start();
      } catch (error) {
        console.error(error);
        this.recording = false;
      }
    },
    stopRecording() {
      this.mediaRecorder.stop();
      this.stream.getTracks().forEach(track => track.stop());
    },
    toggleRecording() {
      if (this.recording) {
        this.stopRecording();
      } else {
        this.startRecording();
      }
    },
    async sendMessage() {
      if (this.audioBlob) {
        const formData = new FormData();
        formData.append('audio', this.audioBlob, 'recording.webm');
        try {
          const response = await axios.post(this.apiUrl, formData);
          this.messages.push(response.data);
          this.audioBlob = null;
        } catch (error) {
          console.error(error);
        }

      }
      if (this.inputMessage) {
        this.messages.push(this.inputMessage);
        this.inputMessage = '';
      }

      try {
        const response = await axios.post(this.botUrl, {
          message: this.inputMessage,
      });
      this.messages.push(response.data.message);
      } catch (error) {
        console.error(error);
      }

    }
  }
};
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Roboto:300,300i,400,400i,500,500i,700,700i,900,900i');

html, body {
  margin: 0; padding: 20px; overflow: hidden;
  height: 100%;
  background: linear-gradient(180deg, #2A48E9, #12C3BE);
  font-family: 'Roboto', sans-serif;
}
#app {
  height: 100%;
}
h1{
  text-align: center;  
  color:#fff;
}
/* Chat container */
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

/* Chat messages */
.chat-messages {
  flex-grow: 1;
  padding: 1rem;
  overflow-y: scroll;
  background-color: #fff;
  height: 60vh;
  opacity: 0.4;
}

.message {
  opacity: 1;
}

/* Chat input */
.chat-input {
  display: flex;
  align-items: center;
  padding: 0.5rem 1rem;
  background-color: #fff;
}

/* Chat input text field */
.chat-input input[type="text"] {
  flex-grow: 1;
  padding: 0.5rem;
  border: none;
  border-radius: 0.25rem;
}

/* Chat input send button */
.chat-input .send-button {
  display: inline-block;
  padding: 0.5rem 1rem;
  margin: 0.5rem;
  font-size: 1rem;
  font-weight: bold;
  color: none;
  background-color: none;
  background-size: cover;
  border: none;
  border-radius: 0.25rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
  background-image: url('./assets/send.png');
  background-repeat: no-repeat;
  background-position: center;
  width: 50px;
  height: 50px;
}

/* Chat input send button - disabled state */
.chat-input .send-button:disabled {
  background-color: none;
  cursor: not-allowed;
}

/* Chat input send button - hover state */
.chat-input .send-button:hover:not(:disabled) {
  background-color: none;
}
.chat-messages {
  flex-grow: 1;
  padding: 1rem;
  overflow-y: scroll;
  background-color: #fff;
  flex-direction: column;
  margin: 0.5rem 0;
  opacity:0.4;
}

/* Chat message sender name */
.chat-message .sender-name {
  font-weight: bold;
  margin-bottom: 0.25rem;
}

/* Chat message content */
.chat-message .message-content {
  background-color: #007bff;
  color: #fff;
  padding: 0.5rem;
  border-radius: 0.25rem;
  align-self: flex-start;
}

/* Chat message from current user */
.chat-message.current-user .message-content {
  background-color: none;
  align-self: flex-end;
}

/* Chat input record button */
.chat-input .record-button {
  display: inline-block;
  padding: 0.5rem 1rem;
  margin: 0.5rem;
  font-size: 1rem;
  font-weight: bold;
  color: #fff;
  background-color: none;
  background-size: cover;
  border: none;
  border-radius: 0.25rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
  background-image: url('./assets/voice.png');
  background-repeat: no-repeat;
  background-position: center;
  width: 50px;
  height: 50px;
}

/* Chat input record button - disabled state */
.chat-input .record-button:disabled {
  background-color: none;
  cursor: not-allowed;
}

/* Chat input record button - hover state */
.chat-input .record-button:hover:not(:disabled) {
  background-color: none;
}
</style>