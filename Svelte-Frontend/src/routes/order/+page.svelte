<script>

    import Ordercard from "../../components/OrderCard.svelte";
    
    export let data 
    const orders = data.orders
    
</script>

<div class="grid grid-cols-3 gap-10 mx-10 my-10">

    {#each orders as product}
        <Ordercard product={product}/>
        
    {/each}
    
</div>
