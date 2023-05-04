export const load = async (loadEvent) => {
    const { fetch, params } = loadEvent
    const { userId } = params
    const title = 'User Details';
    const response = await fetch(`http://localhost:4000/users/${userId}`)
    const user = await response.json()

    return {
        title,
        user
    }
}