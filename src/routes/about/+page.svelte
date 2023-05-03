<script>
  import Navbar from "../../components/Navbar.svelte";
</script>

<main>
  <div class=" flex items-center justify-center my-12">
    <div
      class="w-full max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8 dark:bg-gray-800 dark:border-gray-700"
    >
      <h1 class="text-center mb-10 text-3xl">About Us</h1>
      <p class="text-justify">
        Welcome to our organic food selling website! We are passionate about
        providing high-quality, nutritious, and delicious organic food to our
        customers. Our mission is to promote healthy eating habits and support
        sustainable agriculture by offering a wide range of organic products
        that are free from harmful chemicals and pesticides.
      </p>

      <div class="mx-12 mt-10 grid grid-cols-3">
        <a href="https://web.facebook.com/zuck/?_rdc=1&_rdr">
          <img src="https://img.icons8.com/material/50/null/facebook-new.png" alt=""/>
        </a>
        <a href="https://www.instagram.com/zuck/?hl=en">
          <img src="https://img.icons8.com/material/50/null/instagram-new--v1.png" alt=""/>
        </a>
        <a href="https://www.linkedin.com/in/mark-zuckerberg-618bba58/">
          <img src="https://img.icons8.com/material/50/null/linkedin--v1.png" alt=""/>
        </a>
      </div>

      <div class="mt-10 mx-10 inline-flex items-center">
        <img src="https://img.icons8.com/material-sharp/24/null/ringer-volume.png" alt=""/>
        <p class="ml-3">+94785957581</p>
      </div>
      <div class="mt-4 mx-10 inline-flex items-center">
        <img src="https://img.icons8.com/ios-filled/25/null/marker.png" alt=""/>
        <p class="ml-3"> 1456 Edgewood Dr, Palo Alto, CA 94301, USA</p>
      </div>
    </div>
  </div>
</main>
