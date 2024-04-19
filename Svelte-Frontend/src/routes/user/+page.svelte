<script>
  import { queryStore, gql, getContextClient } from "@urql/svelte";
  import {
    Client,
    setContextClient,
    cacheExchange,
    fetchExchange,
  } from "@urql/svelte";
  export let data;
  const users = data.users; // this property

  const client = new Client({
    url: "http://localhost:3000/query",
    exchanges: [cacheExchange, fetchExchange],
  });

  setContextClient(client);

  // Define the GraphQL query
  const GET_USERS = gql`
    query {
      users {
        ID
        Age
        Name
      }
    }
  `;

  // Fetch data using the queryStore
  const todos = queryStore({
    client: getContextClient(),
    query: GET_USERS,
  });

  // Access the data property of the todos store
  $: userData = $todos.data;
</script>

{#if $todos.fetching}
  <p>Loading...</p>
{:else if $todos.error}
  <p>Error: {$todos.error.message}</p>
{:else}
  <!-- Display the JSON data -->
  <pre>{JSON.stringify(userData, null, 2)}</pre>
{/if}

<!-- <h1 class="text-4xl font-bold ml-12 my-5">{data.title}</h1> -->

<!--<div class="relative overflow-x-auto px-10 mb-12">
  <table class="w-full text-sm text-left text-black dark:text-gray-400">
    <thead
      class="text-s text-white font-bold uppercase bg-black dark:bg-gray-700 dark:text-gray-400"
    >
      <tr>
        <th scope="col" class="px-6 py-3"> ID </th>
        <th scope="col" class="px-6 py-3"> Name </th>
        <th scope="col" class="px-6 py-3"> Age </th>
        <th></th>
      </tr>
    </thead>
    <tbody>
      {#each users as user}
        <tr
          class="bg-orange-200 border-b dark:bg-gray-800 dark:border-gray-700"
        >
          <th
            scope="row"
            class="px-6 py-4 font-medium text-blue-900 whitespace-nowrap dark:text-white"
          >
            <a href="user/{user.ID}">{user.ID}</a>
          </th>
          <td class="px-6 py-4">
            {user.Name}
          </td>
          <td class="px-6 py-4">
            {user.Age}
          </td>
          <td>
            <a href="user/{user.ID}">VISIT</a>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>-->
