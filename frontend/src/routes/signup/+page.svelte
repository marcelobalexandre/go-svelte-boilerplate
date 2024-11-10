<script lang="ts">
  import { Helper, Input, Label, Button } from 'flowbite-svelte';
  import type { ActionData } from './$types';

	let { form }: { form: ActionData } = $props();
  const usernameColor = form?.errors?.username ? 'red' : undefined;
  const emailColor = form?.errors?.email? 'red' : undefined;
</script>

<div class="flex items-center justify-center min-h-screen">
  <form method="POST" class="w-full max-w-sm p-6 rounded-lg shadow-lg bg-gray-800">
    <h2 class="text-2xl font-bold text-center">Sign Up</h2>

    <div class="mt-4">
      <Label for="username" color={usernameColor} class="block mb-2">Username</Label>
      <Input type="text" id="username" name="username" value={form?.username ?? ''} required color={usernameColor} />
      {#if form?.errors?.username}
        <Helper class='mt-2' color='red'>{form.errors.username}</Helper>
      {/if}
    </div>

    <div class="mt-4">
      <Label for="password" color={emailColor} class="block mb-2">Password</Label>
      <Input type="password" id="password" name="password" value={form?.password ?? ''} required color={emailColor} />
      {#if form?.errors?.email}
        <Helper class='mt-2' color='red'>{form.errors.email}</Helper>
      {/if}
    </div>

    <Button type="submit" class="w-full mt-6">Sign Up</Button>

    <Button href="/login" color="alternative" class="w-full mt-4">
       Log In
    </Button>
  </form>
</div>
