<script>
  import { each } from "svelte/internal";
    import Ordercard from "../../components/OrderCard.svelte";
    let products = [
        { 
            name: "Spaghetti",
            desc: "long, thin, solid, cylindrical pasta. It is a staple food of traditional Italian cuisine.",
            image: "https://images.pexels.com/photos/1279330/pexels-photo-1279330.jpeg?auto=compress&cs=tinysrgb&w=400"
        },
        {
            name: "Pasta",
            desc: "typically made from wheat flour mixed with water or eggs, and formed into sheets or other shapes, then cooked by boiling or baking.",
            image: "https://images.pexels.com/photos/1437267/pexels-photo-1437267.jpeg?auto=compress&cs=tinysrgb&w=600"
        },
        {
            name: "Pizza",
            desc: "dish of Italian origin consisting of flat base of leavened wheat-based dough topped with tomatoes, cheese, and other ingredients.",
            image: "https://images.pexels.com/photos/842519/pexels-photo-842519.jpeg?auto=compress&cs=tinysrgb&w=600"
        },
        {
            name: "Biriyani",
            desc: "mixed rice dish originating among South Asians. It is made with Indian spices, vegetables, rice, and some type of meat.",
            image: "https://media.istockphoto.com/id/1345624336/photo/chicken-biriyani.jpg?b=1&s=612x612&w=0&k=20&c=2CBPkgQI2PiNySSNi1tvwVSYFqOBgfgrgDPCyelIVrk="
        },
        {
            name: "Sushi",
            desc: "Japanese dish of prepared vinegared rice, usually with some sugar and salt, accompanied by a variety of ingredients, such as seafood raw and vegetables.",
            image: "https://images.pexels.com/photos/357756/pexels-photo-357756.jpeg?auto=compress&cs=tinysrgb&w=600"

        },
        {
            name: "Kottu",
            desc: "Sri Lankan dish consisting of chopped roti, a meat curry dish of choice, along with scrambled egg, onions, and chillies.",
            image: "https://images.pexels.com/photos/2313686/pexels-photo-2313686.jpeg?auto=compress&cs=tinysrgb&w=600"
        }
    ]
</script>

<div class="grid grid-cols-3 gap-10 mx-10 my-10">

    {#each products as product}
        <Ordercard product={product}/>
        
    {/each}
    
</div>
