import React, { Component } from "react"
import {
  TextField,
  Label,
  DefaultButton,
  MessageBar,
  MessageBarType,
} from "office-ui-fabric-react"
import ReactMarkdown from "react-markdown"
import { submitPost } from "../../API"
import { withRouter } from "react-router"

class Editor extends Component {
  state = {
    title: "",
    content: "",
    ok: true,
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

  clickSubmit = () => {
    submitPost({ ...this.state }).then(res => {
      if (res.status === 201) {
        this.props.history.push("/")
      } else {
        this.setState({ ok: false })
      }
    })
  }

  render() {
    const disableButton =
      this.state.title.length === 0 || this.state.content.length === 0
    return (
      <div style={editorStyle}>
        {!this.state.ok && (
          <MessageBar messageBarType={MessageBarType.error} isMultiline={false}>
            Something went wrong, unable to publish post. Try again later.
          </MessageBar>
        )}
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
        <Label>Markdown Preview</Label>
        <div style={box}>
          <ReactMarkdown source={this.state.content} />
        </div>
        <DefaultButton
          primary={true}
          text="Publish"
          onClick={this.clickSubmit}
          style={btnStyle}
          disabled={disableButton}
        />
      </div>
    )
  }
}

const editorStyle = {
  margin: "2em",
  flexBasis: "50em",
  display: "flex",
  flexDirection: "column",
}

const btnStyle = {
  width: "3em",
  alignSelf: "flex-end",
  marginTop: "1em",
}

const box = {
  border: "1px solid #a6a6a6",
  minHeight: "1em",
  padding: "0.5em",
}

export default withRouter(Editor)
