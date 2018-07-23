import React, { Component } from "react"
import { getPosts } from "../../API"
import PostView from "../PostView/PostView"

class Home extends Component {
  state = {
    posts: [],
  }

  componentDidMount() {
    getPosts().then(data => {
      this.setState({ posts: data })
    })
  }

  render() {
    return (
      <div style={style}>
        {this.state.posts.map(p => <PostView key={p.id} post={p} />)}
      </div>
    )
  }
}

const style = {
  display: "flex",
  flexWrap: "wrap",
}

export default Home
