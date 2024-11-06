<script lang="ts">
  import { Input, Label, Button } from 'flowbite-svelte';
  let username = "";
  let password = "";

  function handleSignup(event: Event) {
    event.preventDefault();

    // TODO: Call the login API.
  }
</script>

<div class="flex items-center justify-center min-h-screen">
  <form
    class="w-full max-w-sm p-6 rounded-lg shadow-lg bg-gray-800"
    on:submit={handleSignup}
  >
    <h2 class="text-2xl font-bold text-center text-white">Sign Up</h2>

    <div class="mt-4">
      <Label for="username" class="block mb-2">Username</Label>
      <Input type="text" id="username" bind:value={username} required />
    </div>

    <div class="mt-4">
      <Label for="password" class="block mb-2">Password</Label>
      <Input type="password" id="password" bind:value={password} required />
    </div>

    <Button type="submit" class="w-full mt-6">Sign Up</Button>

    <Button href="/login" color="alternative" class="w-full mt-4">
       Log In
    </Button>
  </form>
</div>
