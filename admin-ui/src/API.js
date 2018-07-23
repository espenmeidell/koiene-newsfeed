const URL = "http://localhost:8000"

export async function getPosts() {
  const data = await fetch(URL + "/posts")
  const json = await data.json()
  return json
}

export async function submitPost(post) {
  const response = await fetch(URL + "/posts", {
    method: "POST",
    headers: new Headers({ "content-type": "application/json" }),
    body: JSON.stringify(post),
  })
  console.log(response)
}
