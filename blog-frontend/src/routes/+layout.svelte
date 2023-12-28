<script>
  import "../app.pcss";
  import 'bytemd/dist/index.css';
  import Header from "$lib/components/Header.svelte";
</script>

<div>
  <Header />
  <div>
    <slot />
  </div>
  <!-- <Footer /> -->
  
</div>