<script>
  import Navbar from "../components/Navbar.svelte";

  let url =
    "https://images.pexels.com/photos/1028599/pexels-photo-1028599.jpeg?auto=compress&cs=tinysrgb&w=1400";
</script>

<main>
  <Navbar />
  <div class="font-mono mt-12 ml-12">
    <pre>
        <h1 class="text-5xl">   Best <br />   Ordanic <br />   Foods</h1>
    </pre>
    <pre class=" text-xl"> 
          Elevate your health 
          and taste buds with 
          our farm-fresh,
          organic produce.
    </pre>
  </div>
</main>

<style>
  :global(body) {
    background-image: url("https://images.pexels.com/photos/1028599/pexels-photo-1028599.jpeg?auto=compress&cs=tinysrgb&w=1400");
  }
</style>
