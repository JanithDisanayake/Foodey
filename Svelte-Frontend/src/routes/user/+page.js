export const load = async (loadEvent) => {
  const { fetch } = loadEvent;
  const title = "List of available Users";
  const response = await fetch("http://localhost:3000/");
  const users = await response.json();

  return {
    title,
    users,
  };
};
