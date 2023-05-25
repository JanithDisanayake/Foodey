<script>
    export let data
    const user = data.user
</script>

<h1 class="text-4xl font-bold ml-12"> {data.title} </h1>
<table class="ml-12 mb-12">
    <tr>
        <td class="p-4 text-xl pr-9 w-10">
            ID:
        </td>
        <td class="p-5">
            {user.ID}
        </td>
    </tr>
    <tr>
        <td class="p-4 text-xl pr-9 w-10">
            Name:
        </td>
        <td class="p-5">
            {user.Name}
        </td>
    </tr>
    <tr>
        <td class="p-4 text-xl pr-9 w-10">
            AGE:
        </td>
        <td class="p-5">
            {user.Age}
        </td>
    </tr>
</table>
