<script lang="ts">
  import { Toast } from 'flowbite-svelte';
  import { CheckCircleSolid, ExclamationCircleSolid } from 'flowbite-svelte-icons';
	import '../app.css';

	let { data, children } = $props();
  const { flash } = data;
</script>

<div class="relative w-full pt-10">
  {#if flash.message}
    <Toast color={flash.color} class="absolute left-1/2 -translate-x-1/2">
      <svelte:fragment slot="icon">
        {#if flash.color === "red"}
          <ExclamationCircleSolid class="w-5 h-5" />
          <span class="sr-only">Alert icon</span>
        {:else}
          <CheckCircleSolid class="w-5 h-5" />
          <span class="sr-only">Success icon</span>
        {/if}
      </svelte:fragment>
      {flash.message}
    </Toast>
  {/if}
</div>

{@render children()}
