<script>
    export let data
    const user = data.user
</script>

<h1 class="text-4xl font-bold ml-12 my-10"> {data.title} </h1>
<table class="ml-12 mb-12">
    <tr>
        <td class="p-4 text-xl pr-9">
            ID:
        </td>
        <td class="p-5">
            {user.id}
        </td>
    </tr>
    <tr>
        <td class="p-4 text-xl pr-9">
            Name:
        </td>
        <td class="p-5">
            {user.name}
        </td>
    </tr>
    <tr>
        <td class="p-4 text-xl pr-9">
            AGE:
        </td>
        <td class="p-5">
            {user.age}
        </td>
    </tr>
</table>
<div class="p-12"/>
