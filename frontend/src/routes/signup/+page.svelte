<script lang="ts">
  import { Helper, Input, Label, Button } from 'flowbite-svelte';
  import type { ActionData } from './$types';

	let { form }: { form: ActionData } = $props();
  const usernameColor = form?.error?.details?.username ? 'red' : undefined;
  const passwordColor = form?.error?.details?.password? 'red' : undefined;
</script>

<div class="flex items-center justify-center min-h-screen">
  <form method="POST" class="w-full max-w-sm p-6 rounded-lg shadow-lg bg-gray-800">
    <h2 class="text-2xl font-bold text-center">Sign Up</h2>

    <div class="mt-4">
      <Label for="username" color={usernameColor} class="block mb-2">Username</Label>
      <Input type="text" id="username" name="username" required value={form?.username ?? ''} color={usernameColor} />
      {#if form?.error?.details?.username}
        <Helper class='mt-2' color='red'>{form.error.details.username}</Helper>
      {/if}
    </div>

    <div class="mt-4">
      <Label for="password" color={passwordColor} class="block mb-2">Password</Label>
      <Input type="password" id="password" name="password" required value={form?.password ?? ''} color={passwordColor} />
      {#if form?.error?.details?.password}
        <Helper class='mt-2' color='red'>{form.error.details.password}</Helper>
      {/if}
    </div>

    <Button type="submit" class="w-full mt-6">Sign Up</Button>

    <Button href="/login" color="alternative" class="w-full mt-4">
       Log In
    </Button>
  </form>
</div>
