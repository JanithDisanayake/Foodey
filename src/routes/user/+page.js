export const load = async (loadEvent) => {
    const { fetch } = loadEvent
    const title = 'List of available Users';
    const response = await fetch("http://localhost:4000/users/")
    const users = await response.json()

    return {
        title,
        users
    }
}



// import { onMount } from 'svelte';
    // onMount(async () => {
    //     const response = await fetch(`http://localhost:3000/`, { mode: 'no-cors', method: "get" });
    //     users = await response
    //     console.log(users)
    // })