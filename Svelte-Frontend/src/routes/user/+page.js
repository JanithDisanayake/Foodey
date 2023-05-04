// export const load = async (loadEvent) => {
//     const { fetch } = loadEvent
//     const title = 'List of available Users';
//     const response = await fetch("http://localhost:4000/users/")
//     const users = await response.json()

//     return {
//         title,
//         users
//     }
// }


export async function load( {data, fetch} ) {
    const response = await fetch(`https://jsonplaceholder.typicode.com/users`)
    const users = await response.json()
    console.log(users)
    return {
        title: 'User Details',
        users: users
    }
}