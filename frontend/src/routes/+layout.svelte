<script lang="ts">
	import '../app.css';
	let { children } = $props();
</script>

<div class="dark bg-gray-900">
  {@render children()}
</div>
