<script lang="ts">
  import { Input, Label, Button } from 'flowbite-svelte';
  import type { ActionData } from './$types';

	let { form }: { form: ActionData } = $props();
</script>

<div class="flex items-center justify-center min-h-screen">
  <form method="POST" class="w-full max-w-sm p-6 rounded-lg shadow-lg bg-gray-800">
    <h2 class="text-2xl font-bold text-center text-white">Sign Up</h2>

    <div class="mt-4">
      <Label for="username" class="block mb-2">Username</Label>
      <Input type="text" id="username" name="username" value={form?.username ?? ''} required />
    </div>

    <div class="mt-4">
      <Label for="password" class="block mb-2">Password</Label>
      <Input type="password" id="password" name="password" value={form?.password ?? ''} required />
    </div>

    <Button type="submit" class="w-full mt-6">Sign Up</Button>

    <Button href="/login" color="alternative" class="w-full mt-4">
       Log In
    </Button>
  </form>
</div>
