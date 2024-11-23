<script>
  import { onMount } from 'svelte';

  const eventSource = new EventSource("http://localhost:8080/sse")

  eventSource.onmessage = (event) => {
    console.log(event.data)
  }

  const socket = new WebSocket('ws://localhost:8080/ws')
  socket.onopen = (event) => {
    console.log('Connected')
  }

  socket.onclose = (event) => {
    console.log('Connection closed')
  }

  socket.onmessage = (event) => {
    console.log(event.data)
  }

  onMount(() => {
    return () => socket.close()
  })
</script>

<div class="flex flex-col items-center justify-center min-h-screen">
  <h1 class="text-2xl font-bold text-center">Chat</h1>
  <div class="w-full max-w-sm p-6 rounded-lg shadow-lg bg-gray-800 flex flex-col items-center justify-center">

    <input class="text-black rounded-lg" type="text">
  </div>
</div>
