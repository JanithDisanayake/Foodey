<script>
  import "../app.css";
  import Footer from "../components/Footer.svelte";
  import Navbar from "../components/Navbar.svelte";
</script>

<div class="flex flex-col h-screen justify-between">
    <Navbar/>
    <slot />
    <Footer/>
</div>