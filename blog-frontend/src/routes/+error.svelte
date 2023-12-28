<script context="module">
  export function load({ error, status }) {
    return {
      props: {
        title: `${status}: ${error.message}`,
        status: status,
      },
    };
  }
</script>

<script lang="ts">
	import NotFound from "$lib/components/NotFound.svelte";

  export let title: string;
  export let status: number;
</script>


<NotFound />
