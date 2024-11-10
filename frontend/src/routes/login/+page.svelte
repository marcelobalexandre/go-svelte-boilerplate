<script lang="ts">
  import { Alert, Input, Label, Button } from 'flowbite-svelte';
  import type { ActionData } from './$types';

	let { form }: { form: ActionData } = $props();
  const inputColor = form?.errorMessage ? 'red' : undefined;
</script>

<div class="flex items-center justify-center min-h-screen">
  <form method="POST" class="w-full max-w-sm p-6 rounded-lg shadow-lg bg-gray-800">
    <h2 class="text-2xl font-bold text-center">Log In</h2>

    {#if form?.errorMessage}
      <Alert color="red" rounded={false} class="mt-4 border-t-4 dark:bg-gray-700">
        {form.errorMessage}
      </Alert>
    {/if}
      

    <div class="mt-4">
      <Label for="username" color={inputColor} class="block mb-2">Username</Label>
      <Input type="text" id="username" name="username" value={form?.username ?? ''} required color={inputColor} />
    </div>

    <div class="mt-4">
      <Label for="password" color={inputColor} class="block mb-2">Password</Label>
      <Input type="password" id="password" name="password" value={form?.password ?? ''} required color={inputColor} />
    </div>

    <Button type="submit" class="w-full mt-6">Log In</Button>

    <Button href="/signup" color="alternative" class="w-full mt-4">
      Sign Up
    </Button>
  </form>
</div>
