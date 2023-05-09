export const load = async (loadEvent) => {
  const { fetch, params } = loadEvent;
  const { userId } = params;
  const title = "User Details";
  const response = await fetch(`http://localhost:3000/users/${userId}`);
  console.log(userId);
  const user = await response.json();

  return {
    title,
    user,
  };
};
