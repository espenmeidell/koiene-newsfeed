import React, { Component } from "react"
import { TextField, Label } from "office-ui-fabric-react"
import ReactMarkdown from "react-markdown"

class Editor extends Component {
  state = {
    title: "",
    content: "",
  }

  titleChange = title => {
    this.setState({
      title,
    })
  }

  contentChange = content => {
    this.setState({
      content,
    })
  }

  render() {
    return (
      <div>
        <TextField
          label="Story Title"
          required={true}
          value={this.state.title}
          onChanged={this.titleChange}
        />
        <TextField
          label="Content"
          multiline
          rows={4}
          required={true}
          value={this.state.content}
          onChanged={this.contentChange}
        />
        <Label>Preview</Label>
        <ReactMarkdown source={this.state.content} />
      </div>
    )
  }
}

export default Editor
