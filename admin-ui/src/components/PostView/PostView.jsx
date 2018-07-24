import React, { Component } from "react"
import ReactMarkdown from "react-markdown"
import moment from "moment"

const PostView = ({ post }) => {
  return (
    <div style={style}>
      <h1 style={titleStyle}>{post.title}</h1>
      <p style={timeStyle}>{moment(post.timestamp * 1000).fromNow()}</p>
      <ReactMarkdown source={post.content} />
    </div>
  )
}

const style = {
  flex: 1,
  margin: "1em",
  flexBasis: "25em",
}

const titleStyle = {
  marginBottom: 0,
}

const timeStyle = {
  margin: 0,
  padding: 0,
  fontSize: "0.8em",
}

export default PostView
